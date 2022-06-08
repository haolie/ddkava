package httpSv

import (
	"fmt"

	"fileSharker/log"
	"fileSharker/src/Server"
	"fileSharker/src/db/dbManager"
	"fileSharker/src/model"
	"fileSharker/src/tool"

	"github.com/kataras/iris/v12"
)

const name = "httpSv"

func init() {
	Server.RegisterCompleted(name, completed)
}

func completed(param interface{}) []error {
	go startHttp()
	return nil
}

func startHttp() {
	app := iris.New()
	app.Get("/downCheck", downCheck)
	app.Get("/downSave", downSave)

	app.Run(iris.Addr(":8780"))
}

func downCheck(ctx iris.Context) {
	downUrl := ctx.URLParam("durl")
	vName := ctx.URLParam("vname")
	fsize := ctx.URLParam("fsize")

	ps := ctx.Params()
	fmt.Println(ps)

	if len(downUrl) == 0 {
		return
	}

	hr := model.NewHResponse("")

	isDown, err := dbManager.SIPageView(downUrl, vName, tool.PressFileSize(fsize))
	if err != nil {
		log.LogError(err, "downCheck")
		hr.ErrorStr = fmt.Sprintf("%v", err)
	} else {
		hr.Content["IsDown"] = isDown
	}

	ctx.JSON(hr)
}

func downSave(ctx iris.Context) {
	downUrl := ctx.URLParam("durl")
	cName := ctx.URLParam("cname")
	fsize := ctx.URLParam("fsize")

	if len(downUrl) == 0 {
		return
	}

	hr := model.NewHResponse("")
	isDown, err := dbManager.XYPageView(downUrl, cName, tool.PressFileSize(fsize))
	if err != nil {
		log.LogError(err, "downCheck")
		hr.ErrorStr = fmt.Sprintf("%v", err)
	} else {
		hr.Content["IsDown"] = isDown
	}

	ctx.JSON(hr)
}
