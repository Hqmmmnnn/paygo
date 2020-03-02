package entities

type User struct {
	ID         string `json:"id" db:"id"`
	FirstName  string `json:"first_name" db:"first_name"`
	LastName   string `json:"last_name" db:"last_name"`
	Patronymic string `json:"patronymic" db:"patronymic"`
}
