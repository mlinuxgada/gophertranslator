package history

import (
	"github.com/mlinuxgada/gophertranslator/app"
	"github.com/mlinuxgada/gophertranslator/app/services/history"
	"net/http"
	"sort"
)

// Response - word response struct
type Response struct {
	History []map[string]string `json:"history"`
}

// GetAll - get all the history
func GetAll(resp http.ResponseWriter, req *http.Request) {

	app := app.GetApp()

	rawHistory := history.Get()
	sortedEngWords := make([]string, 0, len(rawHistory))

	for enWord := range rawHistory {
		sortedEngWords = append(sortedEngWords, enWord)
	}

	sort.Strings(sortedEngWords)

	histResp := Response{}

	for _, en := range sortedEngWords {

		m := make(map[string]string)
		m[en] = rawHistory[en]
		histResp.History = append(histResp.History, m)
	}

	app.Render(resp, histResp)
}
