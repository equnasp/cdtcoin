package blockchain

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
		Name:    "Blockchain",
		CName:   "区块链模块",
		Version: "1.0",
	}
}

func (p obj) Flag() bool {
	return true
}

func (p obj) OnStart() {

}
