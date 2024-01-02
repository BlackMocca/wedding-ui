package pages

import (
	"fmt"

	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/Blackmocca/wedding-ui/domain/components"
	"github.com/Blackmocca/wedding-ui/models"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
)

const (
	navHeaderTitle = "New Connection"
	tagConnection  = "ConnectionList"
)

type Home struct {
	app.Compo

	connectionList []*models.ConnectionList
}

func (h *Home) ConnectionList() []*models.ConnectionList {
	return h.connectionList
}

func (h *Home) getDataStorage(ctx app.Context) error {
	if err := ctx.LocalStorage().Get(string(constants.STORAGE_CONNECTION_LIST), &h.connectionList); err != nil {
		return err
	}
	return nil
}

func (h *Home) OnMount(ctx app.Context) {
	h.getDataStorage(ctx)
}

func (h *Home) OnNav(ctx app.Context) {
	fmt.Println("on nav")
}

func (h *Home) Event(ctx app.Context, event constants.Event, data interface{}) {
	switch event {
	case constants.EVENT_UPDATE:
		if _, ok := data.(*models.ConnectionList); ok {
			if err := h.getDataStorage(ctx); err != nil {
				app.Log(err)
				return
			}
		}
	}

	h.Update()
}

func (h *Home) Render() app.UI {
	return app.Div().Class("flex w-screen h-screen").Body(
		&components.Nav{
			ConnectionList: h.connectionList,
		},
		app.Div().Class("flext flex-col w-full").Body(
			components.NewNavHeader(components.NavHeaderProp{Title: navHeaderTitle}),
			app.Div().Class().Body(
				&components.FormConnection{Parent: h},
			),
		),
	)
}
