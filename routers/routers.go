package routers

import (
	"ginProject/routers/api/v1"
	"github.com/gin-gonic/gin"
)

func InitRouters() *gin.Engine{
	r := gin.New()

	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	apiV1 := r.Group("/api/v1")
	{
		apiV1.GET("/tags",v1.GetTags)
		apiV1.POST("/tags",v1.AddTag)
		apiV1.PUT("/tags/:id",v1.EditTag)
		apiV1.DELETE("/tags/:id",v1.DeleteTag)
	}



	return r
}
