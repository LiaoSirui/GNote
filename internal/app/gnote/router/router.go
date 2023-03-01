package router

import (
	"github.com/LiaoSirui/GNote/internal/app/gnote/controller/graphql"

	"github.com/gin-gonic/gin"
)

func SetRouter(router *gin.Engine) error {

	router.POST("/graphql", graphql.GraphqlHandler())
	router.GET("/graphql", graphql.GraphqlHandler()) // graphql web

	return nil
}
