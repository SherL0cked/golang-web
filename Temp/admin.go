package controllers

import (
	"github.com/astaxie/beego"
	"sitepointgoapp/models"
	"sitepointgoapp/models/catalog"
)

type AdminController struct {
	beego.Controller
}

func (this *AdminController) ReadCatalogs() {
	this.Data["Catalogs"] = catalog.ReadBlogCatalogs()
}

func (this *AdminController) extractCatalog() *models.Catalog {
	catalogs := &models.Catalog{}
	catalogs.Name = this.GetString("Name")
	catalogs.DisplayName = this.GetString("DisplayName")
	catalogs.DisplayOrder = this.GetInt64("DisplayOrder")

	return catalogs
}

func (this *AdminController) Add() {
	cat := this.extractCatalog()
	_, err := catalog.save(cat)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
	this.Layout = "layout/admin.html"
	this.TplName = "catalog/add.html"

}

func (this *AdminController) Edit() {
	cat := this.extractCatalog()
	_, err := catalog.save(cat)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}
}
