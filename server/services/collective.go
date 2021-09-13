package services

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/golang/glog"
	"github.com/google/uuid"
	serversv1 "github.com/vniche/collective/protocol/v1"
	"github.com/vniche/collective/server/datastore"
	"github.com/vniche/collective/server/stream"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

var service *Service

type Service struct {
	server        *grpc.Server
	kafkaProducer *kafka.Producer
	kafkaConsumer *kafka.Consumer
	mongoDatabase *mongo.Database
}

func (service *Service) Shutdown() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var err error
	if err = service.mongoDatabase.Client().Disconnect(ctx); err != nil {
		glog.Errorf("unable to gracefully disconnect from datastore: %v\n", err)
	}

	service.kafkaProducer.Close()
	if err = service.kafkaConsumer.Close(); err != nil {
		glog.Errorf("unable to gracefully close kafka consumer: %v\n", err)
	}

	stream.Shutdown()

	return nil
}

// server is used to implement me.vniche.collective.v1.Servers
type server struct {
	serversv1.UnimplementedServersServer
}

func (s *server) Get(ctx context.Context, resourceIdentity *serversv1.ResourceIdentity) (*serversv1.Server, error) {
	if resourceIdentity.Name == "" {
		return nil, status.Errorf(codes.Internal, "name is required")
	}

	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if service.mongoDatabase == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to get server")
	}

	res := service.mongoDatabase.Collection(datastore.Collections.Servers).FindOne(ctx, bson.M{"name": resourceIdentity.Name})
	if res.Err() != nil {
		err := status.Errorf(codes.Internal, "unable to get server")
		if res.Err() == mongo.ErrNoDocuments {
			err = status.Errorf(codes.NotFound, "no server found with name %s", resourceIdentity.Name)
		}
		return nil, err
	}

	var (
		server *serversv1.Server
		err    error
	)
	err = res.Decode(&server)

	return server, err
}

func (s *server) List(listFilter *serversv1.ListFilter, stream serversv1.Servers_ListServer) error {
	var (
		err      error
		skip     int32
		pageSize int32
	)

	if listFilter.PageNumber > 0 {
		skip = (listFilter.PageNumber - 1) * listFilter.PageSize
	}

	if listFilter.PageSize == 0 {
		pageSize = 10
	}

	options := options.Find()
	options.SetSort(bson.M{"updatedat": -1})
	options.SetSkip(int64(skip))
	options.SetLimit(int64(pageSize))

	ctx, cancel := context.WithTimeout(stream.Context(), 1*time.Second)
	defer cancel()

	if service.mongoDatabase == nil {
		return status.Errorf(codes.FailedPrecondition, "unable to list servers")
	}

	cursor, err := service.mongoDatabase.Collection(datastore.Collections.Servers).Find(ctx, bson.M{}, options)
	if err != nil {
		return status.Errorf(codes.Internal, "unable to list servers")
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {

		var server *serversv1.Server
		if err = cursor.Decode(&server); err != nil {
			return status.Errorf(codes.Internal, "unable to list servers")
		}

		if err = stream.Send(server); err != nil {
			return status.Errorf(codes.Internal, "unable to list servers")
		}
	}

	if err := cursor.Err(); err != nil {
		status.Errorf(codes.Internal, "unable to list servers")
	}

	return nil
}

func (s *server) Create(ctx context.Context, server *serversv1.Server) (*serversv1.ResourceChange, error) {
	if server.Name == "" {
		return nil, status.Errorf(codes.InvalidArgument, "name is required")
	}

	serverBytes, err := json.Marshal(server)
	if err != nil {
		glog.Errorln(err)
		return nil, status.Errorf(codes.Aborted, "unable to request server creation")
	}

	eventBytes, err := json.Marshal(&serversv1.Event{
		Name:     server.Name,
		Type:     serversv1.Event_REQUEST,
		Action:   serversv1.Event_CREATE,
		Resource: serverBytes,
	})
	if err != nil {
		glog.Errorln(err)
		return nil, status.Errorf(codes.Aborted, "unable to request server creation")
	}

	if service.kafkaProducer == nil {
		glog.Errorln(fmt.Errorf("kafka producer unavailable"))
		return nil, status.Errorf(codes.FailedPrecondition, "unable to request server creation")
	}

	if err = service.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &stream.RequestTopic, Partition: kafka.PartitionAny},
		Value:          eventBytes,
		Key:            []byte(uuid.New().String()),
	}, nil); err != nil {
		return nil, status.Errorf(codes.Aborted, "unable to request server creation")
	}

	return &serversv1.ResourceChange{
		Message: "creation successfully requested",
	}, nil
}

func (s *server) Update(ctx context.Context, patch *serversv1.Patch) (*serversv1.ResourceChange, error) {
	if patch.Name == "" {
		return nil, status.Errorf(codes.Internal, "name is required")
	}

	if len(patch.Changes) == 0 {
		return nil, status.Errorf(codes.Internal, "at least one change is required")
	}

	patchBytes, err := json.Marshal(patch)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "unable to request server update")
	}

	eventBytes, err := json.Marshal(&serversv1.Event{
		Name:     patch.Name,
		Type:     serversv1.Event_REQUEST,
		Action:   serversv1.Event_UPDATE,
		Resource: patchBytes,
	})
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "unable to request server update")
	}

	if service.kafkaProducer == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to request server update")
	}

	if err = service.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &stream.RequestTopic, Partition: kafka.PartitionAny},
		Value:          eventBytes,
		Key:            []byte(uuid.New().String()),
	}, nil); err != nil {
		return nil, status.Errorf(codes.Aborted, "unable to request server update")
	}

	return &serversv1.ResourceChange{
		Message: "update successfully requested",
	}, nil
}

func (s *server) Delete(ctx context.Context, resourceIdentity *serversv1.ResourceIdentity) (*serversv1.ResourceChange, error) {
	if resourceIdentity.Name == "" {
		return nil, status.Errorf(codes.Internal, "name is required")
	}

	identityBytes, err := json.Marshal(resourceIdentity)
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "unable to patch resource")
	}

	eventBytes, err := json.Marshal(&serversv1.Event{
		Name:     resourceIdentity.Name,
		Type:     serversv1.Event_REQUEST,
		Action:   serversv1.Event_DELETE,
		Resource: identityBytes,
	})
	if err != nil {
		return nil, status.Errorf(codes.Aborted, "unable to request server creation")
	}

	if service.kafkaProducer == nil {
		return nil, status.Errorf(codes.FailedPrecondition, "unable to request server creation")
	}

	if err = service.kafkaProducer.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &stream.RequestTopic, Partition: kafka.PartitionAny},
		Value:          eventBytes,
		Key:            []byte(uuid.New().String()),
	}, nil); err != nil {
		return nil, status.Errorf(codes.Aborted, "unable to request server deletion")
	}

	return &serversv1.ResourceChange{
		Message: "deletion successfully requested",
	}, nil
}

func (s *server) Watch(listFilter *serversv1.ResourceIdentity, watchStream serversv1.Servers_WatchServer) error {
	if service.kafkaConsumer == nil {
		return status.Errorf(codes.FailedPrecondition, "unable to request server creation")
	}

	clientStream := stream.New()

	for {
		select {
		case <-watchStream.Context().Done():
			fmt.Printf("close stream from client")
			clientStream.Close()
			return nil
		case event, ok := <-clientStream.Channel:
			if event == nil || !ok {
				return nil
			}

			watchStream.Send(event)
		}
	}
}

func New() *Service {
	s := grpc.NewServer()
	reflection.Register(s)
	serversv1.RegisterServersServer(s, &server{})

	service = &Service{
		server: s,
	}

	go func() {
		mongoDB, ok := os.LookupEnv("MONGO_DB")
		if !ok {
			mongoDB = "collective"
		}

		mongoURI, ok := os.LookupEnv("MONGO_URI")
		if !ok {
			mongoURI = "mongodb://localhost/" + mongoDB
		}

		mongoClient, err := datastore.Connect(mongoURI)
		if err != nil {
			log.Printf("unable to connect to datastore: %v\n", err)
		}

		service.mongoDatabase = mongoClient.Database(mongoDB)

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

		service.kafkaProducer, err = stream.NewProducer(kafkaConfig)
		if err != nil {
			log.Printf("unable to initialize stream producer: %v\n", err)
		}

		service.kafkaConsumer, err = stream.NewConsumer(kafkaConfig)
		if err != nil {
			log.Printf("unable to initialize stream consumer: %v\n", err)
		}

		err = service.kafkaConsumer.Subscribe(stream.EventTopic, nil)
		if err != nil {
			log.Printf("unable to subscribe to events topics: %v\n", err)
		}

		run := true

		for run {
			ev := service.kafkaConsumer.Poll(0)
			switch e := ev.(type) {
			case *kafka.Message:
				var event *serversv1.Event
				if err := json.Unmarshal(e.Value, &event); err != nil {
					fmt.Printf("unable to parse message to product: %v\n", err)
				}

				stream.PublishToAll(event)

				if _, err = service.kafkaConsumer.CommitMessage(e); err != nil {
					fmt.Printf("unable to commit stream consume: %v\n", err)
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
	}()

	return service
}
