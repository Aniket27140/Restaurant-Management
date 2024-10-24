package routes

import (
	controller "RestaurantMangement/controllers"

	"github.com/gin-gonic/gin"
)

func FoodRoutes(incomingRoutes *gin.Engine) {
	incomingRoutes.GET("/foods",controller.GetFood())
	incomingRoutes.GET("/foods/:food_id",controller.GetFood())
	incomingRoutes.POST("/foods",controller.CreateFood())
	incomingRoutes.PATCH("/foods/:food_id",controller.UpdateFood())
}
