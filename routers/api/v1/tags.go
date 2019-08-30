package v1

import (
	"ginProject/models"
	"ginProject/pkg/e"
	"ginProject/pkg/setting"
	"ginProject/pkg/util"
	"github.com/Unknwon/com"
	"github.com/astaxie/beego/validation"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/common/log"
	"net/http"
)

func GetTags(c *gin.Context){
	name := c.Query("name")
	maps := make(map[string]interface{})
	data := make(map[string]interface{})

	if name != ""{
		maps["name"] = name
	}

	var state = -1
	if arg := c.Query("state"); arg != ""{
		state = com.StrTo(arg).MustInt()
		maps["state"] = state
	}

	code := e.SUCCESS

	data["lists"] = models.GetTags(util.GetPage(c),setting.PageSize,maps)
	data["total"] = models.GetTagTotal(maps)


	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : data,
	})
}

func AddTag(c *gin.Context)  {
	name := c.Query("name")
	state := com.StrTo(c.DefaultQuery("state","0")).MustInt()
	createdBy := c.Query("created_by")

	valid := validation.Validation{}
	valid.Required(name,"name").Message("can not be nil")
	valid.MaxSize(name,100,"name").Message("can not more than 100")
	valid.Required(createdBy, "created_by").Message("创建人不能为空")
	valid.MaxSize(createdBy, 100, "created_by").Message("创建人最长为100字符")
	valid.Range(state, 0, 1, "state").Message("状态只允许0或1")

	code := e.INVALID_PARAMS

	if !valid.HasErrors(){
		if !models.TagExists(name){
			code = e.SUCCESS
			models.AddTag(name,state,createdBy)
		}else {
			code = e.ERROR_EXIST_TAG
		}
	}else {
		for _,e := range valid.Errors{
			log.Info(e.Key,e.Message)
		}
	}

	c.JSON(http.StatusOK,gin.H{
		"code" : code,
		"msg" : e.GetMsg(code),
		"data" : make(map[string]string),
	})
	
}

func EditTag(c *gin.Context)  {
	
}

func DeleteTag(c *gin.Context){

}