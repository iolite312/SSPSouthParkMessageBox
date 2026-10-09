package config

import "os"

type Config struct {
	HTTPAddr    string
	RabbitMQURL string
	Queue       string
}

func Load() Config {
	return Config{
		HTTPAddr:    getenv("HTTP_ADDR", ":8080"),
		RabbitMQURL: getenv("RABBITMQ_URL", "amqp://guest:guest@localhost:5672/"),
		Queue:       getenv("RABBITMQ_QUEUE", "messages"),
	}
}

func getenv(key, fallback string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return fallback
}
