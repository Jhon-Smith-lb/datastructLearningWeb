package controllers

import (
	"dataStructLearningWeb/biz/biznews"
	"dataStructLearningWeb/dm"
	"dataStructLearningWeb/dm/dmnews"
	"dataStructLearningWeb/lib"
	"dataStructLearningWeb/utils"
	"encoding/json"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type NewsController struct {
	beego.Controller
}

func (p *NewsController) Prepare() {
	bearerToken := p.Ctx.Input.Header("Authorization")
	arr := strings.Split(bearerToken, " ")
	if len(arr) != 2 {
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("无效token").Error())
		p.ServeJSON()
		p.StopRun()
	}

	// 获取token
	token := arr[1]

	// 解析token
	claims, err := ParseToken(token)
	if err != nil {
		logs.Info("[Prepare] err: %v", err)
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("无效token").Error())
		p.ServeJSON()
		p.StopRun()
	}

	userId := claims.ID
	p.Data["user_id"] = userId
}

func (p *NewsController) AddNews() {
	var req dmnews.AddNewsReq
	json.Unmarshal(p.Ctx.Input.RequestBody, &req)

	req.UserId = p.Data["user_id"].(int64)

	logs.Info("[AddNews] u.Ctx.Input.RequestBody: %v\n", string(p.Ctx.Input.RequestBody))
	logs.Info("[AddNews] req: %v\n", lib.PointerToString(&req))

	newsId, err := biznews.AddNews(&req)
	if err != nil {
		logs.Error("[AddNews]  biznews.AddNews, err: %v\n", err)
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}
	data := map[string]interface{}{
		"news_id": newsId,
	}
	p.Data["json"] = utils.SetResp(dm.HTTP_OK, data, "")
	p.ServeJSON()
}

// @Title QueryUser
// @Description 根据过滤条件查询用户
// @Success 200 {object} models.User
func (p *NewsController) QueryNewsList() {
	title := p.GetString("title")
	newsTypeStr := p.GetString("type")
	offsetStr := p.GetString("offset")
	limitStr := p.GetString("limit")

	logs.Info("[QueryNewsList] title: %v, typeStr: %v, offsetStr: %v, limitStr: %v", title, newsTypeStr, offsetStr, limitStr)

	var offset int64
	var limit int64
	var newsType int64
	var err error

	if offsetStr != "" {
		offset, err = strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			logs.Error("[QueryNewsList] strconv.ParseInt, err: %v\n", err)
			p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("param is incorrect").Error())
			p.ServeJSON()
			p.StopRun()
		}
	}

	if limitStr != "" {
		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			logs.Error("[QueryNewsList] strconv.ParseInt, err: %v\n", err)
			p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("param is incorrect").Error())
			p.ServeJSON()
			p.StopRun()
		}
	}

	if newsTypeStr != "" {
		newsType, err = strconv.ParseInt(newsTypeStr, 10, 64)
		if err != nil {
			logs.Error("[QueryNewsList] strconv.ParseInt, err: %v\n", err)
			p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("param is incorrect").Error())
			p.ServeJSON()
			p.StopRun()
		}
	}

	bizReq := dmnews.NewQueryNewsOption()
	bizReq.Title = title
	bizReq.Type = dmnews.NewsType(newsType)
	bizReq.Offset = offset
	bizReq.Limit = limit

	dmNewsList, total, err := biznews.QueryNewsList(bizReq)
	if err != nil {
		logs.Error("[QueryNewsList] biznews.QueryNewsList, err: %v, bizReq: %v\n", err, lib.PointerToString(bizReq))
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}
	data := map[string]interface{}{
		"news_list": dmNewsList,
		"total":     total,
	}
	p.Data["json"] = utils.SetResp(dm.HTTP_OK, data, "")
	p.ServeJSON()
}

// @Title Update
// @Description update the user
// @Param	uid		path 	string	true		"The uid you want to update"
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {object} models.User
// @Failure 403 :uid is not int
func (p *NewsController) UpdateNews() {
	var req dmnews.UpdateNewsReq
	json.Unmarshal(p.Ctx.Input.RequestBody, &req)

	req.UserId = p.Data["user_id"].(int64)

	logs.Info("[UpdateNews] u.Ctx.Input.RequestBody: %v\n", string(p.Ctx.Input.RequestBody))
	logs.Info("[UpdateNews] req: %v\n", lib.PointerToString(&req))

	if err := biznews.UpdateNews(&req); err != nil {
		logs.Error("[UpdateNews] biznews.UpdateNews, err: %v\n", err)
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}

	p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, "")
	p.ServeJSON()
}
