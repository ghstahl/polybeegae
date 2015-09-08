package ensure_data

import (
    "net/http"
    "appengine"
    "appengine/datastore"
    "time"
)

// [START Product_struct]
type Product struct {
    Height string
    HardinessZone  string
    BloomTime string
    ImageUrl  string
    Description string
    Name  string
    Class  string
    PriceEach  string
    Light  string
    PlugSize  string
}
// [END Product_struct]


/*
 height string
        hardiness_zone  string
        bloom_time string
        image_url  string
        description string
        name  string
        class  string
        price_each  string
        light  string
        plug_size  string
*/

// [START greeting_struct]
type Greeting struct {
        Author  string
        Content string
        Date    time.Time
}
// [END greeting_struct]

// productKey returns the key used for all product entries.
func productKey(c appengine.Context) *datastore.Key {
    // The string "default_products" here could be varied to have multiple product catalogs.
    return datastore.NewKey(c, "Product", "default_products", 0, nil)
}
// guestbookKey returns the key used for all guestbook entries.
func guestbookKey(c appengine.Context) *datastore.Key {
        // The string "default_guestbook" here could be varied to have multiple guestbooks.
        return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}
// [START func_sign]
func InsertProduct(c appengine.Context,prod Product,w http.ResponseWriter) {
       
        
        // We set the same parent key on every Greeting entity to ensure each Greeting
        // is in the same entity group. Queries across the single entity group
        // will be consistent. However, the write rate to a single entity group
        // should be limited to ~1/second.
        key := datastore.NewIncompleteKey(c, "Product", productKey(c))
        _, err := datastore.Put(c, key, &prod)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
}

// [START func_sign]
func Sign(c appengine.Context,w http.ResponseWriter) {
        g := Greeting{
                Content: "Some Comment",
                Date:    time.Now(),
        }
        
        // We set the same parent key on every Greeting entity to ensure each Greeting
        // is in the same entity group. Queries across the single entity group
        // will be consistent. However, the write rate to a single entity group
        // should be limited to ~1/second.
        key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
        _, err := datastore.Put(c, key, &g)
        if err != nil {
                http.Error(w, err.Error(), http.StatusInternalServerError)
                return
        }
}

// [END func_sign]

func EnsureData(w http.ResponseWriter,c appengine.Context) {
    q := datastore.NewQuery("Product").Ancestor(productKey(c)).Limit(1)
    products := make([]Product, 0, 1)
    if _, err := q.GetAll(c, &products); err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
    }
    var prodExists bool = false 
   
    // if you only need e:
    for _, e := range products {
        c.Infof("Product:%v",e)
        prodExists = true
        break
    }
    
    if (!prodExists){
        prod := new(Product)
            prod.Height = "6 - 8 inches"
            prod.HardinessZone = "4 - 9"
            prod.BloomTime = "May - June"
            prod.ImageUrl = "http://www.jamesgreenhouses.com/_code/_dyn_images/products/Strawberry-Gasana.JPG"
            prod.Description = "Everbearing. Great for containers! Or in the ground."
            prod.Name = "FRAGARIA (STRAWBERRY)  ‘GASANA’"
            prod.Class = "Fruit Crops"
            prod.Light = "Sun"
            prod.PlugSize = "72"
            prod.PriceEach = "$.68"
         
            InsertProduct(c,*prod,w)
        
        prod = new(Product)
            prod.Height = "6 - 8 inches"
            prod.HardinessZone = "4 - 9"
            prod.BloomTime = "May - June"
            prod.ImageUrl = "http://www.jamesgreenhouses.com/_code/_dyn_images/products/Strawberry-Gasana.JPG"
            prod.Description = "Everbearing. Great for containers! Or in the ground."
            prod.Name = "FRAGARIA (STRAWBERRY)  ‘GASANA’"
            prod.Class = "Fruit Crops"
            prod.Light = "Sun"
            prod.PlugSize = "72"
            prod.PriceEach = "$.68"
        
            InsertProduct(c,*prod,w)
    }
    
    
}

 
 
 

 
