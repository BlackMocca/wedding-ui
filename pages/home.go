package pages

import (
	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/Blackmocca/wedding-ui/domain/elements"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type Home struct {
	app.Compo
}

func (h *Home) toCelebrate(ctx app.Context, e app.Event) {
	ctx.Navigate("/celebrate")
}
func (h *Home) toPromptPay(ctx app.Context, e app.Event) {
	ctx.Navigate("/promptpay")

}
func (h *Home) toSendGift(ctx app.Context, e app.Event) {
	ctx.Navigate("/sendgift")
}

func (h *Home) Render() app.UI {
	return app.Div().Class("flex flex-col w-screen min-h-screen bg-secondary-base overflow-y-auto").Body(
		app.Div().Class("flex flex-col w-full min-h-screen items-center gap-6 justify-center overflow-y-auto").Body(
			app.Div().Class("flex w-10/12 h-10/12 items-center justify-center").Body(
				app.Img().Class("relative w-full h-10/12 z-1 p-4 opacity-100").Src(string(constants.IMG_HOME_COVER)),
			),
		),
		app.Div().Class("flex flex-col w-screen min-h-[50vh] bg-secondary-base bg-red-500 justify-center items-center").Body(
			app.Div().Class("flex flex-col w-full items-center gap-6").Body(
				app.Div().Class("flex flex-col w-10/12").Body(
					elements.NewButton(constants.BUTTON_STYLE_SECONDARY).
						Text("มาร่วมอวยพร บ่าว-สาว").
						OnClick(h.toCelebrate),
				),
				app.Div().Class("flex flex-col w-10/12").Body(
					elements.NewButton(constants.BUTTON_STYLE_SECONDARY).
						Text("พร้อมเพย์").
						OnClick(h.toPromptPay),
				),
				app.Div().Class("flex flex-col w-10/12").Body(
					elements.NewButton(constants.BUTTON_STYLE_SECONDARY).
						Text("ส่งของขวัญ").
						OnClick(h.toSendGift),
				),
			),
		),
		app.P().Class("relative w-full text-sm text-primary-base font-medium text-center bottom-0 pb-4").Text("© 2024 NengHuag Wedding. All Rights Reserved"),
	)
}
