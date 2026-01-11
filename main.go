package main

import (
	"bcpocker/p2p"
	"log"
	"time"
)

func main() {
	cfg := p2p.ServerConfig{
		Version:     "BCPOKER V0.1-alpha",
		ListenAddr:  ":3000",
		GameVariant: p2p.TexasHoldem,
	}
	server := p2p.NewServer(cfg)
	go server.Start()

	time.Sleep(1 * time.Second)

	remoteCfg := p2p.ServerConfig{
		Version:     "BCPOKER V0.1-alpha",
		ListenAddr:  ":4000",
		GameVariant: p2p.TexasHoldem,
	}
	remoteServer := p2p.NewServer(remoteCfg)
	go remoteServer.Start()
	if err := remoteServer.Connect(":3000"); err != nil {
		log.Fatal(err)
	}

	select {}

	server.Start()
}
