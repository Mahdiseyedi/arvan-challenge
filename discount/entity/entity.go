package entity

type User struct {
	ID          int64
	PhoneNumber string
	Balance     float64
}

type Transaction struct {
	ID     int64
	user   User
	amount float64
}
