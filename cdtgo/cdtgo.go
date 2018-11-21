package cdtgo

import (
	"crypto/md5"
	"fmt"
	"github.com/equnasp/CDS/webapi/faygo"
	"io"
)

const (
	VERSION = "1.0.0"
	BANNER  = `
   __    ____  _  _    ____  _____  _  _ 
  /__\  (_  _)( \/ )  (  _ \(  _  )( \/ )
 /(__)\  _)(_  )  (    ) _ < )(_)(  )  ( 
(__)(__)(____)(_/\_)  (____/(_____)(_/\_)  VERSION ` + VERSION + "\n"
)

func New() *Frame {
	fmt.Println(BANNER[1:])
	return new(Frame)
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str) //将str写入到w中
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}

func ReturnVal(ctx *faygo.Context, code string, data map[string]string, error_msg string) error {
	return ctx.JSON(200,
		faygo.Map{
			"code":      code,
			"error_msg": error_msg,
			"data":      data,
		}, true,
	)
}
