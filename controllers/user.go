package controllers

import (
	"dataStructLearningWeb/biz/bizuser"
	"dataStructLearningWeb/dm"
	"dataStructLearningWeb/dm/dmlogin"
	"dataStructLearningWeb/dm/dmuser"
	"dataStructLearningWeb/lib"
	"dataStructLearningWeb/utils"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
	"github.com/dgrijalva/jwt-go"
)

// Operations about Users
type UserController struct {
	beego.Controller
}

func (p *UserController) Prepare() {
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

	user, err := bizuser.GetUserById(claims.ID)
	if err != nil {
		logs.Error("[Prepare] bizuser.GetUserById, err: %v, claims.ID: %v", err, claims.ID)
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("用户不存在").Error())
		p.ServeJSON()
		p.StopRun()
	}

	if user.IsAdmin != dmuser.IS_ADMIN {
		urlPtr := p.Ctx.Request.URL
		urlStr := fmt.Sprintf("%v", urlPtr)
		urlArr := strings.Split(urlStr, "/")
		logs.Info("urlArr: %v, len: %v", urlArr, len(urlArr))
		if urlArr[4] == "add" || urlArr[4] == "update" || strings.HasPrefix(urlArr[4], "query") {
			logs.Error("[Prepare] bizuser.GetUserById, err: 不是管理员但访问了add, update, query三者中的一条路径")
			p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("没有权限").Error())
			p.ServeJSON()
			p.StopRun()
		}
	} 

	userId := claims.ID
	p.Data["user_id"] = userId
}

func ParseToken(token string) (*dmlogin.Claims, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &dmlogin.Claims{}, func(token *jwt.Token) (interface{}, error) {
	  return []byte(dmlogin.SIGNED_KEY), nil
	})
	if err != nil {
		logs.Error("[parseToken] err: %v", err)
	  	return nil, err
	}
  
	if tokenClaims == nil {
		logs.Error("[parseToken] err: tokenClaims == nil")
		return nil, errors.New("token无效")
	}

	claims, ok := tokenClaims.Claims.(*dmlogin.Claims)
	if !ok || !tokenClaims.Valid {
		logs.Error("[parseToken] err: !ok || tokenClaims.Valid, ok: %v, tokenClaims.Valid: %v", ok, tokenClaims.Valid)
		return nil, errors.New("token无效")
	} 
  
	return claims, nil
  }

// @Title AddUser
// @Description 添加用户
// @Param	body		body 	models.User	true		"body for user content"
// @Success 200 {int} dao.User.Id
// @Failure 403 body is empty
func (p *UserController) AddUser() {
	var req dmuser.AddUserReq
	json.Unmarshal(p.Ctx.Input.RequestBody, &req)

	req.UserId = p.Data["user_id"].(int64)

	logs.Info("[AddUser] u.Ctx.Input.RequestBody: %v\n", string(p.Ctx.Input.RequestBody))
	logs.Info("[AddUser] req: %v\n", req)

	userId, err := bizuser.AddUser(&req)
	if err != nil {
		logs.Error("[AddUser] bizuser.AddUser, err: %v\n", err)
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}
	data := map[string]interface{}{
		"user_id": userId,
	}
	p.Data["json"] = utils.SetResp(dm.HTTP_OK, data, "")
	p.ServeJSON()
}

// @Title QueryUser
// @Description 根据过滤条件查询用户
// @Success 200 {object} models.User
func (p *UserController) QueryUserList() {
	username := p.GetString("username")
	number := p.GetString("number")
	offsetStr := p.GetString("offset")
	limitStr := p.GetString("limit")
	var offset int64
	var limit int64
	var err error

	if offsetStr != "" {
		offset, err = strconv.ParseInt(offsetStr, 10, 64)
		if err != nil {
			logs.Error("[QueryUserList] strconv.ParseInt, err: %v\n", err)
			p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("param is incorrect").Error())
			p.ServeJSON()
			p.StopRun()
		}
	}
	
	if limitStr != "" {
		limit, err = strconv.ParseInt(limitStr, 10, 64)
		if err != nil {
			logs.Error("[QueryUserList] strconv.ParseInt, err: %v\n", err)
			p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, errors.New("param is incorrect").Error())
			p.ServeJSON()
			p.StopRun()
		}
	}

	bizReq := dmuser.NewQueryUserOption()
	bizReq.Username = username
	bizReq.Number = number
	bizReq.Offset = offset
	bizReq.Limit = limit

	dmUserList, total, err := bizuser.QueryUserList(bizReq)
	if err != nil {
		logs.Error("[QueryUserList] bizuser.QueryUserList, err: %v, bizReq: %v\n", err, lib.PointerToString(bizReq))
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}
	data := map[string]interface{}{
		"user_list": dmUserList,
		"total": total,
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
func (p *UserController) UpdateUser() {
	var req dmuser.UpdateUserReq
	json.Unmarshal(p.Ctx.Input.RequestBody, &req)

	req.UserId = p.Data["user_id"].(int64)

	logs.Info("[UpdateUser] u.Ctx.Input.RequestBody: %v\n", string(p.Ctx.Input.RequestBody))
	logs.Info("[UpdateUser] req: %v\n", req)

	if err := bizuser.UpdateUser(&req); err != nil {
		logs.Error("[UpdateUser] bizuser.UpdateUser, err: %v\n", err)
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}

	p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, "")
	p.ServeJSON()
}

func (p *UserController) GetUserByToken() {
	userId := p.Data["user_id"].(int64)
	bizUser, err := bizuser.GetUserById(userId)
	if err != nil {
		logs.Error("[GetUserById] bizuser.GetUserById, err: %v, userId: %v\n", err, userId)
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}

	controllerUser := NewUser()
	controllerUser.Username = bizUser.Username
	controllerUser.Number = bizUser.Number
	controllerUser.IsAdmin = bizUser.IsAdmin
	controllerUser.CreateNews = bizUser.CreateNews

	data := map[string]interface{}{
		"user": controllerUser,
	}

	p.Data["json"] = utils.SetResp(dm.HTTP_OK, data, "")
	p.ServeJSON()
}

func (p *UserController) ResetPwd() {
	var req dmuser.ResetPwdReq
	json.Unmarshal(p.Ctx.Input.RequestBody, &req)

	req.UserId =  p.Data["user_id"].(int64)

	logs.Info("[ResetPwd] req: %v", lib.PointerToString(&req))

	if err := bizuser.ResetPwd(&req); err != nil {
		logs.Error("[GetUserById] bizuser.GetUserById, err: %v, req: %v\n", err, lib.PointerToString(&req))
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}

	p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, "")
	p.ServeJSON()
}

func (p *UserController) CheckPwd() {
	var req dmuser.CheckPwdReq
	json.Unmarshal(p.Ctx.Input.RequestBody, &req)

	req.UserId =  p.Data["user_id"].(int64)

	logs.Info("[CheckPwd] req: %v", lib.PointerToString(&req))

	if err := bizuser.CheckPwd(&req); err != nil {
		logs.Error("[CheckPwd] bizuser.CheckPwd, err: %v, req: %v\n", err, lib.PointerToString(&req))
		p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, err.Error())
		p.ServeJSON()
		p.StopRun()
	}

	p.Data["json"] = utils.SetResp(dm.HTTP_OK, nil, "")
	p.ServeJSON()
}


