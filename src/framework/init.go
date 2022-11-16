package framework

// 初始化框架

import (
	"fmt"
	"signaling/src/glog"
)

var gconf *FrameWorkConf

func Init(confFile string) error {
	var err error
	gconf, err = loadConf(confFile)
	if err != nil {
		panic(err)
	}

	fmt.Printf("%+v\n", gconf)

	glog.SetLogDir(gconf.logDir)
	glog.SetLogFileName(gconf.logFile)
	glog.SetLogLevel(gconf.logLevel)
	// glog.SetLogToStderr(true)
	glog.SetLogAlsoToStedrr(gconf.logAlsoToStdErr)
	return nil
}
