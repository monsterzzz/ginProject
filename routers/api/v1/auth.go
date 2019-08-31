package v1

import (
	"ginProject/models"
	"ginProject/pkg/e"
	"ginProject/pkg/util"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAuthToken(c *gin.Context){
	Username := c.Query("username")
	Password := c.Query("password")
	code := e.SUCCESS
	if Username == "" || Password == ""{
		code = e.INVALID_PARAMS
		c.JSON(http.StatusOK,gin.H{
			"code" : code,
			"msg" : e.GetMsg(code),
			"data" : nil,
		})
		return
	}

	var data = make(map[string]interface{})
	if models.CheckAuth(Username,Password){
		token,err := util.GenerateToken(Username,Password)
		if err != nil {
			code = e.ERROR_AUTH_TOKEN
		}else{
			data["token"] = token
			code = e.SUCCESS
		}
	}else{
		code = e.ERROR_AUTH
	}

	c.JSON(http.StatusOK,gin.H{
		"code":code,
		"msg":e.GetMsg(code),
		"data":data,
	})

}
