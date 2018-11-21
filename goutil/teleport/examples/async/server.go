package main

import (
	"time"

	tp "github.com/equnasp/CDS/goutil/teleport"
)

//go:generate go build $GOFILE

func main() {
	srv := tp.NewPeer(tp.PeerConfig{
		CountTime:  true,
		ListenPort: 9090,
	})
	srv.RouteCall(new(test))
	srv.ListenAndServe()
}

type test struct {
	tp.CallCtx
}

func (t *test) Wait3s(arg *string) (string, *tp.Rerror) {
	time.Sleep(3 * time.Second)
	return *arg + " -> OK", nil
}
