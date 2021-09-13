package processor

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/google/uuid"
	serversv1 "github.com/vniche/collective/protocol/v1"
	"github.com/vniche/collective/server/datastore"
	"github.com/vniche/collective/server/stream"
	"go.mongodb.org/mongo-driver/mongo"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gopkg.in/mgo.v2/bson"
)

const groupName = "requests-processing"

func Start() {
	mongoDB, ok := os.LookupEnv("MONGO_DB")
	if !ok {
		mongoDB = "collective"
	}

	mongoURI, ok := os.LookupEnv("MONGO_URI")
	if !ok {
		mongoURI = "mongodb://localhost/" + mongoDB
	}

	client, err := datastore.Connect(mongoURI)
	if err != nil {
		log.Printf("unable to connect to datastore: %v\n", err)
	}

	database := client.Database(mongoDB)

	kafkaHost, ok := os.LookupEnv("KAFKA_HOST")
	if !ok {
		kafkaHost = "localhost"
	}

	var kafkaConfig *kafka.ConfigMap

	devMode, ok := os.LookupEnv("DEV_MODE")
	if !ok || devMode == "false" {
		kafkaApiKey, ok := os.LookupEnv("KAFKA_API_KEY")
		if !ok {
			log.Fatalf("KAFKA_API_KEY is required")
		}

		kafkaApiSecret, ok := os.LookupEnv("KAFKA_API_SECRET")
		if !ok {
			log.Fatalf("KAFKA_API_SECRET is required")
		}

		kafkaConfig = &kafka.ConfigMap{
			"bootstrap.servers": kafkaHost,
			"security.protocol": "SASL_SSL",
			"sasl.mechanisms":   "PLAIN",
			"sasl.username":     kafkaApiKey,
			"sasl.password":     kafkaApiSecret,
			"group.id":          "watch-" + uuid.New().String(),
			"auto.offset.reset": "latest",
		}
	} else {
		kafkaConfig = &kafka.ConfigMap{
			"bootstrap.servers": kafkaHost,
			"group.id":          "watch-" + uuid.New().String(),
			"auto.offset.reset": "latest",
		}
	}

	producer, err := stream.NewProducer(kafkaConfig)
	if err != nil {
		log.Printf("unable to initialize stream producer: %v\n", err)
	}

	consumer, err := stream.NewConsumer(kafkaConfig)
	if err != nil {
		log.Printf("unable to initialize stream consumer: %v\n", err)
	}

	err = consumer.Subscribe(stream.EventTopic, nil)
	if err != nil {
		log.Printf("unable to subscribe to events topics: %v\n", err)
	}

	serversCollection := database.Collection(datastore.Collections.Servers)
	defer func() {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		if err = client.Disconnect(ctx); err != nil {
			log.Fatalf("unable to gracefully disconnect from datastore: %v\n", err)
		}
	}()

	consumer.SubscribeTopics([]string{stream.RequestTopic, stream.EventTopic}, nil)

	run := true

	for run {
		ev := consumer.Poll(0)
		switch e := ev.(type) {
		case *kafka.Message:
			var event *serversv1.Event
			if err := json.Unmarshal(e.Value, &event); err != nil {
				fmt.Printf("unable to parse message to product: %v\n", err)
			}

			fmt.Printf("event: %v\n", event)
			topic := *e.TopicPartition.Topic
			fmt.Printf("topic: %v\nmessage: %+v\n", topic, string(e.Value))

			switch event.Type {
			case serversv1.Event_REQUEST:
				switch event.Action {
				case serversv1.Event_CREATE:
					var product *serversv1.Server
					if err := json.Unmarshal(event.Resource, &product); err != nil {
						fmt.Printf("unable to parse message to product: %v\n", err)
					}

					product.Status = serversv1.Server_PENDING
					product.CreatedAt = timestamppb.New(e.Timestamp)
					product.UpdatedAt = timestamppb.New(e.Timestamp)

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()

					_, err := serversCollection.InsertOne(ctx, product)
					if err != nil {
						if mongo.IsDuplicateKeyError(err) {
							fmt.Printf("product with name %s already exists\n", product.Name)
							// TODO: notify user of such error
						} else {
							fmt.Printf("unable to create product on datastore: %v\n", err)
						}
					} else {
						productBytes, err := json.Marshal(product)
						if err != nil {
							fmt.Printf("unable to parse product to json: %v\n", err)
						}

						outcomeEvent := &serversv1.Event{
							Name:     event.Name,
							Type:     serversv1.Event_EVENT,
							Action:   event.Action,
							Resource: productBytes,
						}

						eventBytes, err := json.Marshal(outcomeEvent)
						if err != nil {
							fmt.Printf("unable to parse event to json: %v\n", err)
						}

						if err = producer.Produce(&kafka.Message{
							TopicPartition: kafka.TopicPartition{Topic: &stream.EventTopic, Partition: kafka.PartitionAny},
							Value:          eventBytes,
							Key:            e.Key,
						}, nil); err != nil {
							log.Printf("unable to create product on datastore: %v\n", err)
							return
						}
					}

					if _, err = consumer.CommitMessage(e); err != nil {
						fmt.Printf("unable to commit stream consume: %v\n", err)
					}
				case serversv1.Event_UPDATE:
					var patch *serversv1.Patch
					if err := json.Unmarshal(event.Resource, &patch); err != nil {
						fmt.Printf("unable to parse message to product: %v\n", err)
					}

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()

					sets := bson.M{}
					for key, value := range patch.Changes {
						sets[key] = value
					}

					sets["updatedat"] = timestamppb.New(time.Now())

					res := serversCollection.FindOneAndUpdate(ctx, bson.M{"name": patch.Name}, bson.M{"$set": sets})
					if res.Err() != nil {
						if res.Err() == mongo.ErrNoDocuments {
							fmt.Printf("product already exists")
							// TODO: send error to notification service
						}
						fmt.Printf("unable to update product on datastore: %v\n", err)
						// TODO: send error to notification service
					} else {
						outcomeEvent := &serversv1.Event{
							Name:     event.Name,
							Type:     serversv1.Event_EVENT,
							Action:   event.Action,
							Resource: event.Resource,
						}

						eventBytes, err := json.Marshal(outcomeEvent)
						if err != nil {
							fmt.Printf("unable to parse event to json: %v\n", err)
						}

						if err = producer.Produce(&kafka.Message{
							TopicPartition: kafka.TopicPartition{Topic: &stream.EventTopic, Partition: kafka.PartitionAny},
							Value:          eventBytes,
							Key:            e.Key,
						}, nil); err != nil {
							fmt.Printf("unable to update product on datastore: %v\n", err)
							// TODO: send error to notification service
							return
						}
					}

					if _, err = consumer.CommitMessage(e); err != nil {
						fmt.Printf("unable to commit stream consume: %v\n", err)
					}
				case serversv1.Event_DELETE:
					var identity *serversv1.ResourceIdentity
					if err := json.Unmarshal(event.Resource, &identity); err != nil {
						fmt.Printf("unable to parse message to product: %v\n", err)
						// TODO: send error to notification service
					}

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()

					_, err := serversCollection.DeleteOne(ctx, bson.M{"name": identity.Name})
					if err != nil {
						fmt.Printf("unable to delete product on datastore: %v\n", err)
						// TODO: send error to notification service
					} else {
						outcomeEvent := &serversv1.Event{
							Name:     event.Name,
							Type:     serversv1.Event_EVENT,
							Action:   event.Action,
							Resource: event.Resource,
						}

						eventBytes, err := json.Marshal(outcomeEvent)
						if err != nil {
							fmt.Printf("unable to parse event to json: %v\n", err)
						}

						if err = producer.Produce(&kafka.Message{
							TopicPartition: kafka.TopicPartition{Topic: &stream.EventTopic, Partition: kafka.PartitionAny},
							Value:          eventBytes,
							Key:            e.Key,
						}, nil); err != nil {
							fmt.Printf("unable to delete product on datastore: %v\n", err)
							// TODO: send error to notification service
							return
						}
					}

					if _, err = consumer.CommitMessage(e); err != nil {
						fmt.Printf("unable to commit stream consume: %v\n", err)
					}
				default:
					fmt.Printf("action unknown: %s\n", event.Action.String())
					return
				}
			case serversv1.Event_EVENT:
				switch event.Action {
				case serversv1.Event_CREATE:
					var product *serversv1.Server
					if err := json.Unmarshal(event.Resource, &product); err != nil {
						fmt.Printf("unable to parse message to product: %v\n", err)
					}

					// TODO: do any validation to ensure the product can be changed to ready

					ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
					defer cancel()

					sets := bson.M{}

					sets["status"] = serversv1.Server_READY
					sets["updatedat"] = timestamppb.New(time.Now())

					res := serversCollection.FindOneAndUpdate(ctx, bson.M{"name": product.Name}, bson.M{"$set": sets})
					if res.Err() != nil {
						if res.Err() == mongo.ErrNoDocuments {
							fmt.Printf("product already exists")
							// TODO: send error to notification service
						}
						fmt.Printf("unable to update product on datastore: %v\n", err)
						// TODO: send error to notification service
					} else {
						patch, err := json.Marshal(sets)
						if err != nil {
							fmt.Printf("unable to parse update map to json: %v\n", err)
						}

						outcomeEvent := &serversv1.Event{
							Name:     event.Name,
							Type:     serversv1.Event_EVENT,
							Action:   serversv1.Event_UPDATE,
							Resource: patch,
						}

						eventBytes, err := json.Marshal(outcomeEvent)
						if err != nil {
							fmt.Printf("unable to parse event to json: %v\n", err)
						}

						if err = producer.Produce(&kafka.Message{
							TopicPartition: kafka.TopicPartition{Topic: &stream.EventTopic, Partition: kafka.PartitionAny},
							Value:          eventBytes,
							Key:            e.Key,
						}, nil); err != nil {
							fmt.Printf("unable to update product on datastore: %v\n", err)
							// TODO: send error to notification service
							return
						}
					}

					if _, err = consumer.CommitMessage(e); err != nil {
						fmt.Printf("unable to commit stream consume: %v\n", err)
					}
				case serversv1.Event_UPDATE:
				case serversv1.Event_DELETE:
				default:
					fmt.Printf("action unknown: %s\n", event.Action.String())
					return
				}
			default:
				fmt.Printf("type unknown: %s\n", event.Type.String())
				return
			}

		case kafka.PartitionEOF:
			fmt.Printf("%% Reached %v\n", e)
		case kafka.Error:
			fmt.Fprintf(os.Stderr, "%% Error: %v\n", e)
			run = false
		default:
			// fmt.Printf("Ignored %v\n", e)
		}
	}
}
