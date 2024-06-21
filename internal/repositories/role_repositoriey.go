package repositories

import (
	"gin_crud_user_management_panel/internal/models"
	"gin_crud_user_management_panel/pkg/database"
)

func GetRoleByName(name string) (*models.Role, error) {
	var role models.Role
	db := database.GetDB()
	err := db.Where("name = ?", name).First(&role).Error
	return &role, err
}
