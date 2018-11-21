package module

import (
	"github.com/equnasp/CDS/cdsgo"
	"github.com/equnasp/CDS/module/bigdata/handler"
	"github.com/equnasp/CDS/webapi/faygo"
)

func init() {
	o := obj{}
	cdsgo.Register(o)
}

type obj struct{}

func (p obj) GetInfo() cdsgo.Frame {
	return cdsgo.Frame{
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
