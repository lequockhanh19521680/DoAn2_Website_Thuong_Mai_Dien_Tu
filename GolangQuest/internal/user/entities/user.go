package entities

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Username *string   `gorm:"column:username"`
	Password *string   `gorm:"column:password"`
	Salt     *string   `gorm:"column:salt"`
	Name     *string   `gorm:"column:name"`
	Birthday *string   `gorm:"column:birthday"`
	Gender   *bool     `gorm:"column:gender"`
	Email    *string   `gorm:"column:email"`
	Phone    *string   `gorm:"column:phone"`
	Type     *UserType `gorm:"column:type"`
	Avatar   *string   `gorm:"column:avatar"`
}

func (User) WithFields() []string {
	return []string{"username", "name", "gender", "email", "phone", "type"}
}
func (User) SearchFields() []string {
	return []string{"birthday"}
}
func (User) SortFields() []string {
	return []string{"username", "name", "gender", "email", "phone", "type"}
}
func (User) TableName() string {
	return "User"
}

func (u User) HashPassword(plainPassword string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(plainPassword), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(hash), nil
}

func (u User) ComparePassword(hashedPwd string, plainPwd string) bool {
	byteHash := []byte(hashedPwd)
	err := bcrypt.CompareHashAndPassword(byteHash, []byte(plainPwd))
	if err != nil {
		return false
	}

	return true
}

type UserType string

const (
	Seller UserType = "SELLER"
	Buyer  UserType = "BUYER"
	Admin  UserType = "ADMIN"
	Banned UserType = "BANNED"
)
