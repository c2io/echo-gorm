package model

type User struct {
	Id       int
	Username string
	Password string `gorm:"->:false;<-"`
}
