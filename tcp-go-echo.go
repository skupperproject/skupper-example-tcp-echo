package main

import (
	"github.com/skupperproject/skupper-example-tcp-echo/pkg/server"
)
func main() {
	server.Run("9090", make(chan interface{}))
}
