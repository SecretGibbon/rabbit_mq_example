package config

import "fmt"

func GetRabbitMQConnectionString(user, pass, url, port string) string {
	return fmt.Sprintf("amqp://%s:%s@%s:%s", user, pass, url, port)
}
