package httphandler

import (
	"net/http"

	"github.com/kataras/iris/v12"
)

// DownloadFile - downloads file by ud
func (h *HttpHandler) DownloadFile(ctx iris.Context) {
	fileID := ctx.Params().Get("id")

	if fileID == "" {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	file, err := h.usecase.DownloadFile(ctx, fileID)

	if err != nil {
		ctx.StopWithStatus(http.StatusNotFound)
		return
	}

	contentType := http.DetectContentType(file.Bytes())

	ctx.ContentType(contentType)
	ctx.StatusCode(http.StatusOK)
	ctx.Write(file.Bytes())
}
