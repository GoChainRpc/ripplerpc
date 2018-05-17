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

	resultAcounttx, err := client.GetAccountTx("rDu13xhh9pwc55Gpe8Gec9818YCBD8ZqJ7", -1)
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println("account tx :", resultAcounttx)
	}

	tx, err := client.GetTx("2CA49646665A08B9560F5D6E91B6BED1C58B6CD0C8742BACD73EC74D076BD858")
	if err != nil {
		log.Fatal(err)
	} else {
		log.Println(" tx :", tx)
	}
}
