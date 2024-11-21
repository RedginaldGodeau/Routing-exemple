package route

import (
	"api/ent"
	"api/internal/controller"

	"github.com/labstack/echo/v4"
)

func Route(db *ent.Client, r *echo.Echo) {
	userController := controller.NewUserController(db)

	api := r.Group("/api")
	{
		user := api.Group("/user")
		{
			user.GET("/horses/:id/", userController.GetHorsesByUserID)
			user.GET("/user/", userController.GetUserByUsername)
			user.POST("/user/", userController.CreateUser)
		}
	}

}
