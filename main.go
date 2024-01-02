package main

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/Blackmocca/wedding-ui/constants"
	"github.com/Blackmocca/wedding-ui/pages"
	"github.com/maxence-charriere/go-app/v9/pkg/app"
	"github.com/spf13/cast"
)

const (
	branding = `
------------------------------------------------------------------------
http server started on port :%d
------------------------------------------------------------------------
`
)

var (
	port = cast.ToInt(constants.GetEnv("PORT", "8080"))
)

var (
	App = &app.Handler{
		Name:         "Wedding",
		Title:        "N&H Wedding",
		Description:  "Wedding",
		LoadingLabel: "Loading {progress}%",
		Lang:         "th",
		Icon: app.Icon{
			SVG: "/web/resources/assets/logo/logo-color.svg",
		},
		Styles: []string{
			"/web/resources/styles/tailwind/tailwind-min.css",
			"/web/resources/styles/loading.css",
		},
		CacheableResources: []string{
			"/web/resources/styles/tailwind/tailwind-min.css",
			"/web/resources/styles/loading.css",
			"/web/resources/assets/logo.svg",
			"/web/resources/assets/images/home_cover.png",
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
