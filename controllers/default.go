package controllers

import (
	"github.com/astaxie/beego"
	"fmt"
	"math/rand"
	"path"
	"github.com/astaxie/beego/session"
	"github.com/levigross/grequests"
	"encoding/json"
)

type MainController struct {
	beego.Controller
}

var globalSessions *session.Manager

func Init() {
	sessionConfig := &session.ManagerConfig{
		CookieName:      "gosessionid",
		EnableSetCookie: true,
		Gclifetime:      3600,
		Maxlifetime:     3600,
		Secure:          false,
		CookieLifeTime:  360000,
		ProviderConfig:  "./tmp",
	}
	globalSessions, _ = session.NewManager("memory", sessionConfig)
	go globalSessions.GC()
}

func (c *MainController) Get() {
	Init()

	var bm, err = globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer bm.SessionRelease(c.Ctx.ResponseWriter)
	fmt.Println(bm.Get("token"))
	if bm.Get("token") == nil {

		c.Ctx.Redirect(302, "/login/")
	}
	var t = bm.Get("token")
	var token string
	if value, ok := t.(string); ok {
		token = value
	}
	paras := &grequests.RequestOptions{Params: map[string]string{"start": "0", "count": "10", "method": "mobile_list"}, Headers: map[string]string{"Authorization": token}}
	res, err := grequests.Get(url, paras)
	fmt.Println(url)
	if err != nil {
		fmt.Println(err)
	}
	var record Record
	//fmt.Println(res.String())
	json.Unmarshal(res.Bytes(), &record)
	//fmt.Println(record.Rows)
	c.Data["rows"] = record.Rows
	c.TplName = "record.tpl"
}

func (c *MainController) Record() {
	c.Redirect("/list/", 302)
	c.StopRun()

}

func (c *MainController) List() {
	var bm, err = globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	defer bm.SessionRelease(c.Ctx.ResponseWriter)

	head := make(map[string]interface{})

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
	//head["token"] = bm.Get("token")
	head["token"] = bm.Get("token")
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
	//var bm, _ = globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	file, _ := ListDir("static/videos")
	c.Data["file"] = file
	//fmt.Println(bm.Get("token"))
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

func (c *MainController) LoginPost() {
	var bm, _ = globalSessions.SessionStart(c.Ctx.ResponseWriter, c.Ctx.Request)
	var username = c.GetString("user")
	var pass = c.GetString("passwd")
	var token, err = GetToken(username, pass)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(token)
	bm.Set("token", token)
	c.Ctx.Redirect(302, "/")
}
