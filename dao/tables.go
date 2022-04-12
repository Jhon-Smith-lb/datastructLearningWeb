package dao

import (
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	// 需要在init中注册定义的model
	orm.RegisterModel(new(User))
}

type User struct {
	Id         int64
	Username   string
	Number     string
	Password   string
	Status     int16
	IsAdmin    int8
	CreatedAt  time.Time
	UpdatedAt  time.Time
	IsDel      int8
	CreateNews int8
}

func NewUser() *User {
	return &User{}
}

type QueryUserOption struct {
	Username string
	Number   string
	Offset   int64
	Limit    int64
}

func NewQueryUserOption() *QueryUserOption {
	return &QueryUserOption{}
}

type News struct {
	Id        int64     `json:"id"`
	Title     string    `json:"title"`
	UserId    int64     `json:"user_id"`
	Type      int16     `json:"type"`
	Text      string    `json:"text"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	IsDel     int8      `json:"is_del"`
	Offset    int64     `json:"offset"`
	Limit     int64     `json:"limit"`
}

func NewNews() *News {
	return &News{}
}

type QueryNewsOption struct {
	Title  string
	Offset int64
	Limit  int64
}

func NewQueryNewsOption() *QueryNewsOption {
	return &QueryNewsOption{}
}
