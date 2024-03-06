package repository

import (
	"github.com/arvan-challenge/src/entity"
	"gorm.io/gorm"
)

type usersRepository struct {
	CommonBehaviourRepository[entity.User]
}

func NewUsersRepository(db *gorm.DB) UserRepository {
	return &usersRepository{
		CommonBehaviourRepository: NewCommonBehaviour[entity.User](db),
	}
}
