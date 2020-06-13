package router

import (
	"apt/controller/lhbb"
	"github.com/gin-gonic/gin"
)

func InitLhbbRouter(Router *gin.RouterGroup)  {
	LhbbRouter :=Router.Group("/apt/lhbb")
	{
		LhbbRouter.POST("/add",lhbb.Add)
		LhbbRouter.GET("/cx/find/alllist",lhbb.GetCx)
		LhbbRouter.GET("/find/sj/selectMyself",lhbb.Get_Myself)
	}
}

