package main

import (
	"log"

	"go.temporal.io/sdk/client"
	"go.temporal.io/sdk/worker"

	"rohandhamapurkar/temporal-hello-world/app"
)

func main(){
	// create temporal client at worker side
	c,err := client.NewClient(client.Options{})

	if(err != nil) {
		log.Fatalln("unable to create Temporal client", err)
	}

	defer c.Close()

	// register workflows and activities
	w := worker.New(c,"DEFAULT_TASK_QUEUE",worker.Options{})
	w.RegisterWorkflow(app.ModularWorkflow)
	w.RegisterActivity(app.ComposeGreeting)

	// to run the worker
	err = w.Run(worker.InterruptCh())

	if(err != nil) {
		log.Fatalln("unable to start worker", err)
	}
}