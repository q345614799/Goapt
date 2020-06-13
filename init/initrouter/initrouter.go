package initrouter

import (
	"apt/router"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/zap"
	"go.uber.org/zap"
	"time"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	_ "apt/docs"


)

func InitRouter() *gin.Engine {
	r:= gin.New()
	logger, _ := zap.NewProduction()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))


	r.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	Group := r.Group("")
	router.InitLhbbRouter(Group)


	return r
}
