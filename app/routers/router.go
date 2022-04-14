package routers

import (
	"app/controllers"
	"app/filters"

	"github.com/astaxie/beego"
)

func init() {
	// ログ関係
	beego.InsertFilter("/*", beego.BeforeRouter, filters.LogManager)

	beego.Router("/", &controllers.MainController{})

	beego.Router("/bbs", &controllers.BbsController{}, "get:Show")
	beego.Router("/postup", &controllers.BbsController{}, "post:Post")
}
