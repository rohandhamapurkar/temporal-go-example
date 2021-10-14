package app

import (
	"time"

	"go.temporal.io/sdk/workflow"
)


func ModularWorkflow(ctx workflow.Context, config []WorkflowConfig) (string,error) {
	options := workflow.ActivityOptions{
		ScheduleToCloseTimeout: time.Second * 60,
	}

	hash_map := map[string]interface{}{
		"ComposeGreeting": ComposeGreeting,
	}

	var result string

	for i := 0; i < len(config); i++ {
		ctx = workflow.WithActivityOptions(ctx,options)
		err := workflow.ExecuteActivity(ctx, hash_map[config[i].FuncName], config[i].Payload).Get(ctx, &result)

		if err != nil {
			return "",err
		}
    }
	
	return result, nil
}