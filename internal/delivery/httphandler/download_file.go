package httphandler

import (
	"io"
	"net/http"
	"os"

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

	newFile, err := os.OpenFile("file", os.O_CREATE|os.O_WRONLY, 0777)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	_, err = io.Copy(newFile, file)

	if err != nil {
		ctx.StopWithStatus(http.StatusInternalServerError)
		return
	}

	ctx.StatusCode(http.StatusOK)
	ctx.SendFile("file", "file")
	// ctx.Write(file.Bytes())

}
