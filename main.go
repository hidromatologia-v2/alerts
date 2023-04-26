package main

import (
	"context"
	"log"

	"github.com/hidromatologia-v2/alerts/watcher"
	"github.com/hidromatologia-v2/models"
	"github.com/hidromatologia-v2/models/common/postgres"
	"github.com/memphisdev/memphis.go"
	"github.com/sethvargo/go-envconfig"
)

func logFatalErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	var config Config
	eErr := envconfig.Process(context.Background(), &config)
	logFatalErr(eErr)
	controllerOpts := models.Options{
		Database: postgres.New(config.Postgres.DSN),
	}
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
	srCons, srErr := conn.CreateConsumer(
		config.Consumer.Station,
		config.Consumer.Consumer,
	)
	logFatalErr(srErr)
	msgProd, msgErr := conn.CreateProducer(
		config.Consumer.Station,
		config.Consumer.Consumer,
	)
	logFatalErr(msgErr)
	w := &watcher.Watcher{
		Controller:             models.NewController(&controllerOpts),
		SensorRegistryConsumer: srCons,
		MessageProducer:        msgProd,
	}
	rErr := w.Run()
	logFatalErr(rErr)
	<-make(chan struct{}, 1)
}
