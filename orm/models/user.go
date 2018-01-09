package models

import (
	"fmt"
	"strconv"
	"time"

	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	Name  string `gorm:"unique_index:idx_users_name"`
	Admin bool   `gorm:"default:'false'"`
}

func UserAdmins(db *gorm.DB) *gorm.DB {
	return db.Where("admin = ?", true)
}

func UserCreateDefault(db *gorm.DB) {
	defaultUsers := []User{
		User{Name: "Peter", Admin: true},
		User{Name: "Mike", Admin: false},
		User{Name: "John", Admin: false},
	}

	for _, u := range defaultUsers {
		if err := db.Create(&u).Error; err != nil {
			fmt.Println("Catched error: ", err)
		}
	}
}

func (u User) String() string {
	return fmt.Sprintf("%d, %v, %v, %v, %s, %v",
		u.ID,
		u.CreatedAt.Format(time.UnixDate),
		u.UpdatedAt.Format(time.UnixDate),
		u.DeletedAt,
		u.Name,
		strconv.FormatBool(u.Admin),
	)
}
