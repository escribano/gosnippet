package models

type Session struct {
	Id    string `json:"id"`
	Email string `json:"email"`
}

func (m *Session) Insert() {

}
