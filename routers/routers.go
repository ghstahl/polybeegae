package routers

import (
	"controllers"
	"github.com/astaxie/beegae"
)

func init() {

	v1NS := beegae.NewNamespace("/v1",
		beegae.NSNamespace("/object",
			beegae.NSRouter("/", &controllers.ObjectController{}, "get:GetAll;post:Post;put:Put"),
			beegae.NSRouter("/:objectId", &controllers.ObjectController{}, "get:Get;delete:Delete")))
	beegae.AddNamespace(v1NS)
	/*
   beegae.Router("/v1/object",&controllers.ObjectController{},"get:GetAll")
   beegae.Router("/v1/object",&controllers.ObjectController{},"post:Post")
   beegae.Router("/v1/object/:objectId",&controllers.ObjectController{},"get:Get")
   beegae.Router("/v1/object",&controllers.ObjectController{},"put:Put")
   beegae.Router("/v1/object/:objectId",&controllers.ObjectController{},"delete:Delete")
	*/
	beegae.Router("/_session_gc", &controllers.GCController{})
	beegae.Router("/", &controllers.MainController{})
}
