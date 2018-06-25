package queuelib

import (
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

// Holds queue config
type Config struct {
	Username     string `json:"username,omitempty"`
	Password     string `json:"password,omitempty"`
	Host         string `json:"host,omitempty"`
	Vhost        string `json:"vhost,omitempty"`
	Queuename    string `json:"queuename,omitempty"`
	Exchange     string `json:"exchange,omitempty"`
	ExchangeType string `json:"exchange,omitempty"`
	Routingkey   string `json:"routingkey,omitempty"`
}

type Handler struct {
	conn *amqp.Connection
}

var Conn *Handler

// To establish the connection with the queue
func New(cfg *Config) (*Handler, error) {
	handler := new(Handler)
	uri := fmt.Sprintf(`amqp://%s`, cfg.Username)
	uri += fmt.Sprintf(`:%s`, cfg.Password)
	uri += fmt.Sprintf(`@%s:5672`, cfg.Host)
	if cfg.Vhost != "/" {
		uri += fmt.Sprintf(`/%s`, cfg.Vhost)
	}

	conn, err := amqp.Dial(uri)
	if err != nil {
		return nil, err
	}
	handler.conn = conn
	return handler, nil

}

func Init(cfg *Config) error {
	conn, err := New(cfg)
	if err != nil {
		return err
	}
	Conn = conn
	return nil
}

// To publish
func (h *Handler) Publish(cfg Config, body string) error {
	if h != nil {
		ch, err := h.conn.Channel()
		if err != nil {
			return err
		}

		// Check if the exchange is non-default
		if len(cfg.Exchange) > 0 {
			// Make sure the exchange exists by declaring it passively
			if err := ch.ExchangeDeclarePassive(
				cfg.Exchange,     // name
				cfg.ExchangeType, // type
				false,            // durable
				false,            // auto-deleted
				false,            // internal
				false,            // noWait
				nil,              // arguments
			); err != nil {
				return errors.New("The queue exchange doesn't exist")
			}
		}
		err = ch.Publish(
			cfg.Exchange,   // exchange
			cfg.Routingkey, // routing key
			false,          // mandatory
			false,          // immediate
			amqp.Publishing{
				ContentType: "text/plain",
				Body:        []byte(body),
			})
		if err != nil {
			return err
		}

		return nil
	}
	return errors.New("queue: Connect not initialised")
}
