package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"path"
	"syscall"

	"lyh/ddkava/fileSharker/Common/Def"
	"lyh/ddkava/fileSharker/Common/Log"
	"lyh/ddkava/fileSharker/Config"
	"lyh/ddkava/fileSharker/Park"

	_ "lyh/ddkava/fileSharker/httpSv"
)

func main() {
	err := Config.Load(path.Dir(os.Args[0]))
	if err != nil {
		panic("load config file error:" + err.Error())
	}

	logPath, exists := Config.GetTypeValue[string](Def.Config_Log_Path)
	if !exists {
		panic("log path not exists")
	}

	Log.InitLog(logPath)

	ctx, cancel := context.WithCancel(context.Background())
	Log.Info("____ startLoading ____")
	errList := Park.Load(ctx)
	if len(errList) > 0 {
		Log.Info("load failed")
		Log.Error("load failed")
		for _, err := range errList {
			Log.Error(err.Error())
		}

		return
	}

	Log.Info("____ start Server ____")
	errList = Park.Start(ctx)
	if len(errList) > 0 {
		Log.Info("startServer Failed")
		Log.Error("startServer Failed")
		for _, err := range errList {
			Log.Error(err.Error())
		}

		return
	}

	Log.Info("____ start Success ____")
	c := make(chan os.Signal)
	signal.Notify(c, os.Interrupt, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go waitExit(c, cancel)

	waitEx := ctx.Done()
	<-waitEx
}

func waitExit(signalChan chan os.Signal, cancel context.CancelFunc) {
	for i := range signalChan {
		switch i {
		case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
			Log.Info(fmt.Sprintf("exit signal :%v", i))
			cancel()
		}
	}
}
