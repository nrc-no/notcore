package api

type Country struct {
	ID   string `db:"id"`
	Code string `db:"code"`
	Name string `db:"name"`
}