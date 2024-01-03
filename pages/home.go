package pages

import (
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
	return app.Div().Class("flex flex-col w-screen h-dvh").Body(
		app.P().Class("font-sans").Text("จาก HereerererereHereerererere "),
		app.P().Class("font-regular").Text("จาก Hereerererere "),
		app.P().Class("font-medium").Text("สวัสดีครับที่นี้ Hereerererere"),
		app.P().Class("prompt").Text("อันนี้ kanit คับ"),
		// app.Img().Class("absolute w-screen h-dvh z-1 p-4 opacity-100").Src(string(constants.IMG_HOME_COVER)),
	)
}
