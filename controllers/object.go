package controllers

import (
	"appengine"
	"encoding/json"
	"github.com/astaxie/beegae"
    "fmt"
    "models"
    "net/http"
	"stores"
)

// Operations about object
type ObjectController struct {
	beegae.Controller
}

// @Title create
// @Description create object
// @Param	body		body 	models.Object	true		"The object content"
// @Success 200 {string} models.Object.Id
// @Failure 403 body is empty
// @router /v1/object/ [post]
func (o *ObjectController) Post() {
    fmt.Println("ObjectController POST")
    r := o.Ctx.Request
	w := o.Ctx.ResponseWriter
	c := appengine.NewContext(r)
	c.Infof("ObjectController GetAll")
    
    var prod models.Product
	json.Unmarshal(o.Ctx.Input.RequestBody, &prod)
    res2B, _ := json.Marshal(prod)
    fmt.Println(string(res2B))
    
    
    err := productStore.InsertProduct(c, prod)
    if(err != nil){
        http.Error(w, err.Error(), http.StatusInternalServerError)
    }
    
	o.Data["json"] = prod
	o.ServeJson()
}

// @Title Get
// @Description find object by objectid
// @Param	objectId		path 	string	true		"the objectid you want to get"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /v1/object/:objectId [get]
func (o *ObjectController) Get() {

	r := o.Ctx.Request
	c := appengine.NewContext(r)
	c.Infof("ObjectController GetAll")

	classes := productStore.GetAllClassesFromProds(c)

	//products := ensure_data.GetAllProds(ctx)
	o.Data["json"] = classes

	/*
			objectId := o.Ctx.Input.Params[":objectId"]

		    ctx.Infof("ObjectController objectId:%v",objectId)

			if objectId != "" {
				ob, err := models.GetOne(objectId)
				if err != nil {
					o.Data["json"] = err
				} else {
					o.Data["json"] = ob
				}
				ctx.Infof("ObjectController ob:%v",ob)
			}
	*/
	o.ServeJson()
}

// @Title GetAll
// @Description get all objects
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /v1/object/ [get]
func (o *ObjectController) GetAll() {
        fmt.Println("ObjectController GetAll")

	r := o.Ctx.Request
	c := appengine.NewContext(r)
	c.Infof("ObjectController GetAll")

	var item_list []models.Product
	item_list = productStore.GetAllProds(c)

	o.Data["json"] = item_list

	//	obs := models.GetAll()
	//	o.Data["json"] = obs
	o.ServeJson()
}

// @Title update
// @Description update the object
// @Param	objectId		path 	string	true		"The objectid you want to update"
// @Param	body		body 	models.Object	true		"The body"
// @Success 200 {object} models.Object
// @Failure 403 :objectId is empty
// @router /v1/object/:objectId [put]
func (o *ObjectController) Put() {
	objectId := o.Ctx.Input.Params[":objectId"]
	var ob models.Object
	json.Unmarshal(o.Ctx.Input.RequestBody, &ob)

	err := models.Update(objectId, ob.Score)
	if err != nil {
		o.Data["json"] = err
	} else {
		o.Data["json"] = "update success!"
	}
	o.ServeJson()
}

// @Title delete
// @Description delete the object
// @Param	objectId		path 	string	true		"The objectId you want to delete"
// @Success 200 {string} delete success!
// @Failure 403 objectId is empty
// @router /v1/object/:objectId [delete]
func (o *ObjectController) Delete() {
	objectId := o.Ctx.Input.Params[":objectId"]
	models.Delete(objectId)
	o.Data["json"] = "delete success!"
	o.ServeJson()
}
