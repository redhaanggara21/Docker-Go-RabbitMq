package application

import (
	"github.com/redhaanggara21/docker-go-rabbitmq/amqp"
	"github.com/redhaanggara21/docker-go-rabbitmq/util"
)

type (
	Application struct {
		queueName string
		rabbit    *amqp.RabbitConfig
	}

	Logger struct {
		// Stdout is true if the output needs to goto standard out
		Stdout bool `yaml:"stdout"`
		// Level is the desired log level
		Level string `yaml:"level"`
		// OutputFile is the path to the log output file
		OutputFile string `yaml:"outputFile"`
	}
)

func SetupApp() *Application {
	amqpConf := amqp.RabbitConfig{
		Host:     "localhost:5672",
		User:     "guest",
		Password: "guest",
	}

	logger := Logger{Stdout: true, Level: "DEBUG"}
	util.Log = logger.NewLogger()

	return &Application{
		"testing",
		&amqpConf,
	}
}
