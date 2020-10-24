package main

import (
	"fmt"
	"golang-protobuffers-examples/messages/pb"
)

func main() {
	m := pb.Message{Body: "Hello World!"}

	fmt.Println(m.String())
}
