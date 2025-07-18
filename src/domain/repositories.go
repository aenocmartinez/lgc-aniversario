package domain

type UserRepository interface {
	FindByID(id int64) (*User, error)
	FindByEmail(email string) (*User, error)
	Save(user *User) error
	Update(user *User) error
	Delete(id int64) error
}
