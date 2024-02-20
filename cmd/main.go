package main

import (
	"backenddemo/pkg/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {

	r := mux.NewRouter()

	routes.Getstartbackend(r)
	http.Handle("/", r)
	http.ListenAndServe(":8010", CORSMiddleware(r))

	// http.ListenAndServe(":8010",
	// 	handlers.CORS(
	// 		handlers.AllowedOrigins([]string{"*"}),
	// 		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
	// 		handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type", "Authorization"}),
	// 	)(r))
}

func CORSMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*") // change this later
		w.Header().Set("Access-Control-Allow-Credentials", "true")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

		next.ServeHTTP(w, r)
	})
}
