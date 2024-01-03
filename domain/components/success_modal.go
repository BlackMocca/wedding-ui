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
	return app.Div().Class("flex fixed w-screen h-dvh overflow-hidden bg-chacoal bg-opacity-75 justify-center items-center min-h-screen rounded").Body(
		app.Div().Class("flex flex-col gap-4 w-4/5 h-64 bg-secondary-base items-center").Body(
			app.Img().Class("w-24 h-28 pt-4").Src(string(constants.ICON_SUCCESS)),
			app.P().Class("text-md").Text(s.displayText),
			app.Div().Class("flex h-full w-full items-center justify-center items-end").Body(
				app.Button().Class("relative text-xl w-4/5 py-2 px-4 rounded bg-green-500 hover:bg-green-600 text-secondary-base").Text("Close"),
			),
		),
	)
}
