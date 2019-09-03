package routers

import (
	_ "ginProject/docs"
	"ginProject/middleware/JWT"
	"ginProject/routers/api/v1"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

func InitRouters() *gin.Engine {
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/auth", v1.GetAuthToken)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	//r.GET("/swagger/index", ginSwagger.WrapHandler(swaggerFiles.Handler))
	apiV1 := r.Group("/api/v1")
	apiV1.Use(JWT.Jwt())
	{
		apiV1.GET("/ss", v1.GetTags)
		apiV1.GET("/tags", v1.GetTags)
		apiV1.POST("/tags", v1.AddTag)
		apiV1.PUT("/tags/:id", v1.EditTag)
		apiV1.DELETE("/tags/:id", v1.DeleteTag)
	}
	{
		apiV1.GET("/articles", v1.GetArticles)
		apiV1.GET("/articles/:id", v1.GetOneArticle)
		apiV1.POST("/articles", v1.AddArticle)
		apiV1.PUT("/articles/:id", v1.UpdateArticle)
		apiV1.DELETE("/articles/:id", v1.DeleteArticle)
	}

	return r
}
