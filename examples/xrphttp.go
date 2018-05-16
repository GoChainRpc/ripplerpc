package main

import (
	"log"
	"github.com/GoChainRpc/ripplerpc"
)

func main() {
	connCfg := &ripplerpc.ConnConfig{
		Host:         "s.altnet.rippletest.net:51234",
		HTTPPostMode: true, //  supports HTTP POST mode
		DisableTLS:   true, //  does not provide TLS by default
	}

	client, err := ripplerpc.New(connCfg)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Shutdown()

	result, err := client.GetAccountInfo("rDu13xhh9pwc55Gpe8Gec9818YCBD8ZqJ7", "current")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("account info :", result)
	}
}
