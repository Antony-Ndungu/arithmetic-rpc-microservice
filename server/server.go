package server

import (
	"fmt"
	"log"
	"net"
	"net/rpc"

	"github.com/Antony-Ndungu/arithmetic-rpc-microservice/contract"
)

const port = 9090

// ArithmeticHandler is struct with methods for doing arithmetic operations on tw integer operands
type ArithmeticHandler struct {
}

// Add performs arithmetic addition of the two parameters in a request struct and saves the result in the response struct result field
func (a *ArithmeticHandler) Add(args *contract.Request, reply *contract.Response) error {
	reply.Result = args.X + args.Y
	return nil
}

// Subtract performs arithmetic subtraction of the two parameters in a request struct and saves the result in the response struct result field
func (a *ArithmeticHandler) Subtract(args *contract.Request, reply *contract.Response) error {
	reply.Result = args.X + args.Y
	return nil
}

// Start starts an rpc server which a client can connect to and invoke the arithmetic methods defined in the server package
func Start() {
	rpc.Register(&ArithmeticHandler{})

	listener, err := net.Listen("tcp", fmt.Sprintf(":%v", port))
	if err != nil {
		log.Fatalf("Failed to listen on port %v: %v", port, err)
	}
	defer listener.Close()

	for {
		conn, _ := listener.Accept()
		go rpc.ServeConn(conn)
	}
}
