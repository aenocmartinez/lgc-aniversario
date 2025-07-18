package domain

import "lgc/src/view/dto"

type User struct {
	id           int64
	name         string
	email        string
	password     string
	sessionToken string
	repository   UserRepository
}

func NewUser(repository UserRepository) *User {
	return &User{repository: repository}
}

func (u *User) SetID(id int64) {
	u.id = id
}

func (u *User) SetName(name string) {
	u.name = name
}

func (u *User) SetEmail(email string) {
	u.email = email
}

func (u *User) SetPassword(password string) {
	u.password = password
}

func (u *User) SetSessionToken(sessionToken string) {
	u.sessionToken = sessionToken
}

func (u *User) GetID() int64 {
	return u.id
}

func (u *User) GetName() string {
	return u.name
}

func (u *User) GetEmail() string {
	return u.email
}

func (u *User) GetPassword() string {
	return u.password
}

func (u *User) GetSessionToken() string {
	return u.sessionToken
}

func (u *User) Exists() bool {
	return u.id > 0
}

func (u *User) Save() error {
	return u.repository.Save(u)
}

func (u *User) Update() error {
	return u.repository.Update(u)
}

func (u *User) Delete() error {
	return u.repository.Delete(u.id)
}

func (u *User) FindByID(id int64) (*User, error) {
	return u.repository.FindByID(id)
}

func (u *User) FindByEmail(email string) (*User, error) {
	return u.repository.FindByEmail(email)
}

func (u *User) ToDTO() *dto.UserDTO {
	return &dto.UserDTO{
		ID:           u.id,
		Name:         u.name,
		Email:        u.email,
		SessionToken: u.sessionToken,
	}
}
