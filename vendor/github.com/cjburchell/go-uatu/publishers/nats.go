package publishers

import (
	"github.com/cjburchell/go-uatu"
	"github.com/nats-io/go-nats"
)

type NatsSettings struct {
	URL      string
	Token    string
	Password string
	User     string
}

type natsPublisher struct {
	natsConn *nats.Conn
}

func (publisher natsPublisher) Publish(messageBites []byte) error {
	return publisher.natsConn.Publish("logs", messageBites)
}

func SetupNats(newSettings NatsSettings) log.Publisher {
	natsConn, err := nats.Connect(
		newSettings.URL,
		nats.Token(newSettings.Token),
		nats.UserInfo(newSettings.User, newSettings.Password),
		nats.DisconnectHandler(func(nc *nats.Conn) {
			log.Printf("Logger got disconnected\n")
		}),
		nats.ReconnectHandler(func(nc *nats.Conn) {
			log.Printf("Logger reconnected to %v\n", nc.ConnectedUrl())
		}),
		nats.ClosedHandler(func(nc *nats.Conn) {
			log.Printf("Logger connection closed. Reason: %q\n", nc.LastError())
		}))
	if err != nil {
		log.Printf("Can't connect: %v\n", err)
	}

	return natsPublisher{natsConn: natsConn}
}
