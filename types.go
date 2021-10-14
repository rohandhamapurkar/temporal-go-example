package app

type WorkflowConfig struct {
	FuncName string      `json:"funcName"`
	Payload  interface{} `json:"payload"`
}
