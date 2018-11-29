package main

import (
	"fmt"
	"github.com/equnasp/cdtcoin/cdtgo"
	"github.com/equnasp/cdtcoin/module/coin"
	_ "github.com/equnasp/cdtcoin/plugin/blockchain"
)

func main() {
	//var module = flag.String("m", "start", "运行模块（start：开启，wallet：钱包）")
	//flag.Parse()
	app := cdtgo.New()
	app.Start()

	block := coin.NewBlock("0", []byte("test"))

	fmt.Println(block.Hash)
}
