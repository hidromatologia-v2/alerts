package watcher

import (
	"encoding/json"
	"testing"
	"time"

	"bytes"

	"github.com/hidromatologia-v2/models/common/connection"
	"github.com/hidromatologia-v2/models/tables"
	"github.com/stretchr/testify/assert"
)

func testWatcher(t *testing.T, w *Watcher) {
	t.Run("Valid", func(tt *testing.T) {
		prod := connection.DefaultProducer(tt)
		defer prod.Destroy()
		u := tables.RandomUser()
		assert.Nil(tt, w.Controller.DB.Create(u).Error)
		s := tables.RandomStation(u)
		assert.Nil(tt, w.Controller.DB.Create(s).Error)
		sensor := s.Sensors[0]
		a := tables.RandomAlert(u, &sensor)
		*a.Condition = tables.Ge
		*a.Value = 10
		a.Enabled = &tables.True
		assert.Nil(tt, w.Controller.DB.Create(a).Error)
		var buffer bytes.Buffer
		assert.Nil(tt, json.NewEncoder(&buffer).Encode(tables.SensorRegistry{
			SensorUUID: sensor.UUID,
			Value:      11,
		}))
		assert.Nil(tt, prod.Produce(buffer.Bytes()))
		var alert tables.Alert
		tick := time.NewTicker(time.Millisecond)
		defer tick.Stop()
		for i := 0; i < 1000; i++ {
			if w.Controller.DB.
				Where("user_uuid = ?", u.UUID).
				Where("uuid = ?", a.UUID).
				Where("enabled = ?", false).
				First(&alert).Error == nil {
				break
			}
			<-tick.C
		}
		assert.Equal(tt, a.UUID, alert.UUID)
		assert.False(tt, *alert.Enabled)
	})
}

func TestWatcher(t *testing.T) {
	w := &Watcher{
		Controller:             connection.PostgresController(),
		SensorRegistryConsumer: connection.DefaultConsumer(t),
		MessageProducer:        connection.NewProducer(t, "messages"),
	}
	go w.Run()
	defer w.Close()
	testWatcher(t, w)
}
