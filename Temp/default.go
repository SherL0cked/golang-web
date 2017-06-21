package controllers

import (
	"fmt"
	"github.com/astaxie/beego"
	"sitepointgoapp/models/blog"
	"strconv"
	"strings"
)

type MainController struct {
	beego.Controller
}

func urltranslate(inputs string) {
	m1 := map[string]string{"%2F": "/", "%20": " ", "+": "", "%3F": "?", "%25": "%", "%23": "#", "%26": "&", "%3D": "="}
	// Println(a.keys)
	for k, v := range m1 {
		inputs = (strings.Replace(inputs, k, v, -1))
	}
}

func (this *MainController) Index() {
	index := make(map[int]map[string]string)

	keywords := b.Input().Get("contexts")
	b.Data["BlogSummary"] = blog.
	b.Data["ID"] = []string{"id"}
	b.Data["Title"] = []string{"titre"}
	b.Data["Subtitle"] = []string{"soustitre"}
	b.Data["Name"] = []string{"auteur"}
	b.Data["Date"] = []string{"date"}
	b.Data["Text"] = keywords
	// b.Data["addr"] = b.Ctx.Request.RemoteAddr
	b.TplName = "default/index.tpl"
}

func (c *MainController) About() {
	c.TplName = "default/about.tpl"
}

func (a *MainController) Articles() {
	articles := make(map[int]map[string]string)
	page := a.Ctx.Input.Param(":id")
	pageint, _ := strconv.Atoi(page)
	limit := 5
	offset := (pageint - 1) * limit
	query := "select * from articles order by 1 desc limit " + strconv.Itoa(limit) + " offset " + strconv.Itoa(offset)

	// query := "select * from articles order by 1 desc limit 3"
	sqlconnect(articles, query)
	// a.Ctx.Request.PostForm
	a.Data["Mapped"] = articles
	a.Data["ID"] = []string{"id"}
	a.Data["Title"] = []string{"titre"}
	a.Data["Suatitle"] = []string{"soustitre"}
	a.Data["Name"] = []string{"auteur"}
	a.Data["Date"] = []string{"date"}
	// a.Data["Currenturl"] = querya
	switch {
	case pageint == 1:
		a.Data["Previous"] = pageint
		a.Data["Next"] = pageint + 1
	case pageint > 1:
		a.Data["Previous"] = pageint - 1
		a.Data["Next"] = pageint + 1
	}

	a.TplName = "default/article.tpl"
}

func (this *MainController) Contexts() {
	// tablename = "articles"
	contexts := make(map[int]map[string]string)
	articleid := this.Ctx.Input.Param(":id")
	query := "select * from articles where articleid = " + articleid

	sqlconnect(contexts, query)
	this.Data["Mapped"] = contexts
	this.Data["Constr"] = []string{"contenu"}
	this.Data["Queries"] = query

	this.TplName = "default/contexts.tpl"
}

func (t *MainController) Admin() {
	// tablename = "articles"
	keywords := t.Input().Get("Content")
	urltranslate(keywords)
	t.Data["test"] = keywords
	t.TplName = "default/admin.tpl"
}
