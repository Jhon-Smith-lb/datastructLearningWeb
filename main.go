package main

import (
	_ "dataStructLearningWeb/routers"

	"github.com/beego/beego/v2/core/logs"
	beego "github.com/beego/beego/v2/server/web"
)

func init() {
	// 配置日志输出引擎
	logs.SetLogger(logs.AdapterFile, `{"filename":"project.log","level":7,"maxlines":0,
	"maxsize":0,"daily":true,"maxdays":10,"color":true}`)
	// 输出文件名和行号
	logs.EnableFuncCallDepth(true)
	// 异步输出日志
	logs.Async(1e3)
}

func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}
	beego.Run()
}
