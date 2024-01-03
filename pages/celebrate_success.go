package pages

import (
	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/Blackmocca/wedding-ui/domain/elements"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var (
	calendarHeart = string(constants.SVG_CALENDAR_HEART_STRING)
)

type CelebrateSuccessPage struct {
	app.Compo
}

func (c *CelebrateSuccessPage) toHome(ctx app.Context, e app.Event) {
	ctx.Navigate("/")
}

func (c *CelebrateSuccessPage) Render() app.UI {
	return app.Div().Class("w-screen max-w-maximum mx-auto h-dvh bg-secondary-base").Body(
		app.Div().Class("flex flex-col w-full h-dvh items-center gap-6 justify-center").Body(
			app.Div().Class("flex flex-col w-full items-center pt-6").Body(
				app.Raw(calendarHeart),
				app.Div().Class("flex flex-col w-10/12 items-center pt-6").Body(
					app.P().Class("text-base text-primary-base font-medium").Text("ขอขอบพระคุณที่มาร่วมยินดี"),
					app.P().Class("text-base text-primary-base font-medium").Text("และอวยพร พวกเรานะคะ/ครับ"),
				),
			),
			app.Div().Class("flex flex-col w-10/12").Body(),
			/* button */
			app.Div().Class("flex flex-col w-10/12 pt-4").Body(
				elements.NewButton(constants.BUTTON_STYLE_PRIMARY).
					Text("กลับหน้าแรก").
					OnClick(c.toHome),
			),
		),

		app.P().Class("absolute w-full text-sm text-primary-base font-medium text-center bottom-0 pb-4").Text("© 2024 NengHuag Wedding. All Rights Reserved"),
	)
}
