package components

import (
	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/Blackmocca/wedding-ui/domain/core"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type SuccessModal struct {
	app.Compo
	Parent core.ParentNotify

	displayText string
}

func NewSuccessModal(displayText string) *SuccessModal {
	return &SuccessModal{displayText: displayText}
}

func (s *SuccessModal) Render() app.UI {
	return app.Div().Class("flex fixed w-screen h-dvh overflow-hidden bg-red-500 bg-opacity-75 justify-center items-center").Body(
		app.Div().Class("flex flex-col w-3/5 h-40 bg-blue-300 justify-center").Body(
			app.Img().Class("w-2/5 h-24").Src(string(constants.ICON_SUCCESS)),
			app.P().Text("modal"),
		),
	)
}
