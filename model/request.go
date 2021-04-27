package model

type LoginRequest struct {
	BankNumber string `json:"bank_number"`
	AccountId  string `json:"account_id"`
	Password   string `json:"password"`
}

type RegisterRequest struct {
	BankNumber string `json:"bank_number"`
	AccountId  string `json:"account_id"`
	Password   string `json:"password"`
}
