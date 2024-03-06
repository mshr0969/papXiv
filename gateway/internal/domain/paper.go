package domain

type Paper struct {
	Id        string `db:"id"`
	Published string `db:"published"`
	Subject   string
	Title     string `db:"title"`
	Url       string `db:"url"`
	CreatedAt string `db:"created_at"`
	UpdatedAt string `db:"updated_at"`
}

type Papers []Paper
