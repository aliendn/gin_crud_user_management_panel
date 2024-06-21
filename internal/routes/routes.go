package routes

import (
	"gin_crud_user_management_panel/internal/controllers"
	"gin_crud_user_management_panel/internal/middlewares"
	"net/http"

	_ "gin_crud_user_management_panel/cmd/gin_crud_user_management_panel/docs" // Necessary for documentation generation

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

// @BasePath /api/v1

// PingExample godoc
// @Summary ping example
// @Schemes
// @Description do ping
// @Tags example
// @Accept json
// @Produce json
// @Success 200 {string} Helloworld
// @Router /example/helloworld [get]
func Helloworld(g *gin.Context) {
	g.JSON(http.StatusOK, "helloworld")
}

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {

	v1 := r.Group("/api/v1")
	{
		eg := v1.Group("/example")
		{
			eg.GET("/helloworld", Helloworld)
		}
		v1.POST("/register", controllers.Register)
		v1.POST("/login", controllers.Login)
		auth := v1.Group("/auth")
		auth.Use(middlewares.AuthMiddleware())
		{
			auth.GET("/profile", controllers.Profile)
			auth.POST("/logout", controllers.Logout)
		}

		admin := v1.Group("/admin")
		admin.Use(middlewares.AuthMiddleware(), middlewares.AdminMiddleware())
		{
			admin.POST("/create-staff", controllers.CreateStaff)
			admin.POST("/create-user", controllers.CreateUser)
			admin.GET("/users", controllers.ListUsers)
			admin.POST("/grant-permission", controllers.GrantPermission)
		}
	}

}
