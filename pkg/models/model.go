package models

import (
	"backenddemo/pkg/dbconfig"
	"fmt"

	"gorm.io/gorm"
)

var db *gorm.DB

func init() {
	dbconfig.Connect()
	db = dbconfig.GetDB()
	fmt.Println("Successfully connected to database!", db)
}

type bluearchive struct {
	Id        int    `gorm:"primaryKey"`
	Character string `gorm:"column:character"`
}

func GetAllBlueArchive() []bluearchive {
	// db.AutoMigrate(&bluearchive{})
	var bluearchive []bluearchive
	db.Table("bluearchive").Find(&bluearchive)
	fmt.Println(bluearchive)
	return bluearchive
}

func GetBlueById(Id int) (*bluearchive, *gorm.DB) {
	var getdata bluearchive
	db := db.Table("bluearchive").Where("ID=?", Id).Find(&getdata)
	return &getdata, db
}

func CreateBlueArchive(data BlueArchivess) (*bluearchive, error) {
	var maxID int
	err := db.Table("bluearchive").Select("max(Id)").Row().Scan(&maxID)
	if err != nil {
		return nil, err
	}

	maxID++
	newRecord := bluearchive{
		Id:        maxID,
		Character: data.Character,
	}

	if err := db.Table("bluearchive").Create(&newRecord).Error; err != nil {
		return nil, err
	}

	return &newRecord, nil
}

type Dataget struct {
	Title string `json:"Title"`
	Desc  string `json:"Desc"`
}

type BlueArchivess struct {
	Character string `json:"character"`
}
