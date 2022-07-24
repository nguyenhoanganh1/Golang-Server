package request

type LoginRequest struct {
	Email    string `json : "username" form: "email" validate: "required, email"`
	Password string `json : "password" form: "password" validate: "required"`
}
