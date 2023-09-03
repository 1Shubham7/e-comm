package handlers

import (
	"log"
	"net/http"
	"github.com/1shubham7/e-comm/data"
)

type Products struct {
	l *log.Logger
}

func NewProducts(l*log.Logger) *Products {
	return &Products{l}
}

// anything adds to the handler interface  needs to have a serve http method
func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
// we will now use encoding/json package to convert the struct we ahve into a json representative
	// listOfProducts := data.GetProducts()
	// now listOfProducts contains the list of products, but how do we return that to the user
	// we do that by converting listOfProducts into json. and this the the way

	// d, err := json.Marshal(listOfProducts)
	// instead of Marshal we are using this better alternative

	// err := listOfProducts.ToJSON(rw) //ToJSON is in data/products
	// if err != nil {
		// http.Error(rw, "unable to convert data to json", http.StatusInternalServerError)
	// }
	// rw.Write(d) instead of Marshal we are using this better alternative



	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) getProducts (rw http.ResponseWriter, r *http.Request) {
	listOfProducts := data.GetProducts()
	err := listOfProducts.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Unable to convert data to JSON", http.StatusInternalServerError)
	}
}
