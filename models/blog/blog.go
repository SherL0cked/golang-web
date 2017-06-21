package blog

import (
	"fmt"
	"github.com/astaxie/beego/orm"
	. "sitepointgoapp/models"
)

func ReadBlogSummary(catid int64, limit int64, offset int64) []Blog {
	var blogs []Blog
	_, err := Blogs().Filter("CatalogId", catid).Limit(limit, offset).OrderBy("-Created").All(&blogs, "Title", "Subtitle", "Created", "Updated", "Views")

	if err != nil {
		fmt.Println(err)
		return nil
	}

	return blogs
}

func ReadBlogContent(blogid int64) []BlogContent {
	var bc []BlogContent
	_, err := BlogContents().Filter("Id", blogid).All(&bc, "Content")

	if err != nil {
		fmt.Println(err)
		return nil
	}
	return bc
}

func Blogs() orm.QuerySeter {
	return orm.NewOrm().QueryTable("Blog")
}

func BlogContents() orm.QuerySeter {
	return orm.NewOrm().QueryTable("BlogContent")
}
