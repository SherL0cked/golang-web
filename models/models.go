package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/lib/pq"
	"time"
)

type Catalog struct {
	Id           int64 `orm:"auto"`
	Name         string
	DisplayName  string
	DisplayOrder int64 `orm:"unique"` //Navigation bar order
}

type Blog struct {
	Id            int64 `orm:"auto"`
	Title         string
	Subtitle      string
	Keywords      string       `orm:"null"`
	CatalogId     int64        `orm:"index"`
	Content       *BlogContent `orm:"-"`
	BlogContentId int64        `orm:"unique"`
	Views         int64
	Created       time.Time `orm:"auto_now_add;type(datetime)"`
	Updated       time.Time `orm:"auto_now;type(datetime)"`
}

type BlogContent struct {
	Id      int64
	Content string `orm:"type(text)"`
	Type    int64
}

func init() {
	orm.RegisterModel(new(Catalog), new(Blog), new(BlogContent))
}
