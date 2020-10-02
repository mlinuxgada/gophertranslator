package sentence

import (
	"encoding/json"
	"net/http"
	"strings"
	"github.com/mlinuxgada/gophertranslator/app"
	"github.com/mlinuxgada/gophertranslator/app/services/gophertranslator"
)

// Request - word request struct
type Request struct {
	EnglishSentence string `json:"english-sentence"`
}

// Response - word response struct
type Response struct {
	GopherSentence string `json:"gopher-sentence"`
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

	rawSentence := data.EnglishSentence
	punctChar := ""

	// Assume that every sentence ends with dot, question or exclamation mark.
	// but check it
	if strings.ContainsAny(data.EnglishSentence, ".?!") {

		punctChar = string(data.EnglishSentence[len(data.EnglishSentence) - 1])
		rawSentence = data.EnglishSentence[:len(data.EnglishSentence) - 1]
	}

	words := strings.Split(rawSentence, " ")
	gWords := []string{}

	for i, word := range words {

		word = strings.ToLower(word)

		gopherWord, err := gophertranslator.TranslateWord(word)

		if err != nil {
			app.RenderError(resp, err.Error(), 422)
			return
		}

		if i == 0 {
			gopherWord = strings.Title(gopherWord)
		}

		gWords = append(gWords, gopherWord)
	}

	gSentence := strings.Join(gWords, " ") + punctChar

	wordResponse := Response{
		GopherSentence: gSentence,
	}

	app.Render(resp, wordResponse)
}
