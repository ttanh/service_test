package model

type TransactionInfo struct {
	CustomerId    string `json:"customer_id" valid:"required"`
	CustomerName  string `json:"customer_name" valid:"length(0|10)"`
	Token         string `json:"token"`
	DateTime      int64  `json:"date_time" valid:"IsTime"`
	TransactionId string `json:"transaction_id"`
}

type ResponseType struct {
	DateTime      int64  `json:"date_time"`
	TransactionId string `json:"transaction_id"`
	Code          int    `json:"code"`
	Description   string `json:"description"`
}

type Token struct {
	Id    int    `json:"id"`
	Value string `json:"value"`
}
