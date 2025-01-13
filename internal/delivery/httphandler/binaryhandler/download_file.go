package binaryhandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// DownloadFile - downloads file by ud
func (b *BinaryHandler) DownloadFile(ctx iris.Context) {
	fileID := ctx.Params().Get("id")

	if fileID == "" {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	file, err := b.usecase.DownloadFile(ctx, fileID)

	if err != nil {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	contentType := http.DetectContentType(file.Bytes())

	ctx.ContentType(contentType)
	ctx.StatusCode(http.StatusOK)
	ctx.Write(file.Bytes())
}
