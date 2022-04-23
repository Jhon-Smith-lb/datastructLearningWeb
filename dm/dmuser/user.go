package dmuser

import (
	"errors"
	"time"
)

const (
	IS_ADMIN = 1
)

type User struct {
	Id         int64		`json:"id"`
	Username   string		`json:"username"`
	Number     string		`json:"number"`
	Password   string		`json:"password"`
	Status     int16		`json:"status"`
	IsAdmin    int8			`json:"is_admin"`
	CreatedAt  time.Time	`json:"created_at"`
	UpdatedAt  time.Time	`json:"updated_at"`
	IsDel      int8			`json:"is_del"`
	CreateNews int8			`json:"create_news"`
}

func NewUser() *User {
	return &User{}
}

type AddUserReq struct {
	UserId     int64  `json:"user_id"`
	Username   string `json:"username"`
	Number     string `json:"number"`
	Password   string `json:"password"`
	Status     int16  `json:"status"`
	IsAdmin    int8   `json:"is_admin"`
	CreateNews int8   `json:"create_news"`
}

func (p *AddUserReq) CheckParam() error {
	if p.Username == "" || p.Number == "" || p.Password == "" {
		return errors.New("param is incorrect")
	}
	return nil
}

type QueryUserOption struct {
	Username string `json:"username"`
	Number   string `json:"number"`
	Offset   int64  `json:"offset"`
	Limit    int64  `json:"limit"`
}

func (p *QueryUserOption) CheckParam() error {
	if p.Limit == 0 {
		return errors.New("param is incorrect")
	}
	return nil
}

func NewQueryUserOption() *QueryUserOption {
	return &QueryUserOption{}
}

type UpdateUserReq struct {
	Id         int64  `json:"id"`
	Username   string `json:"username"`
	Number     string `json:"number"`
	Password   string `json:"password"`
	Status     int16  `json:"status"`
	IsAdmin    int8   `json:"is_admin"`
	IsDel      int8   `json:"is_del"`
	CreateNews int8   `json:"create_news"`
	UserId     int64  `json:"user_id"`
}

func (p *UpdateUserReq) CheckParam() error {
	if p.Id == 0 || p.Username == "" || p.Number == "" {
		return errors.New("param is incorrect")
	}
	return nil
}

type ResetPwdReq struct {
	UserId int64 `json:"user_id"`
	OldPwd string `json:"old_pwd"`
	NewPwd string `json:"new_pwd"`
}

func (p *ResetPwdReq) CheckParam() error {
	if p.UserId == 0 || p.OldPwd == "" || p.NewPwd == "" {
		return errors.New("param is incorrect")
	}
	return nil
}

type CheckPwdReq struct {
	UserId int64 `json:"user_id"`
	Pwd string `json:"pwd"`
}

func (p *CheckPwdReq) CheckParam() error {
	if p.Pwd == "" {
		return errors.New("param is incorrect")
	}
	return nil
}