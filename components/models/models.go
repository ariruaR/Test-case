package models

import (
	"fmt"
	"log"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	configReader "exampler/components/config"
)

var config = configReader.ReadConfig()

type Subcription struct {
	gorm.Model
	ServiceName string    `json:"service_name"`
	Price       int       `json:"price"`
	UserID      string    `json:"user_id"`
	StartDate   time.Time `json:"start_date"`
}

func CreateNewSub(
	service_name string,
	price int,
	userID string,
	start_date time.Time,
	db gorm.DB,
) error {
	sub := Subcription{
		ServiceName: service_name,
		Price:       price,
		UserID:      userID,
		StartDate:   start_date,
	}
	result := db.Create(&sub)
	return result.Error
}

func GetSubs(
	db gorm.DB,
) ([]Subcription, error) {
	var subs []Subcription
	result := db.Find(&subs)
	return subs, result.Error
}

func GetSubsById(
	db gorm.DB,
	userID int,
) ([]Subcription, error) {
	var subs []Subcription
	result := db.First(&subs, "UserID = ?", userID)
	return subs, result.Error
}

func UpdateSubs(
	db gorm.DB,
	userID int,
) error {
	var sub Subcription
	db.Model(&sub).Update()
}

func CreateDB() gorm.DB {
	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		config.Host, config.Port, config.User, config.Password, config.DBname)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Println(err)
	}

	db.AutoMigrate(&Subcription{})
	return *db
}
