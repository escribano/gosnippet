package controller

import (
	"github.com/microcosm-cc/gosnippet/helpers"
	"net/http"
)

func SnippetHandler(w http.ResponseWriter, r *http.Request) {

	c := SnippetController{Context: helpers.CreateContext(w, r)}

	switch r.Method {
	case "OPTIONS":
		// Change this to encapsulate which methods are available through this handler
		c.Context.SendOptions([]string{"OPTIONS", "GET", "HEAD", "TRACE"})
	case "TRACE":
		c.Context.SendTrace()
	case "GET":
		c.Get()
	case "HEAD":
		c.Get()
	default:
		c.Context.ResponseWriter.Header().Set("Content-Type", "text/plain")
		c.Context.SendResponse(
			http.StatusMethodNotAllowed,
			[]byte(http.StatusText(http.StatusMethodNotAllowed)),
		)
	}
}

type SnippetController struct {
	Context *helpers.Context
}

func (c *SnippetController) Get() {
	c.Context.ResponseWriter.Header().Set("Content-Type", "text/plain")

	c.Context.SendResponse(http.StatusOK, []byte("OK"))

	return
}
