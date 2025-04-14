package main

import (
	"context"
	"os"
	"os/signal"
	"syscall"

	"fileSharker/log"
	_ "fileSharker/src"
	"fileSharker/src/Server"
)

var (
	cancelFn func()
)

func main() {
	var ctx context.Context
	ctx, cancelFn = context.WithCancel(context.Background())
	go func() {
		if !Server.Start(ctx) {
			cancelFn()
			return
		}

		log.LogInfo("server started")

	}()

	<-ctx.Done()
}

func setExit() {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

	go func() {
		for sg := range ch {
			switch sg {
			case syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT:
				doExit()
			default:

			}
		}
	}()
}

func doExit() {
	log.LogInfo("server exit")
	cancelFn()

}
