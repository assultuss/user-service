package db

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var db *gorm.DB

func ConnectDB() error {
	user := os.Getenv("POSTGRES_USER")
	dbname := os.Getenv("POSTGRES_DB")
	password := os.Getenv("POSTGRES_PASSWORD")
	host := os.Getenv("POSTGRES_HOST")
	port := os.Getenv("POSTGRES_PORT")

	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s", user, password, dbname, host, port)

	conn, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	db = conn
	return nil
}

func AutoMigrate() error {
	err := db.AutoMigrate(&User{})
	if err != nil {
		return err
	}
	return nil
}

func CreateUser(user *User) error {
	err := db.Create(user).Error
	if err != nil {
		return err
	}
	return nil
}

func GetUsers() ([]User, error) {
	var users []User
	err := db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func FilterUsers(dateStart, dateEnd, ageStart, ageEnd int64) ([]User, error) {
	query := db.Model(&User{})

	if dateStart != 0 && dateEnd != 0 {
		query = query.Where("recording_date BETWEEN ? AND ?", dateStart, dateEnd)
	} else if dateStart != 0 {
		query = query.Where("recording_date >= ?", dateStart)
	} else if dateEnd != 0 {
		query = query.Where("recording_date <= ?", dateEnd)
	}

	if ageStart != 0 && ageEnd != 0 {
		query = query.Where("age BETWEEN ? AND ?", ageStart, ageEnd)
	} else if ageStart != 0 {
		query = query.Where("age >= ?", ageStart)
	} else if ageEnd != 0 {
		query = query.Where("age <= ?", ageEnd)
	}

	var users []User
	if err := query.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
