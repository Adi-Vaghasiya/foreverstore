package main

import (
	"log"

	p2p "github.com/Aditya-Vaghasiya/fs"
)

func main() {
	conn := p2p.NewTCPConn(":3000")

	if err := conn.GetTCPConn(); err != nil {
		log.Fatal(err)
	}

	select {}
}
