package main

import (
	"github.com/equnasp/cdtcoin/cdtgo"
	_ "github.com/equnasp/cdtcoin/module/bigdata"
)

func main() {
	app := cdtgo.New()
	app.Start()
}
