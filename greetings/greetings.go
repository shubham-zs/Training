package greetings

import "fmt"

func SayHello(name string) string {
	msg := fmt.Sprint("Hello %s", name)
	return msg
}
