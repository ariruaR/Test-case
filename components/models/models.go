package models

import (
	"context"
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
	field string,
	value interface{},
) bool {
	var sub Subcription
	db.Model(&sub).Update(field, value)
	return true
}

func DeleteSubs(
	db gorm.DB,
	userID int,
) error {
	ctx := context.Background()
	_, err := gorm.G[Subcription](&db).Where("userID = ?", userID).Delete(ctx)
	if err != nil {
		return err
	}
	return nil

}

func GetTotalPrice(
	db gorm.DB,
	userID string,
	serviceName string,
	startDate, endDate time.Time,
) (int, error) {
	var total int64
	query := db.Model(&Subcription{}).Where("start_date >= ? AND start_date <= ?", startDate, endDate)
	if userID != "" {
		query = query.Where("user_id = ?", userID)
	}
	if serviceName != "" {
		query = query.Where("service_name = ?", serviceName)
	}
	err := query.Select("SUM(price)").Scan(&total).Error
	return int(total), err
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
