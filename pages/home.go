package pages

import (
	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo
}

func (h *Home) OnMount(ctx app.Context) {
}

func (h *Home) OnNav(ctx app.Context) {
}

func (h *Home) Render() app.UI {
	return app.Div().Class("flex w-screen h-dvh").Body(
		app.Img().Class("absolute w-screen h-dvh z-1 p-4 opacity-100").Src(string(constants.IMG_HOME_COVER)),
	)
}
