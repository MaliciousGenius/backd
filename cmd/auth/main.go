package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/backd-io/backd/internal/db"
	"github.com/backd-io/backd/internal/instrumentation"
	"github.com/backd-io/backd/internal/rest"
	"google.golang.org/grpc"
)

func main() {

	var (
		routes map[string]map[string]rest.APIEndpoint

		server   *rest.REST
		conn     *grpc.ClientConn
		inst     *instrumentation.Instrumentation
		mongo    *db.Mongo
		api      *apiStruct
		mongoURL string
		err      error
	)

	mongoURL = os.Getenv("MONGO_URL")
	if mongoURL == "" {
		fmt.Println("MONGO_URL not found")
		os.Exit(1)
	}

	// TODO: REMOVE! AND CONFIGURE PROPERLY
	address := "sessions:8082"

	// Set up a connection to the sessions server.
	conn, err = grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	mongo, err = db.NewMongo(mongoURL)
	er(err)

	inst, err = instrumentation.New("0.0.0.0:8183", true)
	er(err)

	api = &apiStruct{
		inst:     inst,
		mongo:    mongo,
		sessions: conn,
	}

	routes = map[string]map[string]rest.APIEndpoint{
		"GET": {
			"/session": {
				Handler: api.getSession,
			},
			"/me": {
				Handler: api.getMe,
			},
		},
		"POST": {
			"/session": {
				Handler: api.postSession,
			},
		},
		"DELETE": {
			"/session": {
				Handler: api.deleteSession,
			},
		},
	}

	server = rest.New("0.0.0.0:8083")
	server.SetupRouter(routes, inst)

	// graceful
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt)

	go func() {
		if err := inst.Start(); err != nil {
			inst.Error(err.Error())
		}
	}()

	go func() {
		if err = server.Start(); err != nil {
			inst.Error(err.Error())
		}
	}()

	<-stop

	inst.Info("Shutting down the server.")

	if err = inst.Shutdown(); err != nil {
		inst.Info(err.Error())
	}

	if err = server.Shutdown(); err != nil {
		inst.Info(err.Error())
	}

}

func er(err error) {
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
