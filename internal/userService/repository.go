package userService

import "gorm.io/gorm"

type UserRepository interface {
	CreateUser(user User) (User, error)
	GetAllUsers() ([]User, error)
	GetUserByID(id uint) (User, error)
	UpdateUserByID(id uint, updated User) (User, error)
	DeleteUserByID(id uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

func (u *userRepository) CreateUser(user User) (User, error) {
	err := u.db.Create(&user).Error
	return user, err
}

func (u *userRepository) GetAllUsers() ([]User, error) {
	var users []User
	err := u.db.Find(&users).Error
	return users, err
}

func (u *userRepository) GetUserByID(id uint) (User, error) {
	var user User
	err := u.db.First(&user, id).Error
	return user, err
}

func (u *userRepository) UpdateUserByID(id uint, updated User) (User, error) {
	var existing User
	err := u.db.First(&existing, id).Error
	if err != nil {
		return User{}, err
	}
	existing.Email = updated.Email
	existing.Password = updated.Password
	err = u.db.Save(&existing).Error
	return existing, err
}

func (u *userRepository) DeleteUserByID(id uint) error {
	var existing User
	err := u.db.First(&existing, id).Error
	if err != nil {
		return err
	}
	return u.db.Unscoped().Delete(&existing).Error

}
