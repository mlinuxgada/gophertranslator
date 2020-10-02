package main

import (
	"github.com/mlinuxgada/gophertranslator/app"
	"github.com/mlinuxgada/gophertranslator/app/routes"
)

func main() {

	App := app.GetApp()

	routes.RegisterRoutes(App)

	App.Start()
}
