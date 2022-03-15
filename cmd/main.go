package main

import (
	"ServerWithNatsL0/app"
	"ServerWithNatsL0/config"
	"ServerWithNatsL0/database"
	"ServerWithNatsL0/model"
	"ServerWithNatsL0/servnats"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/nats-io/stan.go"
	"github.com/sirupsen/logrus"

	"context"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	//Creating a new logger object
	MyLogger := logrus.New()
	//Setting the log output format
	MyLogger.Formatter = &logrus.JSONFormatter{
		DisableTimestamp: false,
		PrettyPrint:      true,
	}
	//Getting the Config
	config, err := config.Get()
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "config.Get",
			"package": "main",
		}).Fatal(err)
	}
	//Connecting to the BD
	db, err := database.NewPostgresDB(database.Config{
		User:     config.DB.User,
		Host:     config.DB.Host,
		Password: config.DB.Password,
		Port:     config.DB.Port,
		DBName:   config.DB.Dbname,
		SSLMode:  config.DB.Sslmode,
	})
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "database.NewPostgresDB",
			"package": "main",
		}).Fatal(err)
	}
	//Creating a buffered channel to receive data from nats streaming and send data to the server.
	chanData := make(chan model.OrdersData, 1)
	//Creating a cache in the form map
	cache := make(map[string]model.OrdersData)
	//Creating App structure to communication to other functions.
	app := app.NewApp(db, MyLogger, config, chanData, cache)
	//Connect will form a connection to the NUTS Streaming.
	conn, err := stan.Connect(config.Nats.ClusterID, config.Nats.ClientID)
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "stan.Connect",
			"package": "main",
		}).Fatal(err)
	}
	//Creating new Subscription.
	sub, err := servnats.NewSubscription(conn, config, app.ChanData)
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "servnats.NewSubscription",
			"package": "main",
		}).Fatal(err)
	}
	//NewRouter returns a new router instance.
	router := mux.NewRouter()
	//HandlerFunc registers the handler function.
	router.HandleFunc("/", app.HomePage).Methods("GET")
	router.HandleFunc("/ordersdata/{orderuid}", app.PublishOrdersData).Methods("GET")
	//Server defines parameters for running an HTTP server.
	srv := &http.Server{
		Addr:         config.Host + ":" + config.Port,
		Handler:      router,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}
	//ListenAndServe listens on the TCP network address.
	go func() {
		err = srv.ListenAndServe()
		switch err {
		case http.ErrServerClosed:
			MyLogger.Info("Server at :8080 port Stopped")
		default:
			MyLogger.WithFields(logrus.Fields{
				"func":    "srv.ListenAndServe",
				"package": "main",
			}).Fatal(err)
		}
	}()

	MyLogger.Info("Server at :8080 port Start")
	//Creating a context with a cancel function,
	//for a gorutina receiving data from nats streaming.
	ctx1, cansel1 := context.WithCancel(context.Background())

	go func(ctx context.Context, ch chan model.OrdersData, CacheMap map[string]model.OrdersData) {
		//Getting data from the channel. We save it to the Cache and DB.
		for {
			OrdersData := <-ch

			_, ok := CacheMap[OrdersData.OrderUid]
			if !ok {
				CacheMap[OrdersData.OrderUid] = OrdersData
			}

			err := database.SaveDataInDB(app.Db, OrdersData)
			if err != nil {
				MyLogger.WithFields(logrus.Fields{
					"func":    "database.SaveDataInDB",
					"package": "main",
				}).Info(err)
			}
		}

	}(ctx1, app.ChanData, app.CacheMap)
	//Creating a channel to receive a signal.
	signalChanel := make(chan os.Signal, 1)
	//Notify causes package signal to relay incoming signals to signalChanel.
	signal.Notify(signalChanel, syscall.SIGTERM, syscall.SIGINT)

	<-signalChanel

	MyLogger.Info("server at :8080 port Shutting Down")

	ctx2, cancel2 := context.WithTimeout(context.Background(), 15*time.Second)
	//Shutdown gracefully shuts down the server without interrupting any active connections.
	err = srv.Shutdown(ctx2)
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "srv.Shutdown",
			"package": "main",
		}).Fatal(err)
	}
	//Unsubscribe
	sub.Unsubscribe()
	//Cancel1 stops the execution of the function.
	cansel1()

	close(chanData)
	//Closing the connection to the BD.
	err = db.Close()
	if err != nil {
		MyLogger.WithFields(logrus.Fields{
			"func":    "db.Close",
			"package": "main",
		}).Fatal(err)
	}
	//Cancel1 stops the execution of the function.
	cancel2()
}
