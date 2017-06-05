package controllers

import (
	"database/sql"
	. "fmt"
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	"strconv"
)

type MainController struct {
	beego.Controller
}

type user struct {
	Id       int         `form:"-"`
	keywords interface{} `form:"keywords"`
}

type post struct {
	id        int
	titre     string
	soustitre string
	auteur    string
	contenu   string
	date      string
	href      string
}

func sqlconnect(a map[int]map[string]string, query string) {
	p := new(post)
	i := 0

	db, err := sql.Open("postgres", "user=postgres password=1924zheng. port=8888 dbname=Temp sslmode=disable")
	rows, err := db.Query(query)
	if err != nil {
		Println(err.Error())
		return
	}

	defer rows.Close()

	for rows.Next() {
		i += 1
		err := rows.Scan(&p.id, &p.titre, &p.soustitre, &p.auteur, &p.contenu, &p.date)
		if err != nil {
			Println(err)
			return
		}

		a[i] = make(map[string]string)
		a[i]["id"] = strconv.Itoa(p.id)
		a[i]["titre"] = p.titre
		a[i]["soustitre"] = p.soustitre
		a[i]["auteur"] = p.auteur
		a[i]["contenu"] = p.contenu
		a[i]["date"] = p.date
	}
}

func (b *MainController) Index() {
	index := make(map[int]map[string]string)

	query := "select * from articles order by 1 desc limit 3"
	sqlconnect(index, query)
	// u := user{}
	// if err := b.ParseForm(&u); err != nil {
	// }
	keywords := b.Input().Get("contexts")
	b.Data["Mapped"] = index
	b.Data["ID"] = []string{"id"}
	b.Data["Title"] = []string{"titre"}
	b.Data["Subtitle"] = []string{"soustitre"}
	b.Data["Name"] = []string{"auteur"}
	b.Data["Date"] = []string{"date"}
	b.Data["Text"] = keywords
	b.Data["addr"] = b.Ctx.Request.RemoteAddr
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

func (this *MainController) Search() {
	this.Data["keywords"] = this.Ctx.Input.Param(":keywords")
	this.TplName = "default/article.tpl"
}
