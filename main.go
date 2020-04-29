package main

import (
	"fmt"
	"log"
	"net/rpc"

	"github.com/Antony-Ndungu/arithmetic-rpc-microservice/contract"
	"github.com/Antony-Ndungu/arithmetic-rpc-microservice/server"
)

const port = 9090

func main() {
	var response contract.Response
	request := contract.Request{X: 10, Y: 8}
	go server.Start()
	client, err := rpc.Dial("tcp", fmt.Sprintf("localhost:%v", port))
	if err != nil {
		log.Fatal(err)
	}
	err = client.Call("ArithmeticHandler.Add", &request, &response)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(response.Result)
}
