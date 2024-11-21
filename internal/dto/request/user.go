package request

type UserDTO struct {
	Username string
	Email    string
}

type UserCreateDTO struct {
	Username string
	Email    string
	Password string
}
