package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Blackmocca/wedding-ui/pages"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	branding = `
________           .___ _____.__                 
/  _____/  ____   __| _// ____\  |   ______  _  __
/   \  ___ /  _ \ / __ |\   __\|  |  /  _ \ \/ \/ /
\    \_\  (  <_> ) /_/ | |  |  |  |_(  <_> )     / 
\______  /\____/\____ | |__|  |____/\____/ \/\_/  
		\/            \/                           

------------------------------------------------------------------------
http server started on port :%d
------------------------------------------------------------------------
`
)

var (
	port = cast.ToInt(os.Getenv("PORT"))
)

var (
	App = &app.Handler{
		Name:        "Godflow",
		Title:       "Godflow",
		Description: "Make to Easy ETL",
		Icon: app.Icon{
			SVG: "/web/resources/assets/logo/logo-color.svg",
		},
		Styles: []string{
			"/web/resources/styles/tailwind/tailwind-min.css",
		},
		CacheableResources: []string{
			// "/web/resources/styles/tailwind/tailwind-min.css",
			// "/web/resources/assets/logo.svg",
		},
		Fonts: []string{
			"/web/resources/fonts/Kanit-Regular.ttf",
			"/web/resources/fonts/Kanit-Light.ttf",
			"/web/resources/fonts/Kanit-Bold.ttf",
		},
		// AutoUpdateInterval: time.Duration(30 * time.Second),
	}
)

func main() {
	ctx := context.Background()
	// Components routing:
	app.Route("/", &pages.App{})
	app.RunWhenOnBrowser()

	// HTTP routing:
	http.Handle("/", App)

	start(ctx, port)
}

func start(ctx context.Context, port int) {
	portStr := fmt.Sprintf(":%d", port)
	fmt.Printf(branding, port)
	if err := http.ListenAndServe(portStr, nil); err != nil {
		log.Fatal(err)
	}
}
