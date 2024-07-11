package gomq

import "fmt"

func example() {
	err := Connect("rabbitmq_user", "rabbitmq_password", "103.28.52.86", "5672")
	if err != nil {
		panic(err)
	}
	Publish("hello", "Hello, RabbitMQ World!")
	Consume("hello", func(data string) {
		fmt.Println(data)
	})
}
