package api

import (
	"github.com/equnasp/cdtcoin/cdtgo"
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
