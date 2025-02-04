package main

import (
	"github.com/SalvaRMT/p2p_chat/internal/peer"
	"os"
)

func main() {
	if len(os.Args) < 4 {
		println("Uso: <modo> <puerto o dirección> <usuario>")
		return
	}

	operation := os.Args[1]
	connection := os.Args[2]
	username := os.Args[3]

	if operation == "connect" {
		peer.ConnectToPeer(connection, username)
	} else {
		peer.StartListening(connection, username)
	}
}
