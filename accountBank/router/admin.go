package router

import (
	"GoLearn/accountBank/controller"
	"github.com/gin-gonic/gin"
)

func setupAdminRoutes(c *controller.Controller, api *gin.RouterGroup)  {
	api.POST("/create-account", c.CreateAccount)
	//api.GET("/getAmount/:id", c.GetAmount)
	//api.PUT("/addFund/:amount", c.AddFund)
	//api.PUT("/withdrawMoney/:amount", c.WithDrawMoney)
}
