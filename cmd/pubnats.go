package main

import (
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"

	"io/ioutil"
	"os"
)

const (
	clusterID = "test-cluster"
	clientID  = "client2"
	subject   = "subject1"
)

func main() {

	MyLogger := logrus.New()

	MyLogger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: false,
		PrettyPrint:      true,
	}

	conn, err := stan.Connect(clusterID, clientID)
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "stan.Connect",
			"package": "main",
		}).Fatal(err)
	}
	defer conn.Close()

	file, err := os.Open("/home/ubuntu/ServerForWb/ServerWithNatsL0/model/model.json")
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "os.Open",
			"package": "main",
		}).Fatal(err)
	}
	defer file.Close()

	byteData, err := ioutil.ReadAll(file)
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "ioutil.ReadAll",
			"package": "main",
		}).Fatal(err)
	}

	err = conn.Publish(subject, byteData)
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "ioutil.ReadAll",
			"package": "main",
		}).Fatal(err)
	}

	MyLogger.Info("NATS Streaming, Publish Msg")
}
