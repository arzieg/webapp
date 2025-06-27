package domain

type User struct {
	ID        uint
	FirstName string
	LastName  string
	Emails    []Email
	LastIP    string
}

type Email struct {
	ID      uint
	Address string
	Primary bool
	UserID  int
}
