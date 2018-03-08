package main

import (

	_ "beegolearn02/routers"
	"github.com/astaxie/beego"
)

func indexadd(i int) (i1 int)  {

	return i+1
}
func main() {
	beego.AddFuncMap("addindex",indexadd)
	beego.Run()
}

