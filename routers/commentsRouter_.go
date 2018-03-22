package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["beegolearn02/controllers:MainController"] = append(beego.GlobalControllerRouter["beegolearn02/controllers:MainController"],
		beego.ControllerComments{
			Method: "RedirGithub",
			Router: `/`,
			AllowHTTPMethods: []string{"get","post"},
			MethodParams: param.Make(),
			Params: nil})

}
