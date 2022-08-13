package aliyun

import (
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/k0kubun/go-ansi"
	"github.com/schollz/progressbar/v3"
)

func NewListener() oss.ProgressListener {
	return &listener{}
}

type listener struct {
	bar *progressbar.ProgressBar
}

func (l *listener) ProgressChanged(event *oss.ProgressEvent) {
	switch event.EventType {
	case oss.TransferStartedEvent:
		l.bar = progressbar.NewOptions64(event.TotalBytes,
			progressbar.OptionSetWriter(ansi.NewAnsiStdout()),
			progressbar.OptionEnableColorCodes(true),
			progressbar.OptionShowBytes(true),
			progressbar.OptionSetWidth(30),
			progressbar.OptionSetDescription("开始上传:"),
			progressbar.OptionSetTheme(progressbar.Theme{
				Saucer:        "=",
				SaucerHead:    ">",
				SaucerPadding: " ",
				BarStart:      "[",
				BarEnd:        "]",
			}),
		)
	case oss.TransferDataEvent:
		l.bar.Add64(event.RwBytes)
	case oss.TransferCompletedEvent:
	case oss.TransferFailedEvent:
	}
}
