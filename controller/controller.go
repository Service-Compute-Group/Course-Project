package controller

import (
	"github.com/Service-Compute-Group/Course-Project/model"
	"github.com/gin-gonic/gin"
)

func GetApiList(c *gin.Context) {

	apiMaps := make(map[string]string)
	apiMaps[`/api`] = `使用GET获取可用API列表`
	apiMaps[`api/user`] = `使用GET获取用户列表`
	apiMaps[`/api/user/{user}/article`] = `使用GET获取该用户发表的所有文章列表`
	apiMaps[`/api/article`] = `使用GET获取全站所有文章列表`
	apiMaps[`/api/article/{articalID}}`] = `使用GET获取该文章的具体内容`
	apiMaps[`/api/tag`] = `使用GET获取所有tag列表`
	apiMaps[`/api/tag/{tagname}`] = `使用GET获取所有拥有此tag的文章列表`

	c.JSON(200, gin.H{
		"data": apiMaps,
	})
}

func GetUserList(c *gin.Context) {
	userNameList := model.GetUserNameList()
	c.JSON(200, gin.H{
		"data": userNameList,
	})
}

func GetArticleWithUser(c *gin.Context) {
	username := c.Param("user")
	articleList, userExist := model.GetArticleList(username)
	if !userExist {
		c.JSON(404, gin.H{
			"data": "user not exist",
		})
	} else {
		c.JSON(200, gin.H{
			"data": articleList,
		})
	}
}

func GetAllArticles(c *gin.Context) {
	articleList := model.GetAllArticles()
	c.JSON(200, gin.H{
		"data": articleList,
	})
}

func GetArticleWithArticleID(c *gin.Context) {
	articleID := c.Param("articleID")
	concreteArticle, articleExist := model.GetConcreteInfo(articleID)
	if !articleExist {
		c.JSON(404, gin.H{
			"data": "article not exist",
		})
	} else {
		c.JSON(200, gin.H{
			"data": concreteArticle,
		})
	}
}

func GetAllTags(c *gin.Context) {
	allTags := model.GetAllTags()
	c.JSON(200, gin.H{
		"data": allTags,
	})
}

func GetArticlesWithTag(c *gin.Context) {
	tag := c.Param("tag")
	articleList, tagExist := model.GetArticleListWithTag(tag)
	if !tagExist {
		c.JSON(404, gin.H{
			"data": "tag not exist",
		})
	} else {
		c.JSON(200, gin.H{
			"data": articleList,
		})
	}
}
