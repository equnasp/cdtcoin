# surfer    [![GoDoc](https://godoc.org/github.com/tsuna/gohbase?status.png)](https://godoc.org/github.com/equnasp/cdtcoin/webapi/surfer) [![GitHub release](https://img.shields.io/github/release/henrylee2cn/surfer.svg)](https://github.com/equnasp/cdtcoin/webapi/surfer/releases)


Surfer 是一款Go语言编写的高并发 web 客户端，拥有surf与phantom两种下载内核，高度模拟浏览器行为，可实现模拟登录等功能。

高并发爬虫[Pholcus](https://github.com/equnasp/cdtcoin/webapi/pholcus)的专用下载器。（官方QQ群：Go大数据 42731170，欢迎加入我们的讨论）

## 特性

- 支持 `surf` 和 `phantomjs` 两种下载内核
- 支持大量随机的User-Agent
- 支持缓存cookie
- 支持`http`/`https`两种协议

## 用法
```
package main

import (
    "github.com/equnasp/cdtcoin/webapi/surfer"
    "io/ioutil"
    "log"
)

func main() {
    // 默认使用surf内核下载
    resp, err := surfer.Download(&surfer.Request{
        Url: "http://github.com/equnasp/cdtcoin/webapi/surfer",
    })
    if err != nil {
        log.Fatal(err)
    }
    b, err := ioutil.ReadAll(resp.Body)
    log.Println(string(b), err)

    // 指定使用phantomjs内核下载
    resp, err = surfer.Download(&surfer.Request{
        Url:          "http://github.com/henrylee2cn",
        DownloaderID: 1,
    })
    if err != nil {
        log.Fatal(err)
    }
    b, err = ioutil.ReadAll(resp.Body)
    log.Println(string(b), err)

    resp.Body.Close()
    surfer.DestroyJsFiles()
}
```

[完整示例](https://github.com/equnasp/cdtcoin/webapi/surfer/blob/master/example/example.go)


## 开源协议

Surfer 项目采用商业应用友好的[Apache License v2](https://github.com/equnasp/cdtcoin/webapi/surfer/raw/master/LICENSE).发布
