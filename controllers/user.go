package controllers

import (
	"gin/forms"
	"gin/libservers"
	"gin/response"
	"github.com/gin-gonic/gin"
	"github.com/gookit/validate"
	"log"
)
//用户控制器
type User struct {
}

func initUserRouteAr(r *gin.Engine) {
	user := new(User)
	AppGroup := r.Group("/Member/v1")
	AppGroup.GET("/list", user.list)
	AppGroup.GET("/show/:id", user.show)
	AppGroup.POST("/add", user.add)
	AppGroup.POST("/update", user.update)
	AppGroup.GET("/delete/:id", user.delete)
}
/*
@获取列表信息
params：pagesize page name

*/
func (u User) list(c *gin.Context){
	var lib libservers.User
	Pageindex := GetPage(c)
	lib.Pagesize = GetPagesize(c)
	lib.Offset = (Pageindex-1)*lib.Pagesize
	users := lib.GetList()
	if users["list"] == nil{
		response.ShowFailed(c,nil,"数据不存在！")
		return
	}
	response.ShowSucc(c,gin.H{"code":200,"data":users},"获取成功")
}

/*
@添加用户
params forms.User
*/
func (u User) add(c *gin.Context){
	var lib libservers.User
	var form forms.User
	err := c.BindJSON(&form)
	if err != nil {
		response.ShowFailed(c, nil, "请求参数错误"+"发生错误:"+err.Error())
		return
	}
	v := validate.Struct(form)
	if !v.Validate(){
		response.ShowFailed(c, nil, "请求参数错误:"+v.Errors.One())
		return
	}
	lib.Add(form)
	response.ShowSucc(c,gin.H{"code":200},"添加成功！")
}
/*
@查询一条数据
params id
*/
func (u User) show(c *gin.Context){
	var lib libservers.User
	lib.Id = c.Param("id")
	user,err := lib.Show()
	if err!=nil{
		response.ShowFailed(c,nil,"数据不存在！")
		return
	}
	response.ShowSucc(c,gin.H{"code":200,"data":user},"获取成功")
}
/*
@删除数据
params id
*/
func (u User) delete(c *gin.Context){
	var lib libservers.User
	lib.Id = c.Param("id")
	res := lib.Del()
	if !res{
		response.ShowFailed(c,nil,"数据不存在！")
		return
	}
	response.ShowSucc(c,gin.H{"code":200,"state":true},"删除成功")
}
/*
@修改数据
params  forms.UpUser
*/
func (u User) update(c *gin.Context){
	var lib libservers.User
	var form forms.UpUser
	err := c.BindJSON(&form)
	if err != nil {
		response.ShowFailed(c, nil, "请求参数错误"+"发生错误:"+err.Error())
		return
	}
	v := validate.Struct(form)
	if !v.Validate(){
		response.ShowFailed(c, nil, "请求参数错误:"+v.Errors.One())
		return
	}
	lib.Id = form.Id
	log.Println(form)
	err = lib.Update(form)
	if err!=nil{
		response.ShowFailed(c,nil,"数据不存在")
		return
	}
	response.ShowSucc(c,gin.H{"code":200,"state":true},"修改成功")
}