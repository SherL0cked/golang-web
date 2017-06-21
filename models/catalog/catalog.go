package catalog

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	. "sitepointgoapp/models"
)

func ReadBlogCatalogs() []Catalog {
	var cat []Catalog
	_, err := Catalogs().OrderBy("DisplayOrder").All(&cat, "Id", "Name", "DisplayName", "DisplayOrder")

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return cat
}

func BlogCatalogsbyId(Id int64) []Catalog {
	var cat []Catalog
	_, err := Catalogs().Filter("Id", Id).OrderBy("DisplayOrder").All(&cat, "Id", "Name", "DisplayName", "DisplayOrder")

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return cat
}

func Save(this *Catalog) (int64, error) {
	id, err := orm.NewOrm().Insert(this)
	if err == nil {
		fmt.Println(id)
	}

	return id, err
}

func Del(this *Catalog) error {
	_, err := orm.NewOrm().Delete(this)
	if err != nil {
		return err
	}

	return nil
}

func Update(this *Catalog) error {
	if this.Id == 0 {
		return fmt.Errorf("primary key id not set")
	}
	_, err := orm.NewOrm().Update(this)

	if err == nil {
		fmt.Println(err)
	}
	return err
}

func Catalogs() orm.QuerySeter {
	return orm.NewOrm().QueryTable("Catalog")
}
