package entities

type Transfer struct {
	ID             string  `json:"id" db:"id"`
	SenderLogin    string  `json:"sender_login" db:"sender_login"`
	RecipientLogin string  `json:"recipient_login" db:"recipient_login"`
	Comment        string  `json:"comment" db:"comment"`
	Amount         float64 `json:"amount" db:"amount"`
	Date           string  `json:"date" db:"date"`
}
