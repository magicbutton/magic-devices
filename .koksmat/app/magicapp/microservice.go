// -------------------------------------------------------------------
// Generated by 365admin-publish
// Service wrapper  v1
// -------------------------------------------------------------------
/*
---
title: List Groups
---
*/
package magicapp

import (
	"context"
	"log"
	"time"

	_ "github.com/lib/pq"
	nats "github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/micro"
	"github.com/spf13/viper"
)

func StartMicroService() {
	// Parent context cancels connecting/reconnecting altogether.
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	var err error
	var nc *nats.Conn
	opts := []nats.Option{

		nats.ReconnectWait(2 * time.Second),
		nats.ReconnectHandler(func(c *nats.Conn) {
			log.Println("Reconnected to", c.ConnectedUrl())
		}),
		nats.DisconnectHandler(func(c *nats.Conn) {
			log.Println("Disconnected from NATS")
		}),
		nats.ClosedHandler(func(c *nats.Conn) {
			log.Println("NATS connection is closed.")
		}),
	}

	go func() {
		natsServer := viper.GetString("NATS")
		if natsServer == "" {
			natsServer = "nats://127.0.0.1:4222"
		}
		log.Println("Connecting to", natsServer)
		nc, err = nats.Connect(natsServer, opts...)
	}()
	retryCount := 0
WaitForEstablishedConnection:
	for {
		if err != nil {
			log.Fatal(err)
		}
		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()
		// Wait for context to be canceled either by timeout
		// or because of establishing a connection...
		select {
		case <-ctx.Done():
			break WaitForEstablishedConnection
		default:
		}

		if nc == nil || !nc.IsConnected() {
			if retryCount != 0 {
				log.Println("Connection not ready")
			}
			retryCount++
			time.Sleep(200 * time.Millisecond)
			continue
		}
		break WaitForEstablishedConnection
	}
	if ctx.Err() != nil {
		log.Fatal(ctx.Err())
	}
	log.Println("Connected")

	OpenDatabase()

	// utils.Setup("./.env")

	// dsn := viper.GetString("POSTGRES_DB")
	// sqldb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(dsn)))
	// db := bun.NewDB(sqldb, pgdialect.New())
	// utils.Db = db
	const name = "magic-devices"
	srv, err := micro.AddService(nc, micro.Config{
		Name: name,

		Version:     "0.0.1",
		Description: "Magic People",
	})
	root := srv.AddGroup(name)
	RegisterServiceEndpoints(root)

	for {
		if nc.IsClosed() {
			break
		}

		time.Sleep(1 * time.Second)
	}

	// Disconnect and flush pending messages
	if err := nc.Drain(); err != nil {
		log.Println(err)
	}
	log.Println("Disconnected")

}
