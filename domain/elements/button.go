package elements

import (
	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	primaryButtonStyle   = "px-4 py-2 bg-primary-base text-secondary-base hover:pointer-cursor hover:shadow hover:shadow-primary-base"
	secondaryButtonStyle = "px-4 py-2 text-primary-base bg-secondary-base border-2 border-primary-base hover:bg-gray-100 hover:pointer-cursor hover:shadow"
)

func getButtonBaseStyle(buttonStyle constants.ButtonStyle) string {
	switch buttonStyle {
	case constants.BUTTON_STYLE_PRIMARY:
		return primaryButtonStyle
	case constants.BUTTON_STYLE_SECONDARY:
		return secondaryButtonStyle
	}
	return ""
}

func NewButton(buttonStyle constants.ButtonStyle) app.HTMLButton {
	return app.Button().
		Class(getButtonBaseStyle(buttonStyle)).
		Text("Button")
}
