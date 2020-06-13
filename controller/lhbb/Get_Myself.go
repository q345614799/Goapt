package lhbb

import (
	"apt/model/lhbbmodel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Get_Myself(c *gin.Context)  {
	SID := c.Query("SID")
	LHBB_myself := lhbbmodel.GetLhbbBySID(SID)
	c.JSON(http.StatusOK,gin.H{
		"data":LHBB_myself,
	})
}
