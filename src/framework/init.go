package framework

// 初始化框架

import (
	"signaling/src/glog"
)

func Init() error {
	glog.SetLogDir("./log")
	glog.SetLogFileName("signaling")
	glog.SetLogLevel("DEBUG")
	// glog.SetLogToStderr(true)
	glog.SetLogAlsoToStedrr(true)
	return nil
}
