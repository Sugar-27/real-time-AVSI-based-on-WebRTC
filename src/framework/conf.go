package framework

import "signaling/src/goconfig"

type FrameWorkConf struct {
	logDir          string
	logFile         string
	logLevel        string
	logAlsoToStdErr bool

	httpPort      int
	httpStaticDir string
}

var configFile *goconfig.ConfigFile

func loadConf(confFile string) (*FrameWorkConf, error) {
	var err error

	configFile, err = goconfig.LoadConfigFile(confFile)
	if err != nil {
		return nil, err
	}

	conf := &FrameWorkConf{}
	conf.logDir, err = configFile.GetValue("log", "logDir")
	if err != nil {
		return nil, err
	}
	conf.logFile, err = configFile.GetValue("log", "logFile")
	if err != nil {
		return nil, err
	}
	conf.logLevel, err = configFile.GetValue("log", "logLevel")
	if err != nil {
		return nil, err
	}
	conf.logAlsoToStdErr, err = configFile.Bool("log", "logAlsoToStedrr")
	if err != nil {
		return nil, err
	}
	conf.httpPort, err = configFile.Int("http", "port")
	if err != nil {
		return nil, err
	}
	conf.httpStaticDir, err = configFile.GetValue("http", "staticDir")
	if err != nil {
		return nil, err
	}

	return conf, nil
}

// 暴露方法给非framework的模块使用
func GetStaticDir() string {
	return gconf.httpStaticDir
}
