package entity

type User struct {
	Username string
	Password string
	IsLogin  bool
}

type UserResponse struct {
	Message string
}
