package routes

import (
	"github.com/mlinuxgada/gophertranslator/app"
	"github.com/mlinuxgada/gophertranslator/app/controllers/word"
	"github.com/mlinuxgada/gophertranslator/app/controllers/sentence"
	"github.com/mlinuxgada/gophertranslator/app/controllers/history"
)

// RegisterRoutes - method for adding/registering routes
func RegisterRoutes(App app.Application) {

		App.Router.Post("/word", word.Translate)
		App.Router.Post("/sentence", sentence.Translate)
		App.Router.Get("/history", history.GetAll)

}
