package main

import (
	"flag"
	"signaling/src/framework"
	"signaling/src/glog"
)

func main() {
	flag.Parse()

	err := framework.Init("./conf/framework.conf")
	if err != nil {
		panic(err)
	}

	glog.Info("hello, world!")
	glog.Debug("Test Debug")

	// 静态资源处理 /static
	framework.RegisterStaticUrl()

	// 启动http server
	go startHttp()

	// 启动https server
	startHttps()
}

func startHttp() {
	err := framework.StartHttp()
	if err != nil {
		panic(err)
	}
}

func startHttps() {
	err := framework.StartHttps()
	if err != nil {
		panic(err)
	}
}
