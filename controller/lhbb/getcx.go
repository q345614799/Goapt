package lhbb

import (
	"apt/model/lhbbmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetCx(c *gin.Context)  {
	cx :=lhbbmodel.GetCxList()
		c.JSON(http.StatusOK,gin.H{
			"data":cx,
		})
}
