package server

import (
	"github.com/microcosm-cc/gosnippet/controller"
	"github.com/microcosm-cc/gosnippet/helpers"
	"net/http"
)

func rootStaticFile(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, helpers.Config.GetString("directories.static", "./static/")+r.URL.Path)
}

var (
	handlers = map[string]func(http.ResponseWriter, *http.Request){
		"/":                    controller.RootHandler,
		"/robots.txt":          rootStaticFile,
		"/favicon.ico":         rootStaticFile,
		"/auth/login":          controller.LoginHandler,
		"/auth/logout":         controller.LogoutHandler,
		"/snippets":            controller.SnippetsHandler,
		"/snippet/{id:[0-9]+}": controller.SnippetHandler,
	}
)
