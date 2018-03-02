package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"github.com/astaxie/beego/cache"
	"math/rand"
	"path"
	"time"
)

var bm, err = cache.NewCache("memory", `{"interval":7000}`)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {

	c.Redirect("/list/",302)
	c.StopRun()
}

func (c *MainController) List() {

	head := make(map[string]interface{})
	bm.Put("token",GetToken(),20*time.Hour)
	for k, v := range c.Ctx.Request.Header {
		head[k] = v
	}
	head["url"] = c.Ctx.Request.URL
	head["methond"] = c.Ctx.Request.Method
	head["host"] = c.Ctx.Request.Host
	head["protl"] = c.Ctx.Request.Proto
	head["remote"] = c.Ctx.Request.RemoteAddr
	head["RequestURI"] = c.Ctx.Request.RequestURI
	//fmt.Println(head)
	head["token"]=bm.Get("token")
	if err != nil {
		fmt.Print("初始化内存缓存失败")
	}
	c.Data["head"] = head
	c.TplName = "list.tpl"
}
func (c *MainController) Test() {

	c.TplName = "test.tpl"
}
func (c *MainController) Login() {

	c.Data["pic"] = rand.Intn(10) + 1

	c.TplName = "login.tpl"
}

func (c *MainController) File() {

	file, _ := ListDir("static/videos")
	c.Data["file"] = file
	fmt.Println(bm.Get("token"))
	c.TplName = "file.tpl"
}

func (c *MainController) Upload() {

	c.TplName = "upload.tpl"
}

func (c *MainController) UploadSave() {

	f, h, _ := c.GetFile("file")
	fileName := h.Filename
	fmt.Println(fileName)
	f.Close()
	c.SaveToFile("file", path.Join("static/videos", fileName))
	c.Redirect("/file/", 302)
}

