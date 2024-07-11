package gomq

import (
	"errors"
	"fmt"

	"github.com/streadway/amqp"
)

var _qp *amqp.Connection
var _ch *amqp.Channel

func Connect(user, pass, host, port string) error {
	if _qp != nil {
		return nil
	}
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%v:%v@%v:%v/", user, pass, host, port))
	if err != nil {
		return err
	}
	_qp = conn
	return nil
}

func createChannel(name string) error {
	if _qp == nil {
		return errors.New("Disconnected")
	}

	if _ch == nil {
		ch, err := _qp.Channel()
		if err != nil {
			return err
		}
		_ch = ch
	}

	_, err := _ch.QueueDeclare(
		name,  // Queue name
		false, // Durable (messages survive broker restarts)
		false, // Delete when unused
		false, // Exclusive (for this connection only)
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return err
	}

	return nil
}

func Publish(name, data string) error {
	if _qp == nil {
		return errors.New("Disconnected")
	}
	if _ch == nil {
		createChannel(name)
	}
	return _ch.Publish(
		"",    // Exchange
		name,  // Routing key
		false, // Mandatory
		false, // Immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(data),
		})
}

func Consume(name string, action func(data string)) error {
	if _qp == nil {
		return errors.New("Disconnected")
	}
	if _ch == nil {
		createChannel(name)
	}

	msgs, err := _ch.Consume(
		name,  // Queue name
		"",    // Consumer
		true,  // Auto-acknowledgement
		false, // Exclusive
		false, // No-local
		false, // No-wait
		nil,   // Arguments
	)
	if err != nil {
		return err
	}

	for d := range msgs {
		action(string(d.Body))
	}
	return nil
}
