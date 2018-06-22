package routers

import (
	c "teltechcalcweb/controllers"

	"github.com/astaxie/beego"
)

func init() {
	beego.Include(&c.ArithmeticController{})
}
