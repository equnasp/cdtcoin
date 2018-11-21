package module

import (
	"github.com/equnasp/cdtcoin/cdtgo"
	"github.com/equnasp/cdtcoin/module/bigdata/handler"
	"github.com/equnasp/cdtcoin/webapi/faygo"
)

func init() {
	o := obj{}
	cdtgo.Register(o)
}

type obj struct{}

func (p obj) GetInfo() cdtgo.Frame {
	return cdtgo.Frame{
		Name:    "BigData",
		CName:   "大数据处理模块",
		Version: "1.0",
	}
}

func (p obj) Flag() bool {
	return true
}

func (p obj) OnStart() {
}

func (p obj) InitWebApi(frame *faygo.Framework) {
	frame.Route(
		frame.NewGroup("verification",
			frame.NewNamedAPI("Token", "POST", "/token", &handler.Token{}),
		),
	)
}
