package controllers

import (
	"crypto/tls"
	"time"

	"github.com/astaxie/beego"
	"github.com/toolkits/net/httplib"
)

type GetBetaController struct {
	BaseController
}

func (c *GetBetaController) Get() {
	addr := beego.AppConfig.String("DestBetaAddr")
	req := httplib.Get(addr)
	req.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	req.SetTimeout(10*time.Second, 30*time.Second)
	currentbeta, err := req.String()
	if err != nil {
		c.SetJson(1, nil, "获取版本错误")
	} else {
		c.Data["currentbeta"] = currentbeta
		c.SetJson(0, currentbeta, "")
	}
	c.ServeJSON()
}
