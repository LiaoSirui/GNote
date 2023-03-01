package app_gnote

import (
	"github.com/LiaoSirui/GNote/internal/app/gnote/middleware"
	"github.com/LiaoSirui/GNote/internal/app/gnote/router"
	"github.com/gin-gonic/gin"
)

func StartServer() error {
	engine := gin.Default()
	engine.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"msg": "请求成功",
		})
	})
	// cors for debug
	engine.Use(middleware.Cors())
	router.SetRouter(engine)
	_ = engine.Run(":9090")
	return nil
}
