package router

import (
	"github.com/Service-Compute-Group/Course-Project/controller"
	"github.com/gin-gonic/gin"
)

func CreateServer() *gin.Engine {
	router := gin.Default()
	router.GET("/api", controller.GetApiList)
	router.GET("/api/user", controller.GetUserList)
	router.GET("api/user/:user/article", controller.GetArticleWithUser)
	router.GET("api/article", controller.GetAllArticles)
	router.GET("api/article/:articleID", controller.GetArticleWithArticleID)
	router.GET("api/tag", controller.GetAllTags)
	router.GET("api/tag/:tag", controller.GetArticlesWithTag)

	return router
}
