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

}

func (self *BeegoApp) Run() {

	beegae.Run()
}
