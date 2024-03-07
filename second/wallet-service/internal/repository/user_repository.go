package repository

import (
	"context"
	"github.com/arvan-challenge/wallet/internal/entity"
	"gorm.io/gorm"
)

type UserRepository interface {
	CreateUserWithFirstTransaction(ctx context.Context, user *entity.User, amount float64) error
	GetUserByField(ctx context.Context, field string, value interface{}) (*entity.User, error)
	GetUserByID(ctx context.Context, userID uint) (*entity.User, error)
	UpdateUser(ctx context.Context, user *entity.User) error
	DeleteUser(ctx context.Context, userID uint) error
}

type userRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{
		db: db,
	}
}

func (u *userRepository) CreateUserWithFirstTransaction(ctx context.Context, user *entity.User, amount float64) error {
	// Start a database transaction
	tx := u.db.WithContext(ctx).Begin()

	// Create the user
	if err := tx.Create(user).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Create the user's first transaction
	transaction := &entity.Transaction{
		UserID: user.ID,
		Amount: amount,
	}
	if err := tx.Create(transaction).Error; err != nil {
		tx.Rollback()
		return err
	}

	// Commit the transaction
	return tx.Commit().Error
}

// GetUserByID retrieves a user by ID.
func (u *userRepository) GetUserByID(ctx context.Context, userID uint) (*entity.User, error) {
	return u.GetUserByField(ctx, "id", userID)
}

// GetUserByField queries a record by a specific field.
func (u *userRepository) GetUserByField(ctx context.Context, field string, value interface{}) (*entity.User, error) {
	var user entity.User
	err := u.db.WithContext(ctx).Where(field+"=?", value).First(&user).Error
	return &user, err
}

// UpdateUser updates a user's information.
func (u *userRepository) UpdateUser(ctx context.Context, user *entity.User) error {
	return u.db.WithContext(ctx).Save(user).Error
}

// DeleteUser deletes a user by ID.
func (u *userRepository) DeleteUser(ctx context.Context, userID uint) error {
	return u.db.WithContext(ctx).Where("id = ?", userID).Delete(&entity.User{}).Error
}
