package routes

import (
	"loja-go/controller"
	"net/http"
)

func CarregaRotas() {
	http.HandleFunc("/", controller.Index)
	http.HandleFunc("/create", controller.Create)
	http.HandleFunc("/store", controller.Store)
	http.HandleFunc("/delete", controller.Delete)
	http.HandleFunc("/edit", controller.Edit)
	http.HandleFunc("/update", controller.Update)
}
