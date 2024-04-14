package main

import (
	"fmt"
	"github/com/marktran77/distributed-file-storage/p2p"
	"log"
)

func main() {
	tr := p2p.NewTCPTransport(":4000")
	if err := tr.ListenAndAccept(); err != nil {
		log.Fatal(err)
	}
	select {}
	fmt.Println("We Gucci!")
}
