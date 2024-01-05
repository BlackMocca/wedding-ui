package pages

import (
	"fmt"
	"time"

	"github.com/Blackmocca/wedding-ui/domain/core/api"
	"github.com/Blackmocca/wedding-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

type ListCelebrate struct {
	app.Compo

	celebrates []*models.Celebrate
}

func (l *ListCelebrate) transfromDateWord(dt time.Time) string {
	dt = dt.Add(7 * time.Hour)
	// Thai month names
	thaiMonths := [...]string{
		"",
		"มกราคม", "กุมภาพันธ์", "มีนาคม",
		"เมษายน", "พฤษภาคม", "มิถุนายน",
		"กรกฎาคม", "สิงหาคม", "กันยายน",
		"ตุลาคม", "พฤศจิกายน", "ธันวาคม",
	}

	// Format the timestamp with Thai date format
	formattedTimestamp := fmt.Sprintf("วันที่ %d %s %d เวลา %s", dt.Day(), thaiMonths[dt.Month()], dt.Year()+543, dt.Format("15:04"))
	return formattedTimestamp
}

func (l *ListCelebrate) OnNav(ctx app.Context) {
	celebrate, err := api.WeddingAPI.Fetch(ctx)
	if err != nil {
		panic(err)
	}

	l.celebrates = celebrate
}

func (l ListCelebrate) Render() app.UI {
	return app.Div().Class("w-screen max-w-maximum mx-auto h-screen bg-secondary-base").ID("rootContainer").Body(
		app.Div().Class("flex w-full pt-4 pb-4 items-center justify-center").Body(
			app.H4().Class("text-xl text-primary-base font-medium").Text(fmt.Sprintf("รายการคำอวยพร (%d)", len(l.celebrates))),
		),
		app.Div().Class("flex flex-col pt-4 pb-4 w-full gap-6 pl-4 pr-4").Body(
			app.If(len(l.celebrates) > 0, app.Range(l.celebrates).Slice(func(index int) app.UI {
				return app.Div().Class("flex flex-col w-full p-4 border text-primary-base shadow-md").Body(
					/* head card */
					app.Div().Class("w-full ").Body(
						app.P().Class("text-md text-primary-base font-medium break-words").Text(l.celebrates[index].CelebrateFrom),
						app.P().Class("text-sm text-gray-400 font-base").Text(l.transfromDateWord(l.celebrates[index].CreatedAt)),
					),
					app.Hr().Class("h-px mt-4 mb-4 bg-gray-300"),
					/* detail card */
					app.Div().Class().Body(
						app.P().Class("text-sm text-primary-base font-base break-words").Text(l.celebrates[index].CelebrateText),
					),
				)
			})),
		),
	)
}
