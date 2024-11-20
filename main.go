package main

import (
	"fmt"
	"time"

	"github.com/sher2001/blockchain/network"
)

func main() {

	// Server (container)
	// Transport (tcp/udp)
	// Block 					<== currently working on
	// Transaction
	// KeyPair

	trLocal := network.NewLocalTransport("LOCAL")
	trRemote := network.NewLocalTransport("REMOTE")

	trLocal.Connect(trRemote)
	trRemote.Connect(trLocal)

	go func() {
		for {
			trRemote.Send_message(trLocal.Addr(), []byte("Hello World"))
			time.Sleep(1 * time.Second)
		}
	}()

	server_opts := network.ServerOpts{
		Transports: []network.Transport{trLocal},
	}
	s := network.NewServer(server_opts)
	s.Start()

	fmt.Println("working!!")
}
