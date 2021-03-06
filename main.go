package main

import (
	"fmt"
	"log"

	_ "github.com/lib/pq"

	"github.com/jinzhu/gorm"
)

type UserModel struct {
	gorm.Model
	Name    string
	Address string
}

func main() {
	psqlInfo := "xxx"
	db, err := gorm.Open("postgres", psqlInfo)
	//defer db.Close()
	if err != nil {
		fmt.Println("failed")
	}
	log.Println("succcess")

	//db.AutoMigrate(&UserModel{})

	user2 := &UserModel{Name: "kidzz", Address: "Ranchizz"}

	// db.Find(user2)
	// user2.Address = "Chennai"
	// db.Save(user2)

	// db.Model(&user2).Updates(
	// 	map[string]interface{}{
	// 		"Name":    "Nidzz",
	// 		"Address": "Ranchi-kanke",
	// 	})

	//db.Table("user_models").Where("address = ?", "New York").Update("name", "Walker")

	//https://mindbowser.com/golang-go-with-gorm/  -- CRUD

	tx := db.Begin()
	err = tx.Create(&user2).Error
	if err != nil {
		tx.Rollback()
	}
	tx.Commit()

	//mindbowser.com/golang-go-with-gorm/  -- CRUD

	// -- one to one
	// -- ome to many
	// -- many to many

	db.Close()
}
