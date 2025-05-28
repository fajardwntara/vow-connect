package user

import (
	"context"
	"errors"

	"github.com/fajardwntara/vow-connect/helpers"
	"gorm.io/gorm"
)

type userRepository struct {
	db *gorm.DB
}

// Implements constructor in go
func NewUserRepository(db *gorm.DB) *userRepository {
	return &userRepository{db}
}

// Create implements UserRepository.
func (u *userRepository) Create(ctx context.Context, user *User) error {

	// hash the password
	user.Password = helpers.HashPassword(user.Password)

	err := u.db.WithContext(ctx).Create(user).Error

	if err != nil {
		return err
	}

	return nil
}

// Delete implements UserRepository.
func (u *userRepository) Delete(ctx context.Context, id uint) error {
	if err := u.db.WithContext(ctx).Delete(&User{}, id).Error; err != nil {
		return err
	}

	return nil
}

// GetAll implements UserRepository.
func (u *userRepository) GetAll(ctx context.Context) ([]User, error) {
	var users []User

	result := u.db.WithContext(ctx).Find(&users)

	if result.Error != nil {
		return nil, result.Error
	}

	return users, nil
}

// GetByEmail implements UserRepository.
func (u *userRepository) GetByEmail(ctx context.Context, email string) (*User, error) {
	panic("unimplemented")
}

// GetByID implements UserRepository.
func (u *userRepository) GetByID(ctx context.Context, id uint) (*User, error) {
	var user User

	result := u.db.WithContext(ctx).Find(&user, id)

	if result.Error != nil {
		return nil, result.Error
	}

	if result.RowsAffected == 0 {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

// Update implements UserRepository.
func (u *userRepository) Update(ctx context.Context, user *User) error {
	panic("unimplemented")
}
