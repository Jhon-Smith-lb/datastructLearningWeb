package controllers

import (
	"dataStructLearningWeb/biz/bizlogin"
	"dataStructLearningWeb/dm"
	"dataStructLearningWeb/dm/dmlogin"
	"dataStructLearningWeb/utils"
	"encoding/json"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

// Operations about Users
type LoginController struct {
	beego.Controller
}

// @Title Login
// @Description Logs user into the system
// @Param	username		query 	string	true		"The username for login"
// @Param	password		query 	string	true		"The password for login"
// @Success 200 {string} login success
// @Failure 403 user not exist
// @router /login [get]
func (p *LoginController) Login() {
	var req dmlogin.LoginReq
	json.Unmarshal(p.Ctx.Input.RequestBody, &req)

	logs.Info("[Login] u.Ctx.Input.RequestBody: %v\n", string(p.Ctx.Input.RequestBody))
	logs.Info("[Login] req: %v\n", req)

	token, err := bizlogin.Login(req.Number, req.Password)
	if err != nil {
		logs.Error("[Login] bizuser.Login, err: %v, number: %v", err, req.Number)
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}

	data := map[string]interface{}{
		"token": token,
	}
	p.Data["json"] = utils.SetResp(dm.HTTP_OK, data, "")
	p.ServeJSON()
}