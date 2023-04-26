package main

import "github.com/hidromatologia-v2/models/common/config"

type (
	Config struct {
		config.Consumer `env:",prefix=MEMPHIS_CONSUMER_"` // Memphis
		config.Producer `env:",prefix=MEMPHIS_PRODUCER_"` // Memphis
		config.Postgres `env:",prefix=POSTGRES_"`         // POSTGRESQL
	}
)
