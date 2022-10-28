package Routes

import (
	"gin-restful-api-mysql/Controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	grp1 := r.Group("/api")
	{
		grp1.GET("users", Controllers.GetUsers)
		grp1.POST("user", Controllers.CreateUser)
		grp1.GET("user/:id", Controllers.GetById)
		grp1.PUT("user/:id", Controllers.Update)
		grp1.DELETE("user/:id", Controllers.Delete)
	}

	return r
}
