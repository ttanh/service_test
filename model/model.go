package model

import "time"

type ResponseType struct {
	DateTime      time.Time  `json:"date_time"`
	TransactionId string `json:"transaction_id"`
	Code          int    `json:"code"`
	Description   string `json:"description"`
}

type Token struct {
	ID    int    `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	Value string `json:"value"`
}

func (t *Token) Insert() error {
	return db.Create(t).Error
}

func CheckToken(tokenValue string) bool {
	var token Token
	if err := db.Where("value = ?", tokenValue).First(&token).Error; err != nil {
		return false
	}
	return true
}