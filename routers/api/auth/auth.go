package auth

import (
	"log"
	"net/http"

	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"learn.gin/models"
	"learn.gin/pkg/e"
	"learn.gin/pkg/util"
)

func GetAuth(ctx *gin.Context) {
	var (
		code     = e.SUCCESS
		valid    = validation.Validation{}
		userName = ctx.Query("username")
		passWord = ctx.Query("password")
		data     = make(map[string]interface{})
	)

	valid.Required(userName, "username").Message("用户名不能为空")
	valid.MaxSize(userName, 32, "username").Message("用户名不能超过32个字符")
	valid.Required(passWord, "password").Message("密码不能为空")
	valid.MaxSize(passWord, 32, "password").Message("密码不能超过32个字符")
	if valid.HasErrors() {
		code = e.INVALID_PARAMS
		for _, item := range valid.Errors {
			log.Printf("tag=%s|err=%s\n", item.Key, item.Message)
		}
	}

	if code == e.SUCCESS {
		ok := models.CheckAuth(userName, passWord)
		if ok {
			token, err := util.GenerateToken(userName, passWord)
			if err == nil {
				data["token"] = token
			} else {
				code = e.ERROR_AUTH_TOKEN
				log.Printf("call GenerateToken() faild, err=%v\n", err)
			}
		} else {
			code = e.ERROR_AUTH
		}
	}
	ctx.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  e.GetMsg(code),
		"data": data,
	})
}
