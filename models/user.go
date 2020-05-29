package models

type User struct {
	Id       string
	UserName string
	Email    string
}

//CreateUser creates new user
func CreateUser(id string, userName string, email string) *User {
	user := &User{
		Email:    email,
		Id:       id,
		UserName: userName}

	return user
}

//UpdateUser updates user
func (user *User) UpdateUser(userName string, email string) {
	user.Email = email
	user.UserName = userName
}
