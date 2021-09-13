package stream

import (
	"fmt"

	backoff "github.com/cenkalti/backoff/v4"
	"github.com/golang/glog"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	serversv1 "github.com/vniche/collective/protocol/v1"
)

const Resource = "SERVER"

var (
	RequestTopic = serversv1.Event_REQUEST.String() + "_" + Resource
	EventTopic   = serversv1.Event_EVENT.String() + "_" + Resource
)

func NewProducer(configMap *kafka.ConfigMap) (*kafka.Producer, error) {
	var producer *kafka.Producer
	initialize := func() error {
		var err error
		producer, err = kafka.NewProducer(configMap)
		if err != nil {
			fmt.Printf("unable to initialize stream producer: %s\n", err.Error())
		}

		return err
	}

	if err := backoff.Retry(initialize, backoff.NewExponentialBackOff()); err != nil {
		return nil, err
	}

	glog.V(2).Infoln("connected to kafka")

	return producer, nil
}

func NewConsumer(configMap *kafka.ConfigMap) (*kafka.Consumer, error) {
	var consumer *kafka.Consumer
	initialize := func() error {
		var err error
		consumer, err = kafka.NewConsumer(configMap)
		if err != nil {
			fmt.Printf("unable to initialize stream consumer: %s\n", err.Error())
		}

		return err
	}

	if err := backoff.Retry(initialize, backoff.NewExponentialBackOff()); err != nil {
		return nil, err
	}

	glog.V(2).Infoln("connected to kafka")

	return consumer, nil
}
