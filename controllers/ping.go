package controllers

import (
	"github.com/astaxie/beego"
	"encoding/xml"
	"github.com/deluan/gosonic/responses"
)

type PingController struct{ beego.Controller }

// @router /rest/ping.view [get]
func (this *PingController) Get() {
	response := responses.NewEmpty()
	xmlBody, _ := xml.Marshal(response)
	this.Ctx.Output.Body([]byte(xml.Header + string(xmlBody)))
}



