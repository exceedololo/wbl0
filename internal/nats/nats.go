package nats

import (
	"fmt"
	"github.com/nats-io/stan.go"
)

type NATSConnection struct {
	Conn stan.Conn
}

// establishing new connection with a NATS-stream
func NewConnection(clusterID, clientID, natsURL string) (*NATSConnection, error) {
	conn, err := stan.Connect(clusterID, clientID, stan.NatsURL(natsURL))
	if err != nil {
		return nil, err
	}
	fmt.Println("Connected to NATS Streaming")
	return &NATSConnection{Conn: conn}, nil
}

func (n *NATSConnection) Subscribe(channelName string, callback func([]byte)) (stan.Subscription, error) {
	return n.Conn.Subscribe(channelName, func(msg *stan.Msg) {
		// Вызываем функцию обратного вызова и передаем ей данные из сообщения
		callback(msg.Data)
	})
}

// Closing connection wuth a NATS-stream
func (n *NATSConnection) Close() {
	if n.Conn != nil {
		n.Conn.Close()
		fmt.Println("closed NATS streaming connection")
	}
}
