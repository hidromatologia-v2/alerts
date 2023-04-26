package watcher

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/hidromatologia-v2/models"
	"github.com/hidromatologia-v2/models/common/logs"
	"github.com/hidromatologia-v2/models/tables"
	"github.com/memphisdev/memphis.go"
)

type Watcher struct {
	Controller             *models.Controller
	SensorRegistryConsumer *memphis.Consumer
	MessageProducer        *memphis.Producer
}

func (w *Watcher) Close() {
	w.MessageProducer.Destroy()
	w.SensorRegistryConsumer.Destroy()
}

func (w *Watcher) HandleMessage(message *memphis.Msg) {
	defer func() {
		err, _ := recover().(error)
		if err == nil {
			return
		}
		logs.LogOnError(err)
	}()
	now := time.Now()
	var registry tables.SensorRegistry
	dErr := json.NewDecoder(bytes.NewReader(message.Data())).Decode(&registry)
	logs.PanicOnError(dErr)
	alertsTriggered, tErr := w.Controller.CheckAlert(&registry)
	logs.PanicOnError(tErr)
	ackErr := message.Ack()
	logs.PanicOnError(ackErr)
	for _, alert := range alertsTriggered {
		var alertBuffer bytes.Buffer
		eErr := json.NewEncoder(&alertBuffer).Encode(tables.Message{
			Type:      tables.Email,
			Recipient: *alert.User.Email,
			Subject:   fmt.Sprintf("Alert: %s", *alert.Name),
			Body:      fmt.Sprintf("Your alert was triggered at: %s", now.Format(time.ANSIC)),
		})
		logs.PanicOnError(eErr)
		pErr := w.MessageProducer.Produce(alertBuffer.Bytes())
		logs.LogOnError(pErr)
	}
}

func (w *Watcher) Run() error {
	return w.SensorRegistryConsumer.Consume(
		func(messages []*memphis.Msg, err error, ctx context.Context) {
			logs.LogOnError(err)
			for _, message := range messages {
				w.HandleMessage(message)
			}
		},
	)
}
