package cdtgo

import (
	"fmt"
)

type Frame struct {
	Name    string
	CName   string
	Version string
}

var (
	plugin map[string]Need
)

func init() {
	plugin = make(map[string]Need)
}

type Need interface {
	Flag() bool //必须告诉容器是否启动
	OnStart()   //必须有启动方法
	GetInfo() Frame
}

func (frame *Frame) Start() {
	for name, m := range plugin {
		if m.Flag() {
			go m.OnStart()
			fmt.Printf("启动插件：%s\n", name)
		} else {
			fmt.Printf("不启动插件：%s\n", name)
		}
	}
}

//每个插件在初始化时必须注册
func Register(m Need) {
	info := m.GetInfo()
	plugin[info.Name] = m
}
