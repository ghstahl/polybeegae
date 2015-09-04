package beegoapp

import (   
	"github.com/astaxie/beegae"
    
)

var TheApp = &BeegoApp{}
type BeegoApp struct {
	// provider specifies the policy for authenticating a user.
	//TheAuthHandler *auth.AuthHandler

}



func (self *BeegoApp) Initialize() {

   
    beegae.SetStaticPath("/scripts","scripts")

}


func (self *BeegoApp) Run() {
   
    beegae.SetStaticPath("/bower_components","bower_components")
    beegae.SetStaticPath("/images","images")
    beegae.SetStaticPath("/elements","elements")
    beegae.SetStaticPath("/json","json")
    beegae.SetStaticPath("/styles","styles")
    beegae.SetStaticPath("/scripts","scripts")
    
	beegae.Run()
}