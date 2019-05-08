package controllers

import (
	"fmt"
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
	utils := models.Utils{}

	// Paser user input to user ojbect
	if err := c.ParseForm(&u); err != nil {
		fmt.Println("Fail to read from form")
	}

	// Check if website is empty
	if u.WebSite == "" {
		c.TplName = "index.tpl"
		return
	}

	// Split website into host & path
	split := strings.SplitAfterN(u.WebSite, "/", 2)
	var path string
	if len(split) > 1 {
		path = split[1]
		split[0] = split[0][:len(split[0])-1]
	}

	// Creating socket class object
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

	// Close socket
	defer s.SocketClient.Close()

	// Get html page
	response := s.GetWebPage()

	// Write response on find.html file
	if err := utils.WriteOnFile(response); err != nil {
		fmt.Println("Failed to render html on file")
		c.TplName = "error.html"
		return
	}

	c.TplName = "find.html"
}
