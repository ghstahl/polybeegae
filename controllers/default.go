// Beego (http://beego.me/)
// @description beego is an open-source, high-performance web framework for the Go programming language.
// @link        http://github.com/astaxie/beego for the canonical source repository
// @license     http://github.com/astaxie/beego/blob/master/LICENSE
// @authors     Unknwon

package controllers

import (
	"github.com/astaxie/beegae"
	"fmt"
    "html/template"
     "sync" 
    "appengine"
    "datautils"
    "net/http"
	 
)
var startOnce sync.Once 

type MainController struct {
	beegae.Controller
}

func (this *MainController) Prepare() {
	fmt.Println("func (c *MainController) Prepare()")
	this.LayoutSections = make(map[string]string)
	this.TplNames = "index.tpl"
	this.LayoutSections["SharedHead"] = "shared/_head.tpl"
	this.LayoutSections["Header"] = "shared/_header.tpl"
	this.LayoutSections["Footer"] = "shared/_footer.tpl"
	this.Data["Jumbotron"] = true
	this.LayoutSections["HtmlHead"] = ""
	this.Layout = "shared/_layout.tpl"
	this.Data["xsrf_token"] = this.XsrfToken()
	this.Data["xsrfdata"]=template.HTML(this.XsrfFormHtml())
}

func start(w http.ResponseWriter,c appengine.Context) {   

    ensure_data.Sign(c,w)
    // startup code here; datastore requests, etc 
    ensure_data.EnsureData(w,c)
       

} 


func (this *MainController) Get() {
    r := this.Ctx.Request
	w := this.Ctx.ResponseWriter
    ctx  := appengine.NewContext(r) 
    
    startOnce.Do(func() { 
        ctx.Infof("startOnce")
        start(w,ctx) 
    }) 
	ctx.Infof("Serving the front page.")


	this.Data["Website"] = "beego.me"
	this.Data["Email"] = "astaxie@gmail.com"
	this.TplNames = "index.tpl"
}
