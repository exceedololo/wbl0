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
	/*_, err := n.sc.Subscribe(n.channel, func(msg *stan.Msg) {
		var orderData Order
		err := json.Unmarshal(msg.Data, &orderData)
		if err != nil {
			log.Println("Error decoding message:", err)
			return
		}

		if n.handlerFunc != nil {
			err = n.handlerFunc(orderData)
			if err != nil {
				log.Println("Error processing message:", err)
				return
			}
		}
	})

	if err != nil {
		return err
	}

	return nil*/
}

// Closing connection wuth a NATS-stream
func (n *NATSConnection) Close() {
	if n.Conn != nil {
		n.Conn.Close()
		fmt.Println("closed NATS streaming connection")
	}
}
