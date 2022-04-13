package controllers

import (
	"dataStructLearningWeb/biz/bizresources"
	"dataStructLearningWeb/dm"
	"dataStructLearningWeb/utils"
	"errors"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

type ResourcesController struct {
	beego.Controller
}

func (p *ResourcesController) Prepare() {
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

func (p *ResourcesController) QueryResourcesList() {
	prefix := p.GetString("prefix")
	limitStr := p.GetString("limit")

	logs.Info("[QueryResourcesList] prefix: %v, limit: %v", prefix, limitStr)

	var limit int64
	var err error
	if limitStr != "" {
		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			logs.Error("[QueryResourcesList] strconv.ParseInt, err: %v\n", err)
			p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("param is incorrect").Error())
			p.ServeJSON()
			p.StopRun()
		}
	}
	dmResourcesList, err := bizresources.QueryResourcesList(prefix, int(limit))
	if err != nil {
		logs.Error("[QueryResourcesList] bizresources.QueryResourcesList, err: %v, prefix: %v, limit: %v\n", err, prefix, limit)
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}

	data := map[string]interface{}{
		"resources_list": dmResourcesList,
	}
	p.Data["json"] = utils.SetResp(dm.HTTP_OK, data, "")
	p.ServeJSON()
}
