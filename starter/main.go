package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"rohandhamapurkar/temporal-hello-world/app"
	"time"

	"go.temporal.io/sdk/client"
)



func main() {
	// create temporal client
	c, err := client.NewClient(client.Options{})
	if(err != nil) {
		log.Fatalln(err)
	}

	// defer the client close
	defer c.Close()

	// workflow options
	options := client.StartWorkflowOptions{
		ID: fmt.Sprint(time.Microsecond.Seconds()),
		TaskQueue: "DEFAULT_TASK_QUEUE",
	}

	// Open our json file
	jsonFile, err := os.Open("config.json")
	// if we os.Open returns an error then handle it
	if err != nil {
		log.Fatalln(err)
	}
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	// read our opened jsonFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	var workflow_config []app.WorkflowConfig
	json.Unmarshal(byteValue, &workflow_config)


	we,err := c.ExecuteWorkflow(context.Background(), options, app.ModularWorkflow, workflow_config)

	if err != nil {
		log.Fatalln("unable to complete Workflow", err)
	}
	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}
	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}