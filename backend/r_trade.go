package main

import (
	
	tradeTest "ProjectSource/backend/handlers-trade/test"
	"github.com/gin-gonic/gin"

)

func SetWebTrade(rTrade *gin.RouterGroup) {





	// api/trade/auth
	rAuth := rTrade.Group("/auth")
	{
		//rAuth.GET("/sign-in")
		rAuth.GET("/test", tradeTest.Test)
	}
}
