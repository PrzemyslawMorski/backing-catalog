package service

import (
	"github.com/hudl/fargo"
	"strconv"

	"github.com/codegangsta/negroni"
	"github.com/gorilla/mux"
	"github.com/unrolled/render"
)

// new server from
func NewServerFromApplication(app *fargo.Application) *negroni.Negroni {
	webClient := fulfillmentWebClient{
		rootURL: app.Instances[0].SecureVipAddress + ":" + strconv.Itoa(app.Instances[0].SecurePort) + "/skus",
	}

	return NewServerFromClient(webClient)
}

// NewServerFromClient configures and returns a Server.
func NewServerFromClient(webClient fulfillmentClient) *negroni.Negroni {
	formatter := render.New(render.Options{
		IndentJSON: true,
	})

	n := negroni.Classic()
	mx := mux.NewRouter()

	initRoutes(mx, formatter, webClient)

	n.UseHandler(mx)
	return n
}

func initRoutes(mx *mux.Router, formatter *render.Render, webClient fulfillmentClient) {
	mx.HandleFunc("/", rootHandler(formatter)).Methods("GET")
	mx.HandleFunc("/catalog", getAllCatalogItemsHandler(formatter)).Methods("GET")
	mx.HandleFunc("/catalog/{sku}", getCatalogItemDetailsHandler(formatter, webClient)).Methods("GET")
}
