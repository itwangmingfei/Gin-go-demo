package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func Respose(ctx *gin.Context,httpstatus int,code int,data gin.H,msg string){
	ctx.JSON(httpstatus,gin.H{"code":code,"data":data,"msg":msg})
}
func ShowSucc(ctx *gin.Context,data gin.H,msg string){
	ctx.JSON(http.StatusOK,gin.H{"code":200,"data":data,"msg":msg})
}
func ShowFailed(ctx *gin.Context,data gin.H,msg string){
	ctx.JSON(http.StatusOK,gin.H{"code":400,"data":data,"msg":msg})
}
