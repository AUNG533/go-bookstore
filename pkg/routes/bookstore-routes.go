package routes

import (
	// "./controllers"
	"github.com/gorilla/mux"
	"github.com/akhil/go-bookstore/pkg/controllers"
	// "github.com/AUNG533/go-bookstore/pkg/controllers"
)

var RegisterBookStoreRoutes = func (router *mux.Router)  {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("PODT")
}