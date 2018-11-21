package main

import (
	"github.com/equnasp/cdtcoin/cdtgo"
	_ "github.com/equnasp/cdtcoin/module/api"
)

func main() {
	app := cdtgo.New()
	app.Start()
	for a := 0; a < 100000; a++ {
		//fmt.Print("a: %d", a)
	}
}
