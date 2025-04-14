package Server

import (
	"context"
	"fmt"

	"fileSharker/config"
	"fileSharker/log"
)

type RegisterCallBacker func(ctx context.Context) []error

var (
	loadCallBackMap  = make(map[string]RegisterCallBacker, 4)
	startCallBackMap = make(map[string]RegisterCallBacker, 4)
)

// 注册数据加载
func RegisterLoadFun(key string, cb RegisterCallBacker) {
	if _, exists := loadCallBackMap[key]; exists {
		return
	}

	loadCallBackMap[key] = cb
}

// 注册服务器启动
func RegisterStartFun(key string, cb RegisterCallBacker) {
	if _, exists := startCallBackMap[key]; exists {
		return
	}

	startCallBackMap[key] = cb
}

// 加载
func load(ctx context.Context) bool {
	for key, cb := range loadCallBackMap {
		errs := cb(ctx)
		if len(errs) > 0 {
			log.LogInfo(fmt.Sprintf("%s load failed", key))
			log.LogErrs(errs)

			return false
		}
	}

	return true
}

// 开始运行
func start(ctx context.Context) bool {
	for key, cb := range startCallBackMap {
		errs := cb(ctx)
		if len(errs) > 0 {
			log.LogInfo(fmt.Sprintf("%s start failed", key))
			log.LogErrs(errs)

			return false
		}
	}

	return true
}

func Start(ctx context.Context) bool {

	err := config.Load("")
	if err != nil {
		log.LogError(fmt.Errorf("load config err:%s", err))
		return false
	}

	if !load(ctx) {
		return false
	}

	if !start(ctx) {
		return false
	}

	return true
}
