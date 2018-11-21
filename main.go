package main

import (
	"github.com/equnasp/CDS/cdsgo"
	_ "github.com/equnasp/CDS/module/bigdata"
)

func main() {
	app := cdsgo.New()
	app.Start()
}
