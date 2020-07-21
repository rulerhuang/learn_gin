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
	var (
		code        = e.SUCCESS
		msg, errMsg = "", ""
	)
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	createdBy := c.Query("created_by")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()

	valid := validation.Validation{}
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.Required(content, "content").Message("内容不能为空")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 255, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只能为0或1")

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, value := range valid.Errors {
			log.Printf("tag=%s|err=%s\n", value.Key, value.Message)
			if errMsg == "" {
				errMsg = value.Message
			}
		}
	} else {
		exists, err := models.ExistTagById(tagId)
		if err != nil {
			code = e.ERROR
		} else if !exists {
			code = e.ERROR_NOT_EXIST_TAG
		} else {
			data := make(map[string]interface{})
			data["tag_id"] = tagId
			data["title"] = title
			data["desc"] = desc
			data["content"] = content
			data["created_by"] = createdBy
			data["state"] = state
			models.AddArticle(data)
		}
	}

	if errMsg != "" {
		msg = errMsg
	} else {
		msg = e.GetMsg(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make(map[string]interface{}),
	})
}

// 编辑文章
func EditArticle(c *gin.Context) {
	var (
		code        = e.SUCCESS
		msg, errMsg = "", ""
	)
	id := com.StrTo(c.Query("id")).MustInt()
	tagId := com.StrTo(c.Query("tag_id")).MustInt()
	title := c.Query("title")
	desc := c.Query("desc")
	content := c.Query("content")
	modifiedBy := c.Query("modified_by")

	valid := validation.Validation{}
	valid.Min(id, 1, "id").Message("文章ID必须大于0")
	valid.Min(tagId, 1, "tag_id").Message("标签ID必须大于0")
	valid.Required(title, "title").Message("标题不能为空")
	valid.MaxSize(title, 100, "title").Message("标题最长为100字符")
	valid.Required(desc, "desc").Message("简述不能为空")
	valid.MaxSize(desc, 255, "desc").Message("简述最长为255字符")
	valid.Required(content, "content").Message("内容不能为空")
	valid.MaxSize(content, 65535, "content").Message("内容最长为65535字符")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("内容最长为100字符")

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, value := range valid.Errors {
			log.Printf("tag=%s|err=%s\n", value.Key, value.Message)
			if errMsg == "" {
				errMsg = value.Message
			}
		}
	} else {
		exists, err := models.ExistArticleById(id)
		if err != nil {
			code = e.ERROR
		} else if !exists {
			code = e.ERROR_NOT_EXIST_ARTICLE
		} else {
			exists, err := models.ExistTagById(tagId)
			if err != nil {
				code = e.ERROR
			} else if !exists {
				code = e.ERROR_NOT_EXIST_TAG
			} else {
				data := make(map[string]interface{})
				data["tag_id"] = tagId
				data["title"] = title
				data["desc"] = desc
				data["content"] = content
				data["modified_by"] = modifiedBy
				models.EditArticle(id, data)
			}
		}
	}

	if errMsg != "" {
		msg = errMsg
	} else {
		msg = e.GetMsg(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make(map[string]interface{}),
	})
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	var (
		code = e.SUCCESS
		msg  string
	)
	id := com.StrTo(c.Param("id")).MustInt()
	if id < 1 {
		code = e.INVALID_PARAMS
		msg = "文章ID必须大于0"
		log.Printf("id=%d\n", id)

	}

	exists, err := models.ExistArticleById(id)
	if err != nil {
		code = e.ERROR
	} else if !exists {
		code = e.ERROR_NOT_EXIST_ARTICLE
	} else {
		models.DeleteArticle(id)
	}

	if msg == "" {
		msg = e.GetMsg(code)
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": make(map[string]interface{}),
	})
}
