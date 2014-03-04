package main

import (
	"net"
	"definitions/arithmetic"
)

func main() {
	// Create a server.
	server := arithmetic.NewArithmeticServiceServer(new(arithmeticServiceImplementation))

	// Create a listener.
	listener, err := net.Listen("tcp", ":5555")
	if err != nil {
		panic(err)
	}

	// Serve indefinitely.
	server.Serve(listener)
}
