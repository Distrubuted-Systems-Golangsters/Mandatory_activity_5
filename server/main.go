package main

import (
	serverPackage "auction/server/internal"
	"fmt"
	"os"
)

func main() {
	args := os.Args[1:]

	if len(args) < 2 {
		fmt.Println("Arguments required: 'id' and 'port'")
		os.Exit(1)
	}

	id, port := args[0], args[1]

	if (id != "1" && id != "2" && id != "3") || (port != "8080" && port != "8081" && port != "8082") {
		fmt.Println("Invalid id or port")
		os.Exit(1)
	}

	if (id == "1" && port != "8080") || (id != "2" && port == "8081") || (id != "3" && port == "8082") {
		fmt.Println("Invalid id or port combination")
		os.Exit(1)
	}

	serverPackage.NewServer(id, port)

	//blocker
	bl := make(chan bool)
	<-bl
}