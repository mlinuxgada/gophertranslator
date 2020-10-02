package app

import (
	"net/http"
	"flag"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/sirupsen/logrus"
)

// Application struct - holding all generic components, needed for the app to function
type Application struct {
	IsInitialized bool
	Router *chi.Mux
	Logger *logrus.Logger
}

// App - global app instance
var App = Application{}

// GetApp - method, used to get the App instance OR instanciate it properly, if not yet so
func GetApp() Application {
	
	if false == App.IsInitialized {
		App.InitRouter()
		App.IsInitialized = true
		App.Logger = logrus.New()
	}

	return App
}

// InitRouter - initialize router - add whats needed
// PS - routes will be added later/after this
func (App *Application) InitRouter() {
	
	App.Router = chi.NewRouter()
	App.Router.Use(middleware.Logger)
	App.Router.Use(middleware.Recoverer)
}

// Start - method used to start the http server
func (App *Application) Start() {

	host := "localhost"
	port := flag.String("port", "8080", "b")

	flag.Parse()

	App.Logger.WithFields(logrus.Fields{}).Info("starting server on: " + host + ":" +*port)

	http.ListenAndServe(host + ":" + *port, App.Router)
}
