package ginSev

import (
	"context"
	"fmt"
	"net/http"
	"sync"

	"github.com/gin-gonic/gin"
	"lyh/ddkava/fileSharker/Common/Def"
	"lyh/ddkava/fileSharker/Common/Log"
	"lyh/ddkava/fileSharker/Config"
	"lyh/ddkava/fileSharker/Park"
)

const (
	con_f_name = "ginSev"
)

var (
	startOnce   sync.Once
	rootPostMap = make(map[string]func(c *gin.Context), 8)
	rootGetMap  = make(map[string]func(c *gin.Context), 8)
	status      int32
)

func init() {
	Park.RegisterStart(con_f_name, func(ctx context.Context) []error {

		startHServer(ctx)
		return nil
	})
}

func RegisterRoot(key string, isPost bool, f func(c *gin.Context)) {
	if status > 0 {
		Log.Error("can not register root after started")
		return
	}

	var rootMap map[string]func(c *gin.Context)
	if isPost {
		rootMap = rootPostMap
	} else {
		rootMap = rootGetMap
	}

	_, exist := rootMap[key]
	if exist {
		Log.Error(fmt.Sprintf("%s:k={%s} already exist", con_f_name, key))
		return
	}

	rootMap[key] = f
}

func startHServer(ctx context.Context) (errStr string) {

	startOnce.Do(func() {
		go func() {
			port, exists := Config.GetTypeValue[int64](Def.Config_http_Port)
			if !exists {
				panic("config http port not exist")
			}

			api, exists := Config.GetTypeValue[string](Def.Config_Http_Api)
			if !exists {
				panic("config http api not exist")
			}

			viewPath, exists := Config.GetTypeValue[string](Def.Config_Http_View)
			if !exists {
				panic("config Http_View not exist")
			}

			engine := gin.Default()
			engine.StaticFS("/view", http.Dir(viewPath))
			engine.StaticFS("/static", http.Dir(viewPath+"/static"))

			for k, v := range rootPostMap {
				str := api + k
				engine.POST(str, v)
			}

			for k, v := range rootGetMap {
				str := api + k
				engine.GET(str, v)
			}

			err := engine.Run(fmt.Sprintf(":%d", port))
			if err != nil {
				errStr = fmt.Sprintf("%s fail to start http server err:%s", con_f_name, err.Error())
			}
		}()

	})

	return
}
