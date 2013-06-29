package controller

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/microcosm-cc/gosnippet/helpers"
	"github.com/microcosm-cc/gosnippet/models"
	"io/ioutil"
	"net/http"
	"strings"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {

	c := LoginController{Context: helpers.CreateContext(w, r)}

	switch r.Method {
	case "OPTIONS":
		// Change this to encapsulate which methods are available through this handler
		c.Context.SendOptions([]string{"OPTIONS", "POST", "TRACE"})
	case "TRACE":
		c.Context.SendTrace()
	case "POST":
		c.Post()
	default:
		c.Context.ResponseWriter.Header().Set("Content-Type", "text/plain")
		c.Context.SendResponse(
			http.StatusMethodNotAllowed,
			[]byte(http.StatusText(http.StatusMethodNotAllowed)),
		)
	}
}

type LoginController struct {
	Context *helpers.Context
}

type PersonaAssertion struct {
	Assertion string `json:"assertion"`
	Audience  string `json:"audience"`
}

type PersonaResponse struct {
	Status   string `json:"status"`
	Email    string `json:"email"`
	Audience string `json:"audience"`
	Expires  int32  `json:"expires"`
	Issuer   string `json:"issuer"`
}

func (c *LoginController) Post() {
	c.Context.ResponseWriter.Header().Set("Content-Type", "text/plain")

	// Get the Persona assertion from the form data
	assertion := c.Context.Request.FormValue("assertion")
	if strings.Trim(assertion, " ") == "" {
		c.Context.SendError(http.StatusBadRequest, errors.New("assertion absent in POST data"))
		return
	}

	// Verify the Persona assertion
	// 1) Build post data
	pa := PersonaAssertion{
		Assertion: assertion,
		Audience:  helpers.Config.GetString("persona.audience", "http://localhost:8080"),
	}
	data, err := json.Marshal(pa)
	if err != nil {
		c.Context.SendError(http.StatusBadRequest, err)
		return
	}

	// 2) Verify with Persona
	resp, err := http.Post(
		"https://verifier.login.persona.org/verify",
		"application/json",
		strings.NewReader(string(data)),
	)
	if err != nil {
		c.Context.SendError(http.StatusInternalServerError, err)
		return
	}

	// 3) Read response
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Context.SendError(http.StatusInternalServerError, err)
		return
	}
	resp.Body.Close()
	personaResponse := PersonaResponse{}
	json.Unmarshal(body, &personaResponse)

	// 4) Check values in the response
	if personaResponse.Status != "okay" {
		c.Context.SendError(http.StatusInternalServerError, errors.New(
			fmt.Sprintf("Persona login error: %v", personaResponse.Status),
		))
		return
	}

	if personaResponse.Email == "" {
		c.Context.SendError(http.StatusInternalServerError, errors.New(
			"Persona error: no email address received",
		))
		return
	}

	// TODO: Insert into table
	sess := models.Session{Email: personaResponse.Email}
	sess.Insert()

	// TODO: set cookie
	// TODO: redirect back to target

	cookie := http.Cookie{
		Name:     "SessionId",
		Value:    sess.Id,
		Path:     "/",
		HttpOnly: true,
	}

	http.SetCookie(c.Context.ResponseWriter, &cookie)

	c.Context.SendStatus(http.StatusOK)

	return
}

func LogoutHandler(w http.ResponseWriter, r *http.Request) {

	c := LogoutController{Context: helpers.CreateContext(w, r)}

	switch r.Method {
	case "OPTIONS":
		// Change this to encapsulate which methods are available through this handler
		c.Context.SendOptions([]string{"OPTIONS", "POST", "HEAD", "TRACE"})
	case "TRACE":
		c.Context.SendTrace()
	case "POST":
		c.Post()
	default:
		c.Context.ResponseWriter.Header().Set("Content-Type", "text/plain")
		c.Context.SendResponse(
			http.StatusMethodNotAllowed,
			[]byte(http.StatusText(http.StatusMethodNotAllowed)),
		)
	}
}

type LogoutController struct {
	Context *helpers.Context
}

func (c *LogoutController) Post() {
	c.Context.ResponseWriter.Header().Set("Content-Type", "text/plain")

	c.Context.SendResponse(http.StatusOK, []byte("OK"))

	return
}
