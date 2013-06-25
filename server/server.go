package server

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/microcosm-cc/gosnippet/helpers"
	"github.com/robfig/cron"
	"log"
	"net/http"
)

func StartServer() {

	// Load the config file
	err := helpers.ParseConfig()
	if err != nil {
		log.Fatal("server/server.go: ", err)
	}

	// Now we've loaded the config, we can load the templates from the
	// template directory
	helpers.LoadTemplates()

	// Set up the cron jobs if there are any
	c := cron.New()
	if len(tasks) > 0 {
		for schedule, task := range tasks {
			c.AddFunc(schedule, task)
		}
		c.Start()
	}

	// Add the URL handlers
	r := mux.NewRouter()
	for url, handler := range handlers {
		r.HandleFunc(url, handler)
	}

	// Add the static handler
	// Serves any /static/ request with the contents of the directory
	// specified in the config file
	r.PathPrefix("/static/").Handler(
		http.StripPrefix(
			"/static/",
			http.FileServer(
				http.Dir(helpers.Config.GetString("directories.static")),
			),
		),
	)

	http.Handle("/", r)

	// Start the HTTP server
	log.Fatal(
		http.ListenAndServe(
			fmt.Sprintf("%s:%d",
				helpers.Config.GetString("listen.interface", ""),
				helpers.Config.GetInt64("listen.port", 8080),
			),
			nil,
		),
	)
}
