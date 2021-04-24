package repositories

import (
	"github.com/komtriangle/Bodroe_ytro_GO/db"
	"github.com/komtriangle/Bodroe_ytro_GO/models"
)

type UserIRepository interface {
	Insert(user *models.User) (bool, error)
	GetAll() ([]models.User, error)
}

type UserRepository struct {
	db db.Database
}

func NewUserRepository(database db.Database) *UserRepository {
	return &UserRepository{database}
}

func (u UserRepository) Insert(user *models.User) (bool, error) {
	res, err := user.Validate()
	if !res {
		return false, err
	}
	res, err = u.db.CreateUser(user)
	if !res {
		return false, err
	}
	return true, nil
}

func (u UserRepository) GetAll() ([]models.User, error) {
	var users []models.User
	users, err := u.db.GetAllUsers()
	return users, err
}
