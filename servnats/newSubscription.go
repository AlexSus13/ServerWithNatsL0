package servnats

import (
	"ServerWithNatsL0/config"
	"ServerWithNatsL0/model"

	"github.com/nats-io/stan.go"
	"github.com/pkg/errors"

	"encoding/json"
	"log"
)

var chanal chan model.OrdersData

func handlerMsg(msg *stan.Msg) {

	var data model.OrdersData
	//Save the received data in the structure OrdersData
	err := json.Unmarshal(msg.Data, &data)
	if err != nil {
		log.Printf("Error when decoding Data from nats channel: %v ", err)
		msg.Ack()
		return
	}
	//Sending the structure to the channel
	chanal <- data
	//Manually sending an ACK to the server
	err = msg.Ack()
	if err != nil {
		log.Printf("Error when sending ACK: %d", msg.Sequence)
		return
	}
}

func NewSubscription(conn stan.Conn, conf *config.Conf, ch chan model.OrdersData) (stan.Subscription, error) {

	chanal = ch
	//Subscribe will perform a subscription with the given options to the cluster
	sub, err := conn.Subscribe(
		conf.Nats.Subject, //Getting the message subject from the Config
		handlerMsg,
		stan.DurableName("DurableSubscriptionsName"), //Parameter for a renewable subscription
		stan.SetManualAckMode(),                      //Option for automatic non-sending of ACK. You need to send it manually,
		//in case of an error in processing the received message, in order to resend it.
	)

	if err != nil {
		return nil, errors.Wrap(err, "func NewSubscription")
	}

	return sub, nil
}
