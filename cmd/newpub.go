package main

import (
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"
)

const (
	clientID  = "client2"
	clusterID = "test-cluster"
	subject   = "subject1"
)

func main() {

	MyLogger := logrus.New()

	MyLogger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: false,
		PrettyPrint:      true,
	}

	conn, err := stan.Connect(clusterID, clientID) //url
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "stan.Connect",
			"package": "main",
		}).Fatal(err)
	}
	defer conn.Close()

	SliceJsonString := SliceJsonString()

	for _, data := range SliceJsonString {

		bytesData := []byte(data)

		err = conn.Publish(subject, bytesData)
		if err != nil {
			MyLogger.WithFields(logrus.Fields{
				"func":    "stan.Connect",
				"package": "main",
			}).Fatal(err)
		}
	}

	MyLogger.Info("NATS Streaming, Publish Msg")
}

func SliceJsonString() []string {

	var SliceJsonString []string

	j1 := `
{
  "order_uid": "2_orderuid",
  "track_number": "2_track_number",
  "entry": "2_entry",
  "delivery": {
    "name": "2_name",
    "phone": "2_phone",
    "zip": "2_zip",
    "city": "2_city",
    "address": "2_address",
    "region": "2_region",
    "email": "2_email"
  },
  "payment": {
    "transaction": "2_transaction",
    "request_id": "",
    "currency": "2_currency",
    "provider": "2_provider",
    "amount": 2,
    "payment_dt": 2,
    "bank": "2_bank",
    "delivery_cost": 2,
    "goods_total": 2,
    "custom_fee": 2
  },
  "items": [
    {
      "chrt_id": 21,
      "track_number": "21",
      "price": 21,
      "rid": "21_rid",
      "name": "21_name",
      "sale": 21,
      "size": "",
      "total_price": 21,
      "nm_id": 21,
      "brand": "21_brand",
      "status": 21
    },
    {
      "chrt_id": 22,
      "track_number": "22_track_number",
      "price": 22,
      "rid": "22_rid",
      "name": "22_name",
      "sale": 22,
      "size": "",
      "total_price": 22,
      "nm_id": 22,
      "brand": "22_brand",
      "status": 22
    }
  ],
  "locale": "2_locale",
  "internal_signature": "",
  "customer_id": "2_customer_id",
  "delivery_service": "2_delivery_service",
  "shardkey": "2_shardkey",
  "sm_id": 2,
  "date_created": "2_date_created",
  "oof_shard": "2_oof_shard"
}
`
	j2 := "sdfghdfghjkcvbnm"
	j3 := `{"locale": "2_locale", "nm_id": 21}`
	SliceJsonString = append(SliceJsonString, j1)
	SliceJsonString = append(SliceJsonString, j2)
	SliceJsonString = append(SliceJsonString, j3)

	return SliceJsonString
}
