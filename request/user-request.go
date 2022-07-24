package request

import (
	"time"
)

type UserRequest struct {
	Username   string    `json : "username" form:"username" binding: "required" validate: "required"`
	Password   string    `json : "password" form:"password" `
	Fullname   string    `json : "fullname" form:"fullname"`
	Email      string    `json : "email" form:"email" validate: "required, email"`
	DOB        time.Time `json : "dob" form:"dob"`
	Activated  bool      `json : "activated" form:"activated" default: "false"`
	CreateBy   string    `json : "createBy" form:"createBy"`
	CreateTime time.Time `json : "createTime" form:"CreateTime"`
}

// func (user *UserRequest) SetUsername(username string) {
// 	user.Username = username
// }

// func (user *UserRequest) GetUsername() string {
// 	return user.Username
// }

// func (user *UserRequest) SetPassword(password string) {
// 	user.Password = password
// }

// func (user *UserRequest) GetPassword() string {
// 	return user.Password
// }

// func (user *UserRequest) SetFullname(fullname string) {
// 	user.Fullname = fullname
// }

// func (user *UserRequest) GetFullname() string {
// 	return user.Fullname
// }

// func (user *UserRequest) SetEmail(email string) {
// 	user.Email = email
// }

// func (user *UserRequest) GetEmail() string {
// 	return user.Email
// }
