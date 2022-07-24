package request

import "time"

type RegisterRequest struct {
	Username   string    `json:"username"`
	Password   string    `json : "password"`
	Fullname   string    `json : "fullname"`
	Email      string    `json : "email"`
	DOB        time.Time `json : "dob"`
	Activated  bool      `json : "activated"`
	CreateBy   string    `json : "createBy"`
	CreateTime time.Time `json : "createBy"`
}
