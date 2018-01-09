package models

import (
	"reflect"
	"testing"

	"github.com/jinzhu/gorm"
)

func TestUserAdmins(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
		want *gorm.DB
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := UserAdmins(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UserAdmins() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUserCreateDefault(t *testing.T) {
	type args struct {
		db *gorm.DB
	}
	tests := []struct {
		name string
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UserCreateDefault(tt.args.db)
		})
	}
}

func TestUser_String(t *testing.T) {
	tests := []struct {
		name string
		u    User
		want string
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.u.String(); got != tt.want {
				t.Errorf("User.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
