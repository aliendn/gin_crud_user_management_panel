package repositories

import (
	"gin_crud_user_management_panel/internal/models"
	"gin_crud_user_management_panel/pkg/database"
)

func CreateUser(user *models.User) error {
	db := database.GetDB()
	return db.Create(user).Error
}

func GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	db := database.GetDB()
	err := db.Where("username = ?", username).First(&user).Error
	return &user, err
}

func GetAllUsers() ([]models.User, error) {
	var users []models.User
	db := database.GetDB()
	err := db.Find(&users).Error
	return users, err
}

func UpdateUserRole(userID uint, roleID uint) error {
	db := database.GetDB()
	return db.Model(&models.User{}).Where("id = ?", userID).Update("role_id", roleID).Error
}
