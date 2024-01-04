package pages

import (
	"strings"
	"time"

	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/Blackmocca/wedding-ui/domain/elements"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

var (
	gift    = string(constants.SVG_GIFT_STRING)
	copy    = string(constants.SVG_COPY_STRING)
	address = []string{"คุณเหน่ง เบอร์โทร 083-554-6499", "เดอะไนน์ เรสซิเดนซ์ เลขที่ 9", "ซ. พระราม2 56 แยก 1 ถ. พระราม 2", "แขวงแสมดำ เขตบางขุนเทียน กทม. 10150"}
)

type SendGiftPage struct {
	app.Compo

	isCopy bool
}

func (c *SendGiftPage) OnMount(ctx app.Context) {
	app.Window().ScrollToID("rootContainer")
}

func (c *SendGiftPage) toHome(ctx app.Context, e app.Event) {
	ctx.Navigate("/")
}

func (c *SendGiftPage) clipboard(ctx app.Context, e app.Event) {
	copyText := strings.Join(address, " ")

	execCommandCopyFunc := app.Window().Get("execCommandCopy")
	execCommandCopyFunc.Invoke(copyText)

	c.isCopy = true
	ctx.Async(func() {
		time.Sleep(2 * time.Second)
		c.isCopy = false
		c.Update()
	})
	c.Update()
}

func (c *SendGiftPage) Render() app.UI {
	return app.Div().Class("w-screen max-w-maximum mx-auto h-dvh bg-secondary-base").ID("rootContainer").Body(
		app.Div().Class("flex flex-col w-full h-dvh items-center gap-6 justify-center").Body(
			app.Div().Class("flex flex-col w-full items-center pt-6").Body(
				app.Raw(gift),
				app.Div().Class("flex flex-col w-10/12 items-center pt-4").Body(
					app.P().Class("text-xl text-primary-base font-medium").Text("ส่งของขวัญ"),
				),
			),
			app.Div().Class("flex flex-col w-10/12 gap-4").Body(
				app.P().Class("font-medium text-base text-primary-base font-regular").Text("ที่อยู่ผู้รับ"),
				app.Div().Class("flex flex-col w-full h-[13.1875rem] border-2 border-primary-base").Body(
					app.Div().Class("p-4").Body(
						app.Range(address).Slice(func(i int) app.UI {
							return app.P().Class("text-base").ID(app.FormatString("addr-id-%d", i)).Text(address[i])
						}),
					),
					app.Div().Class("flex w-full h-full p-2 pb-2 items-end justify-end").Body(
						app.Div().Class("justify-end").Body(
							app.If(c.isCopy, app.Img().Class("w-[27px] h-[27px]").Src(string(constants.IMG_CHECKMARK))).
								Else(
									app.Raw(copy),
								),
						),
					).OnClick(c.clipboard),
				),
			),
			/* button */
			app.Div().Class("flex flex-col w-10/12 pt-4").Body(
				elements.NewButton(constants.BUTTON_STYLE_PRIMARY).
					Text("กลับหน้าแรก").
					OnClick(c.toHome),
			),
		),

		app.P().Class("absolute w-full max-w-maximum mx-auto text-sm text-primary-base font-base text-center bottom-0 pb-4").Text("© 2024 NengHuag Wedding. All Rights Reserved"),
	)
}
