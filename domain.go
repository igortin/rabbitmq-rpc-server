package main

import (
	"github.com/streadway/amqp"
)

func NewConnect(url string) (*amqp.Connection, error) {
	connection, err := amqp.Dial(url)
	return connection, err
}

func NewChannel(con *amqp.Connection) (*amqp.Channel) {
	ch,_ := con.Channel()
	return ch
}

func CreateQueue(ch *amqp.Channel, q Queue) (amqp.Queue, error) {
	return ch.QueueDeclare(
		q.Name,
		q.Durable,
		q.AutoDelete,
		q.Durable,
		q.NoWait,
		q.Args,
	)
}

func CreateAmqpChannel(ch *amqp.Channel,consumer Consumer) (<-chan amqp.Delivery,error) {
	return ch.Consume(
		consumer.queueName,
		consumer.Name,
		consumer.autoAck,
		consumer.exclusive,
		consumer.noLocal,
		consumer.noWait,
		consumer.args,
	)}


func CreateExchange(ch *amqp.Channel, e Exchange) error {
	return ch.ExchangeDeclare(
		e.Name,
		e.Kind,
		e.Durable,
		e.AutoDeleted,
		e.Internal,
		e.NoWait,
		e.Args,
	)}


func CreateBind(ch *amqp.Channel ,b Bind) error {
	return ch.QueueBind(
		b.queueName,
		b.routingKey,
		b.exchange,
		b.noWait,
		b.args,
	)}










type Exchange struct {
	Name        string
	Kind        string
	Durable     bool
	AutoDeleted bool
	Internal    bool
	NoWait      bool
	Args        amqp.Table
}

type Queue struct {
	Name       string
	Durable    bool
	AutoDelete bool
	Exclusive  bool
	NoWait     bool
	Args       amqp.Table
}

type Consumer struct {
	queueName      string
	Name  string
	autoAck   bool
	exclusive bool
	noLocal   bool
	noWait    bool
	args      amqp.Table
}

type Bind struct {
	queueName      string
	routingKey  string
	exchange   string
	noWait    bool
	args      amqp.Table
}

type Message struct {
	exchangeName string
	routingKey string
	mandatory bool
	immediate bool
	Publish amqp.Publishing
}

func (msg *Message) Send(ch *amqp.Channel) error {
	return ch.Publish(
		msg.exchangeName,
		msg.routingKey,
		msg.mandatory,
		msg.immediate,
		msg.Publish,
		)}