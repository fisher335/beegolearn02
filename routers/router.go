package routers

import (
	"beegolearn02/controllers"
	"github.com/astaxie/beego"
)

func init() {
    beego.Router("/", &controllers.MainController{})
    beego.Router("/list/", &controllers.MainController{},"get:List")
    beego.Router("/test/", &controllers.MainController{},"get:Test")
    beego.Router("/login/", &controllers.MainController{},"get:Login")
    beego.Router("/file/", &controllers.MainController{},"get:File")
    beego.Router("/upload/", &controllers.MainController{},"get:Upload")
    beego.Router("/upload/", &controllers.MainController{},"post:UploadSave")
    beego.Router("/record/", &controllers.MainController{},"get:Record")

}
