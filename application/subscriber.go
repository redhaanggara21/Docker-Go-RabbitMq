package application

import (
	"context"
	"time"

	"github.com/redhaanggara21/docker-go-rabbitmq/amqp"
	"github.com/redhaanggara21/docker-go-rabbitmq/messaging"
	"github.com/redhaanggara21/docker-go-rabbitmq/util"
)

type (
	subscriber struct {
		queueName string
		broker    messaging.Broker
	}
)

func (app *Application) NewSubscriberDaemon() util.Daemon {
	broker := amqp.NewAmqpBroker(app.rabbit)
	return &subscriber{
		app.queueName,
		broker,
	}
}

func (sub *subscriber) Start() error {
	err := sub.broker.Start()
	if err != nil {
		return err
	}

	_, err = sub.broker.CreateSubscription(sub.queueName, sub.queueName, "", true, 5, sub.handle)
	if err != nil {
		return err
	}

	return nil
}

func (sub *subscriber) Stop() error {
	return sub.broker.Stop()
}

func (sub *subscriber) handle(ctx context.Context, event messaging.Event) error {
	util.SessionLogger(ctx).Debugf("received : %s", string(event.GetBody()))
	time.Sleep(time.Second * 5)
	return nil
}
