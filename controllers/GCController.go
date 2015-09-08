package controllers

import "github.com/astaxie/beegae"

type GCController struct {
	beegae.Controller
}

func (this *GCController) Get() {
	beegae.GlobalSessions.GC(this.AppEngineCtx)
}


