package hello

import "fmt"

// Say makes the greeting message.
//
// if parameter is empty, return empty.
func Say(message string) string {
	if message == "" {
		return ""
	}

	return makeMessage(message)
}

func makeMessage(message string) string {
	return fmt.Sprintf("hello %s", message)
}
