package routes

import (
	"gin_crud_user_management_panel/internal/controllers"
	"gin_crud_user_management_panel/internal/middlewares"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/register", controllers.Register)
	r.POST("/login", controllers.Login)

	auth := r.Group("/auth")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.GET("/profile", controllers.Profile)
		auth.POST("/logout", controllers.Logout)
	}

	admin := r.Group("/admin")
	admin.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
	{
		admin.POST("/create-staff", controllers.CreateStaff)
		admin.POST("/create-user", controllers.CreateUser)
		admin.GET("/users", controllers.ListUsers)
		admin.POST("/grant-permission", controllers.GrantPermission)
	}
}
