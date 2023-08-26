package main

import (
	"context"
	"log"

	"github.com/hidromatologia-v2/alerts/watcher"
	"github.com/hidromatologia-v2/models"
	"github.com/hidromatologia-v2/models/common/postgres"
	"github.com/memphisdev/memphis.go"
	uuid "github.com/satori/go.uuid"
	"github.com/sethvargo/go-envconfig"
)

func logFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func newConsumer(config *Config) *memphis.Consumer {
	var connOpts []memphis.Option
	if config.Consumer.Password != nil {
		connOpts = append(connOpts, memphis.Password(*config.Consumer.Password))
	}
	if config.Consumer.ConnectionToken != nil {
		connOpts = append(connOpts, memphis.ConnectionToken(*config.Consumer.ConnectionToken))
	}
	conn, connErr := memphis.Connect(
		config.Consumer.Host,
		config.Consumer.Username,
		connOpts...,
	)
	logFatalErr(connErr)
	consumer, consumerErr := conn.CreateConsumer(
		config.Consumer.Station,
		config.Consumer.Consumer+uuid.NewV4().String(),
	)
	logFatalErr(consumerErr)
	return consumer
}

func newProducer(config *Config) *memphis.Producer {
	var connOpts []memphis.Option
	if config.Producer.Password != nil {
		connOpts = append(connOpts, memphis.Password(*config.Producer.Password))
	}
	if config.Producer.ConnectionToken != nil {
		connOpts = append(connOpts, memphis.ConnectionToken(*config.Producer.ConnectionToken))
	}
	conn, connErr := memphis.Connect(
		config.Producer.Host,
		config.Producer.Username,
		connOpts...,
	)
	logFatalErr(connErr)
	producer, producerErr := conn.CreateProducer(
		config.Producer.Station,
		config.Producer.Producer+uuid.NewV4().String(),
	)
	logFatalErr(producerErr)
	return producer
}

func main() {
	var config Config
	eErr := envconfig.Process(context.Background(), &config)
	logFatalErr(eErr)
	controllerOpts := models.Options{
		Database: postgres.New(config.Postgres.DSN),
	}
	srCons := newConsumer(&config)
	msgProd := newProducer(&config)
	w := &watcher.Watcher{
		Controller:             models.NewController(&controllerOpts),
		SensorRegistryConsumer: srCons,
		MessageProducer:        msgProd,
	}
	rErr := w.Run()
	logFatalErr(rErr)
	<-make(chan struct{}, 1)
}
