package impl

import (
	"github.com/gin-gonic/gin"
	"lyh/ddkava/fileSharker/Common/Def"
	"lyh/ddkava/fileSharker/FileScan"
	"lyh/ddkava/fileSharker/httpSv/ginSev"
	HttpTools "lyh/ddkava/fileSharker/httpSv/httpTools"
)

func init() {
	ginSev.RegisterRoot(Def.Http_StartScan, false, startScan)
}

func startScan(ctx *gin.Context) {
	dirPath, exists := ctx.GetQuery("dir")
	if !exists || dirPath == "" {
		ctx.JSON(200, HttpTools.CreateErrHSResponse("need dir params"))
		return
	}

	err := FileScan.StartScanDir(ctx, dirPath, true)
	if err != nil {
		ctx.JSON(200, HttpTools.CreateErrHSResponse(err.Error()))
	} else {
		ctx.JSON(200, HttpTools.CreateSuccessHSResponse("Success"))
	}
}
