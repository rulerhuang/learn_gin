package article

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"

	"learn.gin/models"
	"learn.gin/pkg/e"
	"learn.gin/pkg/setting"
	"learn.gin/pkg/util"
)

// 获取单个文章
func GetArticle(c *gin.Context) {
	var (
		data interface{}
		code = e.SUCCESS
	)

	id := com.StrTo(c.Param("id")).MustInt()
	if id <= 0 {
		code = e.INVALID_PARAMS
		log.Printf("id = %v\n", c.Param("id"))
	}

	data = models.GetArticle(id)
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 获取多个文章
func GetArticles(c *gin.Context) {
	var (
		data  = make(map[string]interface{})
		maps  = make(map[string]interface{})
		valid = validation.Validation{}
		code  = e.SUCCESS
	)

	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	tagId := com.StrTo(c.DefaultQuery("tag_id", "0")).MustInt()
	valid.Range(state, 0, 1, "state").Message("state只允许为0或1")
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, item := range valid.Errors {
			log.Printf("tag=%s|err=%s\n", item.Key, item.Message)
		}
	} else {
		maps["state"] = state
		maps["tag_id"] = tagId
		data["lists"] = models.GetArticles(util.GetOffset(c), setting.PageSize, maps)
		data["total"] = models.GetArticleTotal(maps)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章
func AddArticle(c *gin.Context) {
}

// 编辑文章
func EditArticle(c *gin.Context) {
}

// 删除文章
func DeleteArticle(c *gin.Context) {
}
