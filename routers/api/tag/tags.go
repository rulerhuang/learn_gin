package tag

import (
	"net/http"

	"learn.gin/models"
	"learn.gin/pkg/e"
	"learn.gin/pkg/setting"
	"learn.gin/pkg/util"

	"github.com/gin-gonic/gin"
	"github.com/unknwon/com"
)

// 获取多个文章的标签
func GetTags(c *gin.Context) {
	var (
		maps = make(map[string]interface{})
		data = make(map[string]interface{})
	)

	name := c.Query("name")
	if name != "" {
		maps["name"] = name
	}

	var state int = -1
	arg := c.Query("state")
	if arg != "" {
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	data["total"] = models.GetTagTotal(maps)
	data["lists"] = models.GetTags(util.GetOffset(c), setting.PageSize, maps)

	code := e.SUCCESS
	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}

// 新增文章的标签
func AddTag(c *gin.Context) {
}

// 编辑文章的标签
func EditTag(c *gin.Context) {
}

// 删除文章的标签
func DeleteTag(c *gin.Context) {
}
