package model

import "time"

const (
	StatusInsert = 1
	StatusDone   = 2
	StatusReply  = 3
)

type Transaction struct {
	ID            int       `json:"id" gorm:"primary_key;AUTO_INCREMENT"`
	CustomerId    string    `json:"customer_id" gorm:"type:varchar(9);not null" valid:"required, length(0|9)"`
	CustomerName  string    `json:"customer_name" gorm:"type:varchar(200);not null" valid:"required,length(0|200)"`
	Token         string    `json:"token" gorm:"type:varchar(20);not null" valid:"required,length(0|20)"`
	DateTime      time.Time `json:"date_time" gorm:"not null" valid:"IsTime" valid:"required,rfc3339"`
	TransactionId string    `json:"transaction_id" gorm:"type:varchar(128);not null;unique" valid:"required,length(0|128)"`
	Status        int       `json:"status" gorm:"not null"`
}

func (t *Transaction) Insert() (err error) {
	return db.Create(t).Error
}

func (t *Transaction) UpdateStatusDone() (err error) {
	t.DateTime = time.Now()
	if err = db.Model(t).Updates(map[string]interface{}{"status": StatusDone, "date_time": t.DateTime.UTC()}).Error; err != nil {
		return err
	}
	return nil
}

func (t *Transaction) CheckDoneAndUpdateReply() error {
	err := db.Where("transaction_id = ? AND status = ?", t.TransactionId, StatusDone).First(&t).Error
	if err != nil {
		return err
	}
	t.DateTime = time.Now()
	if err = db.Model(t).Updates(map[string]interface{}{"status": StatusReply, "date_time": t.DateTime.UTC()}).Error; err != nil {
		return err
	}
	return nil
}
