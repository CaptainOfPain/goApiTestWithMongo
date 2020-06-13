package models

import "golang.org/x/crypto/bcrypt"

//User represents user model
type User struct {
	Id           string
	UserName     string
	Email        string
	PasswordHash string
}

//CreateUser creates new user
func CreateUser(id string, userName string, email string, password string) *User {
	user := &User{
		Email:    email,
		Id:       id,
		UserName: userName}
	user.setUserPassword(password)

	return user
}

//UpdateUser updates user
func (user *User) UpdateUser(userName string, email string) {
	user.Email = email
	user.UserName = userName
}

//ComparePassword checks user password hash with provided password
func (user *User) ComparePassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(password))
	return err == nil
}

func (user *User) setUserPassword(password string) {
	hash, _ := bcrypt.GenerateFromPassword([]byte(password), 15)
	user.PasswordHash = string(hash)
}
