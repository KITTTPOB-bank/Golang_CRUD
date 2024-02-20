package routes

import (
	"backenddemo/pkg/controllers"

	"github.com/gorilla/mux"
)

// ตั้งค่า mux router
var Getstartbackend = func(router *mux.Router) {
	// ดึงข้อมูลของ id ที่เลือกใน database
	router.HandleFunc("/getallCharacter", controllers.GetallCharacter).Methods("GET")
	// ดึงข้อมูลทั้งหมดใน database
	router.HandleFunc("/getcharacterbyID", controllers.GetcharacterbyID).Methods("GET")
	// // สร้างข้อมูลในดาต้า database
	router.HandleFunc("/createCharacter", controllers.CreateCharacter).Methods("POST")
	// // แก้ไขข้อมูลใน database
	router.HandleFunc("/updateCharacter", controllers.UpdateCharacter).Methods("PUT")
	// // ลบข้อมูลในดาต้าเบส โดยใช้ id
	router.HandleFunc("/deleteCharacterbyID", controllers.DeleteCharacterbyID).Methods("DELETE")
	// // อับโหดลไฟล์
	router.HandleFunc("/uploadCharacterFile", controllers.UploadCharacterFile).Methods("POST")

}
