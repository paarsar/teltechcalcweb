package routers

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/context/param"
)

func init() {

	beego.GlobalControllerRouter["teltechcalcweb/controllers:ArithmeticController"] = append(beego.GlobalControllerRouter["teltechcalcweb/controllers:ArithmeticController"],
		beego.ControllerComments{
			Method: "Add",
			Router: `/:action`,
			AllowHTTPMethods: []string{"get"},
			MethodParams: param.Make(),
			Params: nil})

}
