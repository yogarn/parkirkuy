package repository

import (
	"github.com/yogarn/parkirkuy/entity"
	"github.com/yogarn/parkirkuy/pkg/response"
	"gorm.io/gorm"
)

type IUserRepository interface {
	CreateUser(user *entity.User) (err error)
	GetUserByUsername(username string) (user *entity.User, err error)
	GetUserByEmail(email string) (user *entity.User, err error)
}

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository(db *gorm.DB) IUserRepository {
	return &UserRepository{
		db: db,
	}
}

func (r *UserRepository) CreateUser(user *entity.User) (err error) {
	err = r.db.Create(user).Error
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) GetUserByUsername(username string) (user *entity.User, err error) {
	user = new(entity.User)
	err = r.db.Where("username = ?", username).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &response.UserNotFound
		}
	}

	return user, nil
}

func (r *UserRepository) GetUserByEmail(email string) (user *entity.User, err error) {
	user = new(entity.User)
	err = r.db.Where("email = ?", email).First(user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil, &response.UserNotFound
		}
		return nil, err
	}

	return user, nil
}
