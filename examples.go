package gomq

func example() {
	err := Connect("rabbitmq_user", "rabbitmq_password", "103.28.52.86", "5672")
	if err != nil {
		panic(err)
	}
	CreateChannel("hello")
	Publish("hello", "Hello, RabbitMQ World!")
	Consume("hello", nil)
}
