package user

import (
	"github.com/hanzohasashi17/farm/internal/user/model"
	"github.com/jmoiron/sqlx"
)

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db}
}

func (r *UserRepo) GetAllUsers() ([]model.User, error) {
	var users []model.User

	query := "select * from users"
	err := r.db.Select(&users, query)

	return users, err
}
