package routes

import (
	"resturant/controllers"

	"github.com/gin-gonic/gin"
)

func InvoiceRoutes(incomingRoutes *gin.Engine) {

	incomingRoutes.GET("/invoice", controllers.GetInvoices())
	incomingRoutes.GET("/invoice/:invoice_id", controllers.GetInvoice())
	incomingRoutes.POST("/invoice", controllers.CreateInvoice())
	incomingRoutes.PATCH("/invoice:invoice_id", controllers.UpdateInvoices())

}
