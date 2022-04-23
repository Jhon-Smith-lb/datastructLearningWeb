// @APIVersion 1.0.0
// @Title beego Test API
// @Description beego has a very cool tools to autogenerate documents for your API
// @Contact astaxie@gmail.com
// @TermsOfServiceUrl http://beego.me/
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"dataStructLearningWeb/controllers"

	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	ns :=
    beego.NewNamespace("/api",
      beego.NSNamespace("/v1",
		beego.NSNamespace("/login", 
				beego.NSRouter("/", &controllers.LoginController{}, "post:Login"),
		),
		beego.NSNamespace("/user",
			beego.NSRouter("/add", &controllers.UserController{}, "post:AddUser"),
			beego.NSRouter("/query", &controllers.UserController{}, "get:QueryUserList"),
			beego.NSRouter("/update", &controllers.UserController{}, "post:UpdateUser"),
			beego.NSRouter("/get_info", &controllers.UserController{}, "get:GetUserByToken"),
			beego.NSRouter("/reset_pwd", &controllers.UserController{}, "post:ResetPwd"),
			beego.NSRouter("/check_pwd", &controllers.UserController{}, "post:CheckPwd"),
		),
		beego.NSNamespace("/news", 
			beego.NSRouter("/add", &controllers.NewsController{}, "post:AddNews"),
			beego.NSRouter("/query", &controllers.NewsController{}, "get:QueryNewsList"),
			beego.NSRouter("/update", &controllers.NewsController{}, "post:UpdateNews"),
		),
		beego.NSNamespace("/resources",
			beego.NSRouter("/query", &controllers.ResourcesController{}, "get:QueryResourcesList"),
		),
      ),
  )
  beego.AddNamespace(ns)
}
