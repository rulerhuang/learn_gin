package tag

import (
	"log"
	"net/http"

	"learn.gin/models"
	"learn.gin/pkg/e"
	"learn.gin/pkg/setting"
	"learn.gin/pkg/util"

	"github.com/astaxie/beego/validation"
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
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state", "0")).MustInt()
	createdBy := c.Query("created_by")
	code := e.SUCCESS

	valid := validation.Validation{}
	valid.Required(name, "name").Message("标签名称不能为空")
	valid.MaxSize(name, 100, "name").Message("标签名称不能超过100个字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许为0或者1")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人不能超过100字符")
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, item := range valid.Errors {
			log.Printf("tag=%s|err=%s\n", item.Key, item.Message)
		}
	}

	if code == e.SUCCESS {
		exist, err := models.ExistTagByName(name)
		if err != nil {
			code = e.ERROR
			log.Printf("ExistTagByName faild, err=%s\n", err)
		} else if exist {
			code = e.ERROR_EXIST_TAG
		} else {
			err = models.AddTag(name, state, createdBy)
			if err != nil {
				code = e.ERROR
				log.Printf("AddTag faild, err=%s\n", err)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 编辑文章的标签
func EditTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	name := c.Query("name")
	modifiedBy := c.Query("modified_by")
	state := com.StrTo(c.Query("state")).MustInt()
	code := e.SUCCESS

	valid := validation.Validation{}
	valid.Required(name, "name").Message("标签名称不能为空")
	valid.MaxSize(name, 100, "name").Message("标签名称不能超过100字符")
	valid.Required(id, "id").Message("ID不能为空")
	valid.Required(modifiedBy, "modified_by").Message("修改人不能为空")
	valid.MaxSize(modifiedBy, 100, "modified_by").Message("修改人不能超过100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许为0或1")

	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, item := range valid.Errors {
			log.Printf("tag=%s|err=%s\n", item.Key, item.Message)
		}
	}

	if code == e.SUCCESS {
		exist, err := models.ExistTagById(id)
		if err != nil {
			code = e.ERROR
			log.Printf("ExistTagById faild, err=%s\n", err)
		} else if !exist {
			code = e.ERROR_NOT_EXIST_TAG
		} else {
			data := make(map[string]interface{})
			data["name"] = name
			data["modified_by"] = modifiedBy
			data["state"] = state
			err = models.EditTag(id, data)
			if err != nil {
				code = e.ERROR
				log.Printf("EditTag faild, err=%s\n", err)
			}
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}

// 删除文章的标签
func DeleteTag(c *gin.Context) {
	id := com.StrTo(c.Param("id")).MustInt()
	code := e.SUCCESS

	if id < 0 {
		code = e.ERROR
		log.Printf("id = %v\n", c.Param("id"))
	}

	exist, err := models.ExistTagById(id)
	if err != nil {
		code = e.ERROR
		log.Printf("ExistTagById faild, err=%s\n", err)
	} else if !exist {
		code = e.ERROR_NOT_EXIST_TAG
	} else {
		err = models.DeleteTag(id)
		if err != nil {
			code = e.ERROR
			log.Printf("DeleteTag faild, err=%s\n", err)
		}
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": make(map[string]string),
	})
}
