package viewmodels

type AddUserViewModel struct {
	UserName string
	Email    string
	Password string
}

type SignInViewModel struct {
	UserName string
	Password string
}

type UserSignedInViewModel struct {
	UserName string
	Email    string
	Id       string
	JwtToken string
}

type UserViewModel struct {
	Id       string
	UserName string
	Email    string
}
