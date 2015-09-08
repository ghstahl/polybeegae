package routers

import (
	"controllers"
	"github.com/astaxie/beegae"
)

func init() {



beegae.Router("/v1/object",&controllers.ObjectController{},"*:GetAll")
	beegae.Router("/v1/object",&controllers.ObjectController{},"post:Post")
	beegae.Router("/v1/object/:objectId",&controllers.ObjectController{},"get:Get")
	beegae.Router("/v1/object",&controllers.ObjectController{},"put:Put")
	beegae.Router("/v1/object/:objectId",&controllers.ObjectController{},"delete:Delete")
	
	beegae.Router("/", &controllers.MainController{})
}
