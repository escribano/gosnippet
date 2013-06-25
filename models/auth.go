package models

import (
	"github.com/microcosm-cc/gosnippet/helpers"
)

type Session struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func (m *Session) Insert() {

}
