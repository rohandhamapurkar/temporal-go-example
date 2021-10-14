package app

import "fmt"



func ComposeGreeting(payload interface{}) (string, error) {

	value, ok := payload.(map[string]interface{})["name"].(string)
	if !ok {
		fmt.Println("Wrong assertion", payload)
		return "", nil
	}
	return fmt.Sprintf("Hello %s!", value), nil
}