package controllers

import (
	"backenddemo/pkg/dbconfig"
	"backenddemo/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"gorm.io/gorm"
)

// ตั้งค่าและเรียกใช้ ฐานข้อมูล
var db *gorm.DB

func init() {
	dbconfig.Connect()
	db = dbconfig.GetDB()
	fmt.Println("Successfully connected to database!", db)
}

type Characters models.CharacterData

func GetallCharacter(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)

	var getalldata []Characters
	db.Table("Characters").Find(&getalldata)

	res, _ := json.Marshal(getalldata)
	fmt.Println(res)

	json.NewEncoder(w).Encode(getalldata)

}

func CreateCharacter(w http.ResponseWriter, r *http.Request) {
	db.AutoMigrate(&Characters{})
	w.WriteHeader(http.StatusOK)
	var getdatabyid = &Characters{}
	if err := json.NewDecoder(r.Body).Decode(&getdatabyid); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Newcharacter := Characters{
		Character: getdatabyid.Character,
	}

	if err := db.Table("Characters").Create(&Newcharacter).Error; err != nil {
		return
	}
	json.NewEncoder(w).Encode(&Newcharacter)
}

func GetcharacterbyID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	var getdatabyid Characters

	getid := r.URL.Query().Get("id")
	Id, _ := strconv.Atoi(getid)
	db.Table("Characters").Where("ID=?", Id).Find(&getdatabyid)
	json.NewEncoder(w).Encode(&getdatabyid)

}
func UpdateCharacter(w http.ResponseWriter, r *http.Request) {
	var updateData = &Characters{}
	var getdatabyid Characters
	if err := json.NewDecoder(r.Body).Decode(&updateData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	databyid := db.Table("Characters").Where("ID=?", updateData.Id).Find(&getdatabyid)
	if updateData.Character != "" {
		getdatabyid.Character = updateData.Character
	}

	databyid.Save(&getdatabyid)

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(&getdatabyid)
}

func DeleteCharacterbyID(w http.ResponseWriter, r *http.Request) {
	var data Characters
	var dropData = &Characters{}
	if err := json.NewDecoder(r.Body).Decode(&dropData); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db.Table("Characters").Where("ID=?", dropData.Id).Delete(data)
}

func UploadCharacterFile(w http.ResponseWriter, r *http.Request) {
	file, handler, err := r.FormFile("file")
	fileName := r.FormValue("file_name")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	f, err := os.OpenFile(handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	io.WriteString(w, "File "+fileName+" Uploaded successfully")
	io.Copy(f, file)
}
