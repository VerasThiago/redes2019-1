package controllers

import (
	"encoding/json"
	"fmt"
	"projeto2/models"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {
	u := models.User{}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println("Fail to read from form")
	}
	s := models.Socket{
		Ip:   "codeforces.com",
		Port: 80,
	}
	s.Init()

	response := s.GetMessage(u.Handle)
	var dat map[string]interface{}
	if err := json.Unmarshal([]byte(response), &dat); err != nil {
		fmt.Println(err)
	} else {
		c.Data["photo"] = dat["avatar"]
		c.Data["firstName"] = dat["firstName"]
		c.Data["handle"] = dat["handle"]
		c.Data["city"] = dat["city"]
		c.Data["organization"] = dat["organization"]
		c.Data["maxRank"] = dat["maxRank"]
	}
	c.TplName = "index.tpl"
}
