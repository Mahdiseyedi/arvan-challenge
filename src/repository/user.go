package repository

import (
	"github.com/arvan-challenge/src/entity"
	repository "github.com/arvan-challenge/src/repository/interface"
	"gorm.io/gorm"
)

type usersRepository struct {
	repository.CommonBehaviourRepository[entity.User]
}

func NewUsersRepository(db *gorm.DB) repository.UserRepository {
	return &usersRepository{
		CommonBehaviourRepository: repository.NewCommonBehaviour[entity.User](db),
	}
}
