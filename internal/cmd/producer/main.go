package main

import (
	"bytes"
	"encoding/json"
	"log"
	"os"

	"github.com/hidromatologia-v2/models"
	"github.com/hidromatologia-v2/models/common/postgres"
	"github.com/hidromatologia-v2/models/tables"
	"github.com/memphisdev/memphis.go"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "messaging",
		Usage: "messaging microservice for the ResupplyOrg project",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:  "host",
				Value: "localhost",
				Usage: "Memphis host",
			},
			&cli.StringFlag{
				Name:  "username",
				Value: "root",
				Usage: "Memphis username",
			},
			&cli.StringFlag{
				Name:  "password",
				Value: "memphis",
				Usage: "Memphis password",
			},
			&cli.BoolFlag{
				Name:  "no-pwd",
				Value: false,
				Usage: "Ignore password",
			},
			&cli.StringFlag{
				Name:  "token",
				Value: "",
				Usage: "Memphis connection token",
			},
			&cli.StringFlag{
				Name:  "station",
				Value: "alerts",
				Usage: "Memphis station",
			},
			&cli.StringFlag{
				Name:  "producer",
				Value: "alerts",
				Usage: "Memphis producer name",
			},
		},
		Action: func(ctx *cli.Context) error {
			var opts []memphis.Option
			if !ctx.Bool("no-pwd") {
				opts = append(opts, memphis.Password(ctx.String("password")))
			}
			if token := ctx.String("token"); token != "" {
				opts = append(opts, memphis.ConnectionToken(token))
			}
			conn, err := memphis.Connect(
				ctx.String("host"),
				ctx.String("username"),
				opts...,
			)
			if err != nil {
				log.Fatal(err)
			}
			defer conn.Close()
			prod, pErr := conn.CreateProducer(ctx.String("station"), ctx.String("producer"))
			if pErr != nil {
				log.Fatal(pErr)
			}
			defer prod.Destroy()
			c := models.NewController(&models.Options{Database: postgres.NewDefault()})
			defer c.Close()
			u := tables.RandomUser()
			uErr := c.DB.Create(u).Error
			if uErr != nil {
				log.Fatal(uErr)
			}
			s := tables.RandomStation(u)
			sErr := c.DB.Create(s).Error
			if sErr != nil {
				log.Fatal(sErr)
			}
			sensor := s.Sensors[0]
			a := tables.RandomAlert(u, &sensor)
			*a.Condition = tables.Ge
			*a.Value = 10
			a.Enabled = &tables.True
			aErr := c.DB.Create(a).Error
			if aErr != nil {
				log.Fatal(aErr)
			}
			var buffer bytes.Buffer
			eErr := json.NewEncoder(&buffer).Encode(tables.SensorRegistry{
				SensorUUID: sensor.UUID,
				Value:      11,
			})
			if eErr != nil {
				log.Fatal(eErr)
			}
			prodErr := prod.Produce(buffer.Bytes())
			if pErr != nil {
				log.Fatal(prodErr)
			}
			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
