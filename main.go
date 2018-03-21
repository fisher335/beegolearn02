package main

import (
	_ "beegolearn02/routers"
	"github.com/astaxie/beego"
)

func addindex(i int) int {

	return i + 1
}
func main() {
	beego.AddFuncMap("addindex", addindex)
	beego.Run()
}
