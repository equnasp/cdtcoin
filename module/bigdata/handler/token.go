package handler

import (
	"github.com/equnasp/CDS/cdsgo"
	"github.com/equnasp/CDS/webapi/faygo"
	"strconv"
	"time"
)

type Token struct {
	App_id     string `param:"<in:header> <desc:AppID> <nonzero>"`
	App_secret string `param:"<in:header> <desc:AppSecret> <nonzero>"`
	Signature  string `param:"<in:header> <desc:校验码> <nonzero>"`
}

func (a *Token) Serve(ctx *faygo.Context) error {
	var signature = cdsgo.Md5(a.App_id + a.App_secret + "aa3461bd-9105-4191-b19c-944563f26348")
	var data map[string]string = map[string]string{}

	if signature != a.Signature {
		return cdsgo.ReturnVal(ctx, "996", data, "signature验证失败")
	}

	data["token"] = CreateToken(a.App_id, a.App_secret)

	return cdsgo.ReturnVal(ctx, "1", data, "")
}

func CreateToken(App_id string, App_secret string) string {
	return cdsgo.Md5(App_id + App_secret + strconv.FormatInt(time.Now().UnixNano(), 20))
}

// 补充API文档信息
func (a *Token) Doc() faygo.Doc {
	return faygo.Doc{
		// API接口说明
		Note: "Token",
		// 响应说明或示例
		Return: "获取Token",
	}
}
