package domain

type User struct {
	ID       string
	Email    string
	FullName string `db:"full_name"`
}
