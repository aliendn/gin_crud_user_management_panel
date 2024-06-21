package services

import (
	"gin_crud_user_management_panel/internal/models"
	"gin_crud_user_management_panel/internal/repositories"
)

func CreateUser(user *models.User) error {
	return repositories.CreateUser(user)
}

func GetUserByUsername(username string) (*models.User, error) {
	return repositories.GetUserByUsername(username)
}
