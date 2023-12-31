package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Rabbit struct {
	conn    *amqp.Connection
	errChan <-chan *amqp.Error
}

func newRabbit() (*Rabbit, error) {
	log.Println("Connecting to Rabbit...")

	conn, err := amqp.Dial("amqp://guest:guest@rabbit:5672")
	if err != nil {
		return nil, err
	}

	errs := make(chan *amqp.Error, 1)
	errChan := conn.NotifyClose(errs)

	// ...

	log.Println("Successfully connected to Rabbit")

	return &Rabbit{
		conn:    conn,
		errChan: errChan,
	}, nil
}

func (r *Rabbit) Err() <-chan *amqp.Error {
	return r.errChan
}

func (r *Rabbit) Disconnect(timeout time.Duration) error {
	if r.conn == nil {
		return nil
	}
	return r.conn.CloseDeadline(time.Now().Add(timeout))
}
