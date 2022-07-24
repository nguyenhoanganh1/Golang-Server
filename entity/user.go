package entity

import "time"

type User struct {
	Username   string    `gorm : "primary-key; auto-increment" json:"username"`
	Password   string    `gorm : "type: varchar(150)" json : "password"`
	Fullname   string    `gorm : "type: varchar(150)" json : "fullname"`
	Email      string    `gorm : "type: varchar(150)" json : "email"`
	DOB        time.Time `json : "dob"`
	Activated  bool      `json : "activated"`
	CreateBy   string    `gorm : "type: varchar(150)" json : "createBy"`
	CreateTime time.Time `json : "createBy"`
	UpdateTime string    `gorm : "type: varchar(150)" json : "updateTime"`
	UpdateBy   time.Time `json : "updateBy"`
}

func (user *User) SetUsername(username string) {
	user.Username = username
}

func (user *User) GetUsername() string {
	return user.Username
}

func (user *User) SetPassword(password string) {
	user.Password = password
}

func (user *User) GetPassword() string {
	return user.Password
}

func (user *User) SetFullname(fullname string) {
	user.Fullname = fullname
}

func (user *User) GetFullname() string {
	return user.Fullname
}

func (user *User) SetEmail(email string) {
	user.Email = email
}

func (user *User) GetEmail() string {
	return user.Email
}
