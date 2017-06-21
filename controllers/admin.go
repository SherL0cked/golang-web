package controllers

import (
	"sitepointgoapp/models"
	"sitepointgoapp/models/catalog"
)

type AdminController struct {
	MainController
}

func (this *AdminController) extractCatalog() *models.Catalog {
	catalogs := &models.Catalog{}
	catalogs.Name = this.GetString("Name")
	catalogs.DisplayName = this.GetString("DisplayName")
	catalogs.DisplayOrder, _ = this.GetInt64("DisplayOrder")
	return catalogs
}

func (this *AdminController) ReadCatalogs() {
	this.Data["Catalogs"] = catalog.ReadBlogCatalogs()
	this.Layout = "layout/admin.html"
	this.TplName = "catalog/display.html"

}

func (this *AdminController) Add() {
	this.Data["IsAddCatalog"] = true
	this.Layout = "layout/admin.html"
	this.TplName = "catalog/add.html"
}

func (this *AdminController) DoAdd() {
	cat := this.extractCatalog()
	_, err := catalog.Save(cat)

	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.AdminRedirect()

}

func (this *AdminController) Del() {
	id := this.ReadBlogId()
	catalogs := &models.Catalog{}
	catalogs.Id = id
	err := catalog.Del(catalogs)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.AdminRedirect()
}

func (this *AdminController) Edit() {
	id := this.ReadBlogId()
	catalogs := catalog.BlogCatalogsbyId(id)
	this.Data["Catalog"] = catalogs
	this.Layout = "layout/admin.html"
	this.TplName = "catalog/edit.html"
}

func (this *AdminController) DoEdit() {
	cat := this.extractCatalog()
	cat.Id = this.ReadBlogId()

	err := catalog.Update(cat)
	if err != nil {
		this.Ctx.WriteString(err.Error())
		return
	}

	this.AdminRedirect()
}
