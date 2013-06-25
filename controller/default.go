package controller

import (
	"github.com/microcosm-cc/gosnippet/helpers"
	"net/http"
)

func RootHandler(w http.ResponseWriter, r *http.Request) {

	c := RootController{Context: helpers.CreateContext(w, r)}

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

type RootController struct {
	Context *helpers.Context
}

func (c *RootController) Get() {

	c.Context.ResponseWriter.Header().Set("Content-Type", "text/html")

	type Data struct {
		Title string
		User  string
	}

	c.Context.SendTemplate(
		http.StatusOK,
		"home",
		Data{
			Title: "FooBar",
			User:  "",
		},
	)

	//	c.Context.SendResponse(http.StatusOK, []byte("OK"))

	return
}
