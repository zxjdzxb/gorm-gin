package util

import (
	"ES/model"
	"math/rand"

	"gorm.io/gorm"
)

func RandomString(n int) string {
	var letters = []byte("asdfghjklzxcvbnmqwertyuiop1234567890")
	result := make([]byte, n)
	// rand.Seed(time.Now().Unix())

	for i := range result {
		result[i] = letters[rand.Intn(len(letters))]
	}
	return string(result)
}

func IsTelephoneExist(db *gorm.DB, telephone string) bool {

	var user model.User
	db.Where("telephone = ?", telephone).First(&user)
	if user.ID != 0 {
		return true
	}
	return false
}
