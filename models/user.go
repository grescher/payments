package models

type User struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	IsActive     bool   `json:"is_active"`
}

func (u *User) Validate() (err error) {

	return nil
}
