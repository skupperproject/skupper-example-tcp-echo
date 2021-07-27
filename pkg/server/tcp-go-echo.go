package server

import (
	"fmt"
	"io"
	"net"
	"os"
	"strings"
)

var fp = fmt.Fprintf

// Handle a connection. Read messages from it and reply with the
// same message uppercased. Terminate the connection when we reach
// EOF,or an error.
func cnx_handler(cnx_number int,
	hostname string,
	cnx net.Conn) {
	buffer := make([]byte, 512)

	for {
		n, err := cnx.Read(buffer)

		if err != nil {
			if err != io.EOF {
				fp(os.Stderr, "Connection read error: |%s|\n", err.Error())
			}
			break
		}

		// If we have received a message, uppercase it and send it
		// back as the reply.
		if n > 0 {
			message := string(buffer[0:n])
			fp(os.Stdout, "received from cnx %d : |%s|\n", cnx_number, message)
			reply := hostname + " : " + strings.ToUpper(message)
			cnx.Write([]byte(reply + "\n"))
		}
	}
	cnx.Close()
	fp(os.Stdout, "TCP handler for connection %d exiting.\n", cnx_number)
}

func Run(port string, stopCh chan interface{}) {
	cnx_count := 0

	// Listen for TCP connections on any interface on this port.
	tcp_listener, _ := net.Listen("tcp", ":"+port)

	fp(os.Stdout, "tcp-echo server listening on port %s.\n", port)

	hostname := os.Getenv("HOSTNAME")

	// Handle each new connection in its own goroutine. Tell each
	// one what number it is so the user can see which handler is
	// printing out each message.
	for {
		select {
		case <-stopCh:
			break
		default:
			cnx, _ := tcp_listener.Accept()
			cnx_count++
			fp(os.Stdout, "server: made connection %d.\n", cnx_count)
			go cnx_handler(cnx_count, hostname, cnx)
		}
	}
}
