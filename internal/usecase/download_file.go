package usecase

import (
	"bytes"
	"context"
)

// DownloadFile - downloads file by ud
func (u *Usecase) DownloadFile(ctx context.Context, fileID string) (*bytes.Buffer, error) {
	return u.repository.DownloadFile(ctx, fileID)
}
