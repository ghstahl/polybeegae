package productStore

import (
	"appengine"
	"appengine/datastore"
	"appengine/memcache"
	. "github.com/ahmetalpbalkan/go-linq"
	"models"
	"net/http"
	"reflect"
	"time"
)

const (
	secondsInADay  = 60 * 60 * 24        // from memcache server code
	secondsInAYear = secondsInADay * 365 // from memcache server code
	oneDay         = time.Duration(secondsInADay) * time.Second
	oneYear        = time.Duration(secondsInAYear) * time.Second
)

// productKey returns the key used for all product entries.
func productKey(c appengine.Context) *datastore.Key {
	// The string "default_products" here could be varied to have multiple product catalogs.
	return datastore.NewKey(c, "Product", "default_products", 0, nil)
}

func InsertProduct(c appengine.Context, prod models.Product) ( error) {

	// We set the same parent key on every Greeting entity to ensure each Greeting
	// is in the same entity group. Queries across the single entity group
	// will be consistent. However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Product", productKey(c))
	_, err := datastore.Put(c, key, &prod)
	 
    return err;
}

func datastoreGetAllClassesFromProds(c appengine.Context) []T {
	prods := GetAllProds(c)
	classes, err := From(prods).
		Select(func(prod T) (T, error) {
		return prod.(models.Product).Class, nil
	}).
		Distinct().
		OrderStrings().
		Results()

	c.Infof("Unable to retrieve list from Memcache: %v", reflect.TypeOf(classes))
	c.Infof("classes: %v", classes)
	c.Infof("err: %v", err)
	return classes
}

func GetAllClassesFromProds(c appengine.Context) []T {
	c.Infof("GetAllClassesFromProds: Enter")
	const cacheKey_product_classes = "product_classes"
	var item_list []T
	_, err := memcache.JSON.Get(c, cacheKey_product_classes, &item_list)
	if err != nil {
		//Memcache failed to extract the list.
		c.Infof("GetAllClassesFromProds: Unable to retrieve list from Memcache: %v", err)
		item_list = datastoreGetAllClassesFromProds(c)
		memcache_item := &memcache.Item{
			Key:        cacheKey_product_classes,
			Object:     item_list,
			Expiration: oneDay,
		}
		//Store into Memcache
		// Add the item to the memcache, if the key does not already exist
		if err := memcache.JSON.Set(c, memcache_item); err == memcache.ErrNotStored {
			c.Infof("item with key %q already exists", memcache_item.Key)
		} else if err != nil {
			c.Errorf("error adding item: %v", err)
		} else {

			c.Infof("got stored as json ")
		}
	}
	return item_list

}

func GetAllProds(c appengine.Context) []models.Product {

	const productsCacheKey = "products"

	var item_list []models.Product
	_, err := memcache.JSON.Get(c, productsCacheKey, &item_list)
	if err != nil {
		//Memcache failed to extract the list.
		c.Infof("Unable to retrieve list from Memcache: %v", err)

		q := datastore.NewQuery("Product").Ancestor(productKey(c)).Limit(1000)
		item_list = make([]models.Product, 0, 1000)
		if _, err := q.GetAll(c, &item_list); err != nil {
			return nil
		}

		c.Infof("got products")
		memcache_item := &memcache.Item{
			Key:        productsCacheKey,
			Object:     item_list,
			Expiration: oneDay,
		}
		//Store into Memcache
		// Add the item to the memcache, if the key does not already exist
		if err := memcache.JSON.Set(c, memcache_item); err == memcache.ErrNotStored {
			c.Infof("item with key %q already exists", memcache_item.Key)
		} else if err != nil {
			c.Errorf("error adding item: %v", err)
		} else {

			c.Infof("got stored as json ")
		}
	}
	return item_list
}

func EnsureData(w http.ResponseWriter, c appengine.Context) {
	q := datastore.NewQuery("Product").Ancestor(productKey(c)).Limit(1)
	products := make([]models.Product, 0, 1)
	if _, err := q.GetAll(c, &products); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	var prodExists bool = false

	// if you only need e:
	for _, e := range products {
		c.Infof("Product:%v", e)
		prodExists = true
		break
	}

	if !prodExists {
		prod := new(models.Product)
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

        err := InsertProduct(c, *prod)
        if(err != nil){
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

		prod = new(models.Product)
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

		err = InsertProduct(c, *prod)
        if(err != nil){
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

		prod = new(models.Product)
		prod.Height = "9 - 12 inches"
		prod.HardinessZone = "4 - 8"
		prod.BloomTime = "Early Spring, Summer, Fall"
		prod.ImageUrl = "http://www.jamesgreenhouses.com/_code/_dyn_images/products/Achillea-Little-Moonshine-(2).JPG"
		prod.Description = "This exciting new introduction from Blooms of Bressingham has the bright yellow flowers of 'Moonshine' but at half the height! Great for containers and in the front of the border."
		prod.Name = "ACHILLEA  ‘LITTLE MOONSHINE’"
		prod.Class = "Hardy Perennials"
		prod.Light = "full sun"
		prod.PlugSize = "72"
		prod.PriceEach = "$.65"
		prod.Royalty = "$.16"
		prod.Colors = "Vibrant Yellow"
		prod.PatentStatus = "USPPAF"
		err = InsertProduct(c, *prod)
        if(err != nil){
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }

		prod = new(models.Product)
		prod.Height = "12 inches"
		prod.HardinessZone = "4 - 9"
		prod.BloomTime = "May - July"
		prod.ImageUrl = "http://www.jamesgreenhouses.com/_code/_dyn_images/products/Achillea-New-Vinage-Red-(4).JPG"
		prod.Description = "This new series from Ball Seed offers a compact and tidy habit with deep red flowers all summer."
		prod.Name = "ACHILLEA  NEW VINTAGE™ SERIES ‘RED’"
		prod.Class = "Hardy Perennials"
		prod.Light = "full sun"
		prod.PlugSize = "72"
		prod.PriceEach = "$.62"
		prod.Royalty = "$.08"
		prod.Colors = "Deep Red"
		prod.PatentStatus = "USPPAF"
		err = InsertProduct(c, *prod)
        if(err != nil){
            http.Error(w, err.Error(), http.StatusInternalServerError)
        }
	}

}
