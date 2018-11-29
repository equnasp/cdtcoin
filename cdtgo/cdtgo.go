package cdtgo

import (
	"crypto/md5"
	"fmt"
	"github.com/equnasp/cdtcoin/config"
	"io"
)

func New() *Frame {
	fmt.Println(config.BANNER[1:])
	return new(Frame)
}

func Md5(str string) string {
	w := md5.New()
	io.WriteString(w, str)
	md5str := fmt.Sprintf("%x", w.Sum(nil))
	return md5str
}
