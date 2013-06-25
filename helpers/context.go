package helpers

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
)

type Context struct {
	Request        *http.Request
	ResponseWriter http.ResponseWriter
	RouteVars      map[string]string

	// This is a good place to put per request auth info if you need it,
	// such as a session identifier
}

func CreateContext(w http.ResponseWriter, r *http.Request) *Context {
	c := new(Context)
	c.ResponseWriter = w
	c.Request = r
	c.RouteVars = mux.Vars(r)
	return c
}

func (c *Context) SendResponse(status int, body []byte) {
	c.ResponseWriter.WriteHeader(status)

	if c.Request.Method != "HEAD" {
		c.ResponseWriter.Write(body)
	}
}

func (c *Context) SendStatus(status int) {
	c.SendResponse(status, []byte(http.StatusText(status)))
}

func (c *Context) SendOptions(options []string) {
	c.ResponseWriter.Header().Set("Allow", strings.Join(options, ","))
	c.ResponseWriter.Header().Set("Content-Length", "0")
	c.SendResponse(http.StatusOK, []byte{})
}

func (c *Context) SendTrace() {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain")

	dump, err := httputil.DumpRequest(c.Request, true)
	if err != nil {
		c.ResponseWriter.WriteHeader(http.StatusInternalServerError)
		c.ResponseWriter.Write([]byte(err.Error()))
	} else {
		c.ResponseWriter.WriteHeader(http.StatusOK)
		c.ResponseWriter.Write(dump)
	}
}

func (c *Context) SendError(status int, err error) {
	c.ResponseWriter.Header().Set("Content-Type", "text/plain")
	c.ResponseWriter.WriteHeader(status)
	c.ResponseWriter.Write([]byte(err.Error()))
}

func (c *Context) SendTemplate(status int, templateName string, data interface{}) {
	if _, exists := templates[templateName]; !exists {
		c.SendError(
			http.StatusInternalServerError,
			errors.New(fmt.Sprintf("helpers/context.go: Template %s not found", templateName)),
		)
		return
	}

	// Assumption that everything going through this function is HTML
	c.ResponseWriter.Header().Set("Content-Type", "text/html")

	body := new(bytes.Buffer)
	err := templates[templateName].ExecuteTemplate(body, "base", data)
	if err != nil {
		c.SendError(
			http.StatusInternalServerError,
			errors.New(
				fmt.Sprintf("helpers/context.go: Template %s render error: %s", templateName, err.Error()),
			),
		)
		return
	}

	// Prevents chunking and is required by HEAD
	c.ResponseWriter.Header().Set("Content-Length", strconv.Itoa(len(body.String())))

	if c.Request.Method != "HEAD" {
		c.ResponseWriter.Write(body.Bytes())
	}
}
