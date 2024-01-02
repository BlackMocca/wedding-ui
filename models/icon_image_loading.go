package models

import (
	"os"

	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type IconImageLoading struct {
	imagePath string
	cssPath   string
}

func NewIconImageLoading(imgPath string, cssPath string) app.Library {
	return &IconImageLoading{
		imagePath: imgPath,
		cssPath:   cssPath,
	}
}

func (i IconImageLoading) Styles() (path, styles string) {
	stylesBu, err := os.ReadFile(i.cssPath)
	if err != nil {
		panic(err)
	}

	path = i.imagePath
	styles = string(stylesBu)

	return path, styles
}
