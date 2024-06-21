package services

import (
	"gin_crud_user_management_panel/internal/models"
	"gin_crud_user_management_panel/internal/repositories"
)

func GetAllUsers() ([]models.User, error) {
	return repositories.GetAllUsers()
}

func UpdateUserRole(userID uint, roleID uint) error {
	return repositories.UpdateUserRole(userID, roleID)
}
