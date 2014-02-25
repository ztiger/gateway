package main

import (
	"base/config"
	"base/log"
	"base/socket"
	"fmt"
	"gateway/services"
	"os"
	"time"
)

var (
	SrvID      string
	configFile string = "gateway.conf"
	iniConfig  config.ConfigContainer

	loggerType string
	loggerName string
	Logger     *log.BeeLogger

	socketSrv *socket.Server = nil
)

func main() {
	if err := initConfig(); err != nil {
		panic(err)
		os.Exit(1)
	}
	if err := initLogger(); err != nil {
		panic(err)
		os.Exit(1)
	}
	if err := initServer(); err != nil {
		panic(err)
		Logger.Error("Error", err)
		Logger.Flush()
		os.Exit(1)
	}
	Logger.Flush()
}

//初始化配置文件
func initConfig() error {
	var err error
	iniConfig, err = config.NewConfig(config.Ini, configFile)
	return err
}

//初始化日志系统
func initLogger() error {
	loggerType = iniConfig.String("logger::loggerType")
	loggerName = iniConfig.String("logger::loggerName")
	loggerLevel, err := iniConfig.Int("logger::loggerLevel")
	if err != nil {
		return err
	}

	Logger = log.NewLogger(1000)
	Logger.SetLogger(loggerType, "{\"fileName\":\""+loggerName+"\"}")
	Logger.Setlevel(loggerLevel)
	return nil
}

func initServer() error {
	SrvID = iniConfig.String("server::srvID")
	listenPort, err := iniConfig.Int("server::port")
	if err != nil {
		return err
	}
	timeOut, e := iniConfig.Int("server::timeOut")
	if e != nil {
		return e
	}
	var addr string = "0.0.0.0:" + fmt.Sprint(listenPort)

	socketConf := socket.NewConfig()
	socketConf.CloseingTimeout = time.Duration(timeOut)
	socketConf.Addr = addr
	socketConf.CodecFactory = socket.NewDefaultCodecFactory()
	socketConf.ConnectedHandler = services.ConnectedHandler
	socketConf.DisconnectHandler = services.DisconnectHandler
	socketConf.MessageHandler = services.MessageHandler

	socketSrv, err = socket.NewServer(socketConf)
	if err != nil {
		return err
	}
	socketSrv.Start()
	return nil
}
