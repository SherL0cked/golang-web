package routers

import (
	"github.com/astaxie/beego"
	_ "github.com/lib/pq"
	// "golang.org/x/net/context"
	"sitepointgoapp/controllers"
)

func init() {
	beego.Router("/", &controllers.BlogController{}, "get:Index")
	beego.Router("/articles&page:id:int", &controllers.BlogController{}, "get:Blogs")
	beego.Router("/journals&page:id:int", &controllers.BlogController{}, "get:Blogs")

	// beego.Router("/zfwadmin", &controllers.MeController{}, "get:Default")
	beego.Router("/zfwadmin/catalog/", &controllers.AdminController{}, "get:ReadCatalogs")
	beego.Router("/zfwadmin/catalog/add", &controllers.AdminController{}, "get:Add;post:DoAdd")
	beego.Router("/zfwadmin/catalog/:id:int/del", &controllers.AdminController{}, "get:Del")
	// beego.Router("/zfwadmin/catalog/edit", &controllers.AdminController{}, "post:DoEdit")
	beego.Router("/zfwadmin/catalog/:id:int/edit", &controllers.AdminController{}, "get:Edit;post:DoEdit")
}
