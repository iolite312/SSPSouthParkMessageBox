package queue

import (
	"encoding/json"
	"fmt"
	"ssp-sp-messaging/api/internal/core/domain"
	"ssp-sp-messaging/api/internal/core/ports"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

type Publisher struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	queue   string
}

var _ ports.MessagePublisher = (*Publisher)(nil)

func NewPublisher(url, queue string) (*Publisher, error) {
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, fmt.Errorf("dial rabbitmq: %w", err)
	}

	ch, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("open channel: %w", err)
	}

	// durable queue so messages survive broker restarts
	if _, err := ch.QueueDeclare(queue, true, false, false, false, nil); err != nil {
		ch.Close()
		conn.Close()
		return nil, fmt.Errorf("declare queue: %w", err)
	}

	return &Publisher{conn: conn, channel: ch, queue: queue}, nil
}

func (p *Publisher) Publish(msg domain.Message) error {
	body, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	// default exchange ("") routes by queue name
	return p.channel.Publish("", p.queue, false, false, amqp.Publishing{
		ContentType:  "application/json",
		DeliveryMode: amqp.Persistent,
		MessageId:    msg.ID,
		Timestamp:    time.Now(),
		Body:         body,
	})
}

func (p *Publisher) Close() error {
	if err := p.channel.Close(); err != nil {
		return err
	}
	return p.conn.Close()
}
