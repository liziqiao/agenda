package entities

// A User specifies a user entity
//
// ID - The id of the user
// Key - The api key of the user
// Username - The name of the user (unique for now)
// Password - The password of the user
// Email - The email address of the user
// Phone - The phone number of the user
type User struct {
	ID       int    `json:"id"`
	Key      string `json:"key"`
	Username string `json:"username"`
	Password string `json:"password"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

// NewUser returns a new instance of a user
func NewUser(username string, password string,
	email string, phone string) *User {
	return &User{-1, "", username, password, email, phone}
}
