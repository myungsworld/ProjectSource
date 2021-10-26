package tradeTest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type oj struct {
	Id int `json:"id"`
}


// @tags trade
// @Summary 테스트
// @Description 테스트
// @Accept json
// @Produce json
// @Success 200 {string} HelloWorld
// @Router /api/trade/auth/test [get]
func Test(c *gin.Context) {
	c.JSON(http.StatusOK, "test")
}