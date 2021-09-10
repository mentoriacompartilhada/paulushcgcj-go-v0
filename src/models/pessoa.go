package models

type Pessoa struct {
	Id       int    `json:"id"`
	Nome     string `json:"nome"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Page struct {
	Page  int      `json:"page"`
	Size  int      `json:"size"`
	Total int      `json:"total"`
	Items []Pessoa `json:"items"`
}
