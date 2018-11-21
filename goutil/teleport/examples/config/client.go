package main

import (
	tp "github.com/equnasp/CDS/goutil/teleport"
	"github.com/equnasp/CDS/webapi/cfgo"
)

//go:generate go build $GOFILE

func main() {
	cfg := tp.PeerConfig{}

	// auto create and sync config/config.yaml
	cfgo.MustGet("config/config.yaml", true).MustReg("cfg_cli", &cfg)

	cli := tp.NewPeer(cfg)
	defer cli.Close()

	sess, err := cli.Dial(":9090")
	if err != nil {
		tp.Fatalf("%v", err)
	}

	var result int
	rerr := sess.Call("/math/add?push_status=yes",
		[]int{1, 2, 3, 4, 5},
		&result,
	).Rerror()

	if rerr != nil {
		tp.Fatalf("%v", rerr)
	}
	tp.Printf("result: 1+2+3+4+5 = %d", result)
}
