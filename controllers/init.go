package controllers

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"strconv"
	"strings"
)

//定义路由
func Init(r *gin.Engine) {
	initUserRouteAr(r)
}

func GetPage(c *gin.Context) int{
	page := 1
	_p := strings.TrimSpace(c.Query("page"))
	log.Infof("page : %s", _p)
	if _p != "" {
		_converted, err := strconv.Atoi(_p)
		if err == nil && _converted > 0 {
			page = _converted
		}
	}
	return page
}
func GetPagesize(c *gin.Context) int{
	pagesize := 10
	_p := strings.TrimSpace(c.Query("pagesize"))
	log.Infof("pagesize :%s", _p)
	if _p != "" {
		_converted, err := strconv.Atoi(_p)
		if err == nil && _converted > 0 {
			log.Errorf("err :%v", err)
			pagesize = _converted
		}
	}
	return pagesize
}