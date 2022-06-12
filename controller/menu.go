package controller

import (
	"net/http"
	"server/model"

	"github.com/gin-gonic/gin"
)

func HandeListOfFoodItemInMenu(ctx *gin.Context) {
	menu := []model.FoodItem{
		{Name: "Dosa", Price: 50},
		{Name: "Fried Rice", Price: 65},
		{Name: "Roti", Price: 40},
		{Name: "Panner Butter Masala", Price: 150},
		{Name: "Veg Biryani", Price: 90},
		{Name: "Coffee", Price: 10},
		{Name: "Mango Juice", Price: 60}}
	ctx.JSON(http.StatusOK, menu)
}
