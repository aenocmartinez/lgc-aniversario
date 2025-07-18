package dto

type UserDTO struct {
	ID           int64  `json:"id"`
	Name         string `json:"name"`
	Username     string `json:"username"`
	Email        string `json:"email"`
	SessionToken string `json:"token"`
}
