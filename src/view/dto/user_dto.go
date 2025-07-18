package dto

type UserDTO struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	SessionToken string `json:"token"`
}
