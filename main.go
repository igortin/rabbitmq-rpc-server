package main

import (
	"github.com/streadway/amqp"
	"log"
)

var (
	url      = "amqp://guest:guest@192.168.122.102:5672/"
	queueReq = "Q_rpc"
)

func main() {
	conn, err := NewConnect(url)
	defer conn.Close()

	ch1 := NewChannel(conn)
	defer ch1.Close()
	ch2 := NewChannel(conn)
	defer ch2.Close()


	q1 := Queue{
		queueReq,
		false,
		false,
		true,
		false,
		nil,
	}
	queue, err := CreateQueue(ch1, q1)
	if err != nil {
		log.Fatalf("%v", err)
	}

	err = ch1.Qos(
		1,     // prefetch count
		0,     // prefetch size
		false, // global
	)

	consumer := Consumer{
		queue.Name,
		"",
		true, // acknowledged for every consume msg
		false,
		false,
		false,
		nil,
	}
	msgs, err := CreateAmqpChannel(ch1, consumer)
	if err != nil {
		log.Fatalf("%v", err)
	}

	forever := make(chan bool)
	defer close(forever)
	go func() {
		for d := range msgs {
			response := ProccessMsg(d.Body)
			err = ch2.Publish(
				"",        // exchange
				d.ReplyTo, // routing key
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType:   "text/plain",
					CorrelationId: d.CorrelationId,
					Body:          response,
				})
		}
	}()
	<-forever
}
