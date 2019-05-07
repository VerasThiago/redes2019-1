package controllers

import (
	"fmt"
	"os"
	"projeto2/models"
	"strings"

	"github.com/astaxie/beego"
)

type UserController struct {
	beego.Controller
}

func (c *UserController) Get() {

	// Creating user input model
	u := models.User{}
	if err := c.ParseForm(&u); err != nil {
		fmt.Println("Fail to read from form")
	}

	if u.WebSite == "" {
		c.TplName = "index.tpl"
		return
	}

	// Split website into host & path
	split := strings.SplitAfterN(u.WebSite, "/", 2)
	path := ""
	if len(split) > 1 {
		path = split[1]
		split[0] = split[0][:len(split[0])-1]
	}

	// Creating socket class
	s := models.Socket{
		Ip:   split[0],
		Path: path,
		Port: 80,
	}
	// Stablshing conn
	if err := s.Init(); err != nil {
		fmt.Println("Failed to create a socket")
		fmt.Println(err)
		c.TplName = "error.html"
		return
	}

	response := s.GetWebPage()

	f, err := os.Create("views/find.html")
	if err != nil {
		fmt.Println(err)
	}
	l, err := f.WriteString(response)
	if err != nil {
		fmt.Println(err)
		f.Close()
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
	}

	c.TplName = "find.html"
}
