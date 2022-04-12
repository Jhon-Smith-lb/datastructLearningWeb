package controllers

import (
	"dataStructLearningWeb/dm"
	"dataStructLearningWeb/utils"
	"errors"
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

