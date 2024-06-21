package services

import (
	"gin_crud_user_management_panel/internal/models"
	"gin_crud_user_management_panel/internal/repositories"
)

func CreateStaff(staff *models.User) error {
	return repositories.CreateUser(staff)
}
