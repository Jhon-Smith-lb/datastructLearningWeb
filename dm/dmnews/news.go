package dmnews

import (
	"errors"
	"time"
)

type NewsType int16

const (
	NEWS     NewsType = 1 // 新闻动态
	RESEARCH NewsType = 2 // 教学研究的最新进展
	NOTICE   NewsType = 3 // 通知公告
)

type AddNewsReq struct {
	Title  string   `json:"title"`
	UserId int64    `json:"user_id"`
	Type   NewsType `json:"type"`
	Text   string   `json:"text"`
}

func (p *AddNewsReq) CheckParam() error {
	if p.Title == "" || p.Type == 0 {
		return errors.New("param is incorrect")
	}
	return nil
}

func NewAddNewsReq() *AddNewsReq {
	return &AddNewsReq{}
}

type UpdateNewsReq struct {
	Id     int64  `json:"id"`
	Title  string `json:"title"`
	UserId int64  `json:"user_id"`
	Type   int16  `json:"type"`
	Text   string `json:"text"`
	IsDel  int8   `json:"is_del"`
}

func NewUpdateNewsReq() *UpdateNewsReq {
	return &UpdateNewsReq{}
}

func (p *UpdateNewsReq) CheckParam() error {
	if p.Title == "" || p.Type == 0 {
		return errors.New("param is incorrect")
	}
	return nil
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
	Offset    *int64    `json:"offset,omitempty"`
	Limit     *int64    `json:"limit,omitempty"`
}

type QueryNewsOption struct {
	Title  string
	Offset int64
	Limit  int64
}

func (p *QueryNewsOption) CheckParam() error {
	if p.Limit == 0 {
		return errors.New("param is incorrect")
	}
	return nil
}

func NewQueryNewsOption() *QueryNewsOption {
	return &QueryNewsOption{}
}
