package model

type LoginResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type RegisterResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type ValidateResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type HealthResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}

type ErrorResponse struct {
	Ok      bool   `json:"ok"`
	Message string `json:"message"`
}
