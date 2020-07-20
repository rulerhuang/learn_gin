package routers

import (
	"github.com/gin-gonic/gin"
	"learn.gin/pkg/setting"
	"learn.gin/routers/api/article"
	"learn.gin/routers/api/tag"
)

func helloHandler(c *gin.Context) {
	c.JSON(200, gin.H{"message": "how are you"})
}

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	r.GET("/", helloHandler)
	// 路由组
	apis := r.Group("/api/v1")

	// 文章标签
	apis.GET("/tags", tag.GetTags)
	apis.POST("/tags", tag.AddTag)
	apis.PUT("/tags/:id", tag.EditTag)
	apis.DELETE("/tags/:id", tag.DeleteTag)

	// 文章
	apis.GET("/article", article.GetArticle)
	apis.GET("/articles", article.GetArticles)
	apis.POST("/article", article.AddArticle)
	apis.PUT("/article/:id", article.EditArticle)
	apis.DELETE("/article/:id", article.DeleteArticle)
	return r
}
