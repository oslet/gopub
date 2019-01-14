package controllers

import (
	"crypto/tls"
	"time"

	"github.com/astaxie/beego"
	"github.com/toolkits/net/httplib"
)

type GetReleaseController struct {
	BaseController
}

func (c *GetReleaseController) Get() {
	addr := beego.AppConfig.String("DestReleaseAddr")
	req := httplib.Get(addr)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(10*time.Second, 30*time.Second)
	currentrelease, err := req.String()
	if err != nil {
		c.SetJson(1, nil, "获取版本错误")
	} else {
		c.Data["currentrelease"] = currentrelease
		c.SetJson(0, currentrelease, "")
	}
	c.ServeJSON()
}
