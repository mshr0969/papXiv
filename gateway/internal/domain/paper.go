package domain

type Paper struct {
	Id        string
	Published string
	Subject   string
	Title     string
	Url       string
	CreatedAt string
	UpdatedAt string
}

type Papers []Paper
