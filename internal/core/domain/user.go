package domain

type User struct {
	ID          int
	Version     int
	FullName    string
	PhoneNumber string
}

func NewUser(
	ID int,
	Version int,
	FullName string,
	PhoneNumber string,
) User {
	u := User{}
	u.ID = ID
	u.Version = Version
	u.FullName = FullName
	u.PhoneNumber = PhoneNumber
	return u
}
