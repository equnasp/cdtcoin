package cdtgo

import (
	"fmt"
	"github.com/equnasp/cdtcoin/webapi/faygo"
)

type Frame struct {
	Name    string
	CName   string
	Version string
}

var (
	module map[string]Need
)

func init() {
	module = make(map[string]Need)
}

type Need interface {
	Flag() bool //必须告诉容器是否启动
	InitWebApi(frame *faygo.Framework)
	OnStart() //必须有启动方法
	GetInfo() Frame
}

func (frame *Frame) Start() {
	app := faygo.New("cdtcoin")
	for name, m := range module {
		if m.Flag() {
			go m.InitWebApi(app)
			go m.OnStart()
			fmt.Printf("启动插件：%s\n", name)
		} else {
			fmt.Printf("不启动插件：%s\n", name)
		}
	}

	app.Run()
}

//每个插件在初始化时必须注册
func Register(m Need) {
	info := m.GetInfo()
	module[info.Name] = m
}
