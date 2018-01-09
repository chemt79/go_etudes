package main

// Example of OOP in GoLang

import (
	"fmt"
	"os"

	"github.com/chemt79/go_etudes/orm/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func main() {
	db, err := gorm.Open("postgres", "host=localhost user=postgres dbname=postgres sslmode=disable password=")
	defer db.Close()

	if err != nil {
		fmt.Println("DB connection error", err)
		os.Exit(-1)
	}

	// db.LogMode(true)

	db.AutoMigrate(&models.User{})
	// models.UserCreateDefault(db)

	var users []models.User

	// All users (general query)
	db.Limit(100).Find(&users)

	userHeader := []string{
		"ID", "CreatedAt", "UpdatedAt", "DeletedAt", "Name", "Admin",
	}

	rows := make([]models.Stringer, len(users))
	for i, v := range users {
		rows[i] = v
	}
	fmt.Println("All users:")
	fmt.Println(models.RenderTable(&userHeader, &rows))

	// Only admins (scopes using)
	db.Limit(100).Scopes(models.UserAdmins).Find(&users)

	rows = make([]models.Stringer, len(users))
	for i, v := range users {
		rows[i] = v
	}
	fmt.Println("All admins:")
	fmt.Println(models.RenderTable(&userHeader, &rows))
}
