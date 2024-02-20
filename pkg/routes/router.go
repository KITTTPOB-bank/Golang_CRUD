package routes

import (
	"backenddemo/pkg/controllers"

	"github.com/gorilla/mux"
)

var Getstartbackend = func(router *mux.Router) {
	router.HandleFunc("/go", controllers.Quryparamblue).Methods("GET")
	router.HandleFunc("/test", controllers.Gotestget).Methods("GET")
	router.HandleFunc("/test", controllers.Gotestcrate).Methods("POST")
	router.HandleFunc("/go", controllers.GocreateBlueArchive).Methods("POST")
	router.HandleFunc("/upload", controllers.UploadFile).Methods("POST")

}
