package controllers

import (
	"backenddemo/pkg/models"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"
)

type Datas []models.Dataget
type testdata struct {
	Getcharecter string `json:"getcharacter"`
	Good         string `json:"good"`
}

func Gotestget(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	data := models.GetAllBlueArchive()
	json.NewEncoder(w).Encode(data)
}

func GocreateBlueArchive(w http.ResponseWriter, r *http.Request) {
	var data models.BlueArchivess
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	res, nil := models.CreateBlueArchive(data)
	fmt.Print(nil)
	json.NewEncoder(w).Encode(res)

}

func Gotestcrate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	var getdata models.Dataget
	json.NewDecoder(r.Body).Decode(&getdata)
	testdatas := testdata{
		Getcharecter: getdata.Title + getdata.Desc,
		Good:         getdata.Title + getdata.Desc,
	}
	json.NewEncoder(w).Encode(testdatas)

}

func Quryparamblue(w http.ResponseWriter, r *http.Request) {

	Idget := r.URL.Query().Get("id")
	Id, _ := strconv.Atoi(Idget)
	// check := r.URL.Query().Get("check")

	data, _ := models.GetBlueById(Id)
	json.NewEncoder(w).Encode(data)

}

// res, _ := json.Marshal(id)

func UploadFile(w http.ResponseWriter, r *http.Request) {
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
