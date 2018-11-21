package main

import (
	tp "github.com/equnasp/CDS/goutil/teleport"
	"github.com/equnasp/CDS/webapi/cfgo"
)

//go:generate go build $GOFILE

func main() {
	go tp.GraceSignal()
	cfg := tp.PeerConfig{
		CountTime:  true,
		ListenPort: 9090,
	}

	// auto create and sync config/config.yaml
	cfgo.MustGet("config/config.yaml", true).MustReg("cfg_srv", &cfg)

	srv := tp.NewPeer(cfg)
	srv.RouteCall(new(math))
	srv.ListenAndServe()
}

type math struct {
	tp.CallCtx
}

func (m *math) Add(arg *[]int) (int, *tp.Rerror) {
	var r int
	for _, a := range *arg {
		r += a
	}
	return r, nil
}
