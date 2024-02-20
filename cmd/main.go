package main

import (
	"backenddemo/pkg/routes"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

func main() {
	// ติดตั้ง mux สำหรับการใช้ใน router
	r := mux.NewRouter()
	// ตั้งค่า ROUTER ที่ฟังชั่น Getstartbackend ใน backenddemo/pkg/routes
	routes.Getstartbackend(r)
	http.Handle("/", r)
	// ตั้งค่า CORS และ PORT
	http.ListenAndServe(":8010",
		handlers.CORS(
			handlers.AllowedOrigins([]string{"http://localhost:3000"}),
			handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS", "PATCH"}),
			handlers.AllowedHeaders([]string{"Content-Type", "Authorization", "Accept", "AllowedOrigins", "AllowedHeaders", "AllowedMethods", "Content-Length", "Accept-Encoding", "X-Requested-With"}),
		)(r))

}

// วิธีที่สองสำหรับการทำ CORS
// http.ListenAndServe(":8010", CORSMiddleware(r))
// func CORSMiddleware(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		w.Header().Set("Access-Control-Allow-Origin", "*")
// 		w.Header().Set("Access-Control-Allow-Credentials", "true")
// 		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
// 		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, OPTIONS")

// 		next.ServeHTTP(w, r)
// 	})
// }
