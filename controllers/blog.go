package controllers

import (
	"sitepointgoapp/models/blog"
	"sitepointgoapp/models/catalog"
	"strings"
)

type BlogController struct {
	MainController
}

var limit int64
var offset int64
var catid int64

func SurfPages(page int64) (int64, int64) {
	var Previous int64
	var Next int64
	switch {
	case page == 1:
		Previous = page
		Next = page + 1
	case page > 1:
		Previous = page - 1
		Next = page + 1
	}
	return Next, Previous
}

func ParseCatId(p *BlogController) int64 {
	current_url := p.Ctx.Request.RequestURI
	switch {
	case strings.Contains(current_url, "articles"):
		catid = 2
	case strings.Contains(current_url, "journals"):
		catid = 3
	}
	return catid
}

func (this *BlogController) Index() {
	limit = 3
	offset = 0
	this.Data["BlogContent"] = blog.ReadBlogSummary(2, limit, offset)
	this.Data["BlogCatalog"] = catalog.ReadBlogCatalogs()
	this.TplName = "default/index.tpl"
}

func (this *BlogController) Blogs() {
	limit = 5

	catid = ParseCatId(this)
	id := ParseCatId(this)

	offset = (int64(id) - 1) * limit
	Next, Previous := SurfPages(id)

	this.Data["BlogContent"] = blog.ReadBlogSummary(catid, limit, offset)
	this.Data["Previous"] = Previous
	this.Data["Next"] = Next
	this.TplName = "default/index.tpl"
}

/*
func (this *MainController) Contexts() {
	blogid := strconv.Atoi(this.Ctx.Input.Param(":id"))
	this.Data["BlogContext"] = blog.ReadBlogContent(blogid)
	this.TplName = "default/about.tpl"
}
*/
