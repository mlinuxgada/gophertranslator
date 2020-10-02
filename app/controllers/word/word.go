package word

import (
	"net/http"
	"encoding/json"
	"github.com/mlinuxgada/gophertranslator/app"
	"github.com/mlinuxgada/gophertranslator/app/services/gophertranslator"
	"github.com/mlinuxgada/gophertranslator/app/services/history"
)

// Request - word request struct
type Request struct {
	EnglishWord string `json:"english-word"`
}

// Response - word response struct
type Response struct {
	GopherWord string `json:"gopher-word"`
}

// Translate - http controller method
func Translate(resp http.ResponseWriter, req *http.Request) {

	app := app.GetApp()
	var data Request
	decoder := json.NewDecoder(req.Body)

	err := decoder.Decode(&data)

	if err != nil {
		app.RenderError(resp, "Data not properly defined", 422)
		return
	}

	gopherWord, err := gophertranslator.TranslateWord(data.EnglishWord)

	if err != nil {
		app.RenderError(resp, err.Error(), 422)
		return
	}

	wordResponse := Response{
		GopherWord: gopherWord,
	}

	history.Add(data.EnglishWord, gopherWord)

	app.Render(resp, wordResponse)
}
