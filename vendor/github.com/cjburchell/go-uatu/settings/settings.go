package settings

import (
	"github.com/cjburchell/go-uatu"
	"github.com/cjburchell/go-uatu/publishers"
	"github.com/cjburchell/tools-go/env"
)

func createSettings() log.Settings {
	return log.Settings{
		ServiceName:  env.Get("LOG_SERVICE_NAME", ""),
		MinLogLevel:  log.GetLogLevel(env.Get("LOG_LEVEL", log.INFO.Text)),
		LogToConsole: env.GetBool("LOG_CONSOLE", true),
	}
}

func createHttpSettings() publishers.HttpSettings {
	return publishers.HttpSettings{
		Address: env.Get("LOG_REST_URL", "http://logger:8082/log"),
		Token:   env.Get("LOG_REST_TOKEN", "token"),
	}
}

func createNatsSettings() publishers.NatsSettings {
	return publishers.NatsSettings{
		URL:      env.Get("LOG_NATS_URL", "tcp://nats:4222"),
		Token:    env.Get("LOG_NATS_TOKEN", "token"),
		User:     env.Get("LOG_NATS_USER", "admin"),
		Password: env.Get("LOG_NATS_PASSWORD", "password"),
	}
}

func SetupLogger() error {
	newPublishers := make([]log.Publisher, 0)
	if env.GetBool("LOG_USE_NATS", true) {
		publisher := publishers.SetupNats(createNatsSettings())
		newPublishers = append(newPublishers, publisher)
	}

	if env.GetBool("LOG_USE_REST", false) {
		publisher := publishers.SetupHttp(createHttpSettings())
		newPublishers = append(newPublishers, publisher)
	}

	return log.Setup(createSettings(), newPublishers)
}
