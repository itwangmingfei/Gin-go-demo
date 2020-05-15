package controllers

import (
	"fmt"
	"gin/config"
	"gin/models"
	"gin/tools"
	"github.com/gin-gonic/gin"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/queue"
	"github.com/gocolly/redisstorage"
	"log"
	"regexp"
	"time"
)

/*
@爬虫
*/

//用户控制器
type Gocoll struct {
}

func initGocollRouteAr(r *gin.Engine) {
	col := new(Gocoll)
	AppGroup := r.Group("/Gocoll/v1")
	AppGroup.GET("/get", col.get)
}
/*
@爬取七猫小说网
*/
func (g Gocoll) get(c *gin.Context){
	dlink :="https://www.qimao.com"
	cfg := config.GetConfig()
	var novel models.Novel
	var clientredis tools.GoRedis
	//*********************************

	coll := colly.NewCollector(
		colly.AllowedDomains("www.qimao.com"),
		colly.Async(true),
		colly.UserAgent("Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"),
	)

	///这个地方是配置redis存储请求返回的数据信息吧？

	storage := &redisstorage.Storage{
		Address: fmt.Sprintf("%s:%s", cfg.Redis.Host,cfg.Redis.Port),
		Password:cfg.Redis.Passwd,
		DB: cfg.Redis.Db,
		Prefix : "HTTP_QIMAO",
	}
	if err := coll.SetStorage(storage);err!=nil{
		panic(err)
	}
	if err:=storage.Clear();err!=nil{
		log.Fatal(err)
	}
	//defer storage.Client.Close()
	//*----------------------------------------------

	q,_ :=queue.New(2,storage)
	/*<h2 class="tit">????</h2>*/
	/*获取样式 div.data-txt 模块下的数据信息*/
	coll.OnHTML(`div.data-txt`, func(e *colly.HTMLElement) {
		ls,_ := e.DOM.Html()

		//标题
		tit := e.ChildText("h2.tit")
		//作者
		pname := e.ChildText(`p.p-name a`)
		//状态
		status := e.ChildText(`span.qm-tags.black.clearfix em:first-child`)

		nums := e.ChildText(`p.p-num span:nth-child(1)`)
		nums1 := e.ChildText(`p.p-num span:nth-child(3)`)
		nums2 := e.ChildText(`p.p-num span:nth-child(5)`)

		//没有标签正则一下
		reg := regexp.MustCompile(`<em>主角：</em>(.*?)<`)
		regstr := reg.FindAllStringSubmatch(ls,-1)

		uptime := e.ChildText(`p.p-update em.time`)

		newpage :=e.ChildText(`p.p-update a`)


		if len(tit)>0 {
			novel.Title = tit
			fmt.Printf("名称：%s \n",tit)
		}
		if len(regstr) >0 {
			novel.Mster = regstr[0][1]
			fmt.Printf("主角：%s \n",regstr[0][1])
		}
		if len(pname) >0 {
			fmt.Printf("作者：%s \n",pname)
			novel.Author = pname
		}
		if len(status) >0 {
			fmt.Printf("状态：%s \n",status)
			novel.Status = status
		}
		if len(nums)>0{
			fmt.Printf("数量：%s %s %s \n",nums,nums1,nums2)
			novel.Show = nums+nums1+nums2
		}
		if len(uptime)>0{
			fmt.Printf("跟新时间：%s \n",uptime)
			novel.Uptime =uptime

		}
		if len(newpage)>0{
			fmt.Printf("最新章节：%s \n",newpage)
			novel.Newpage = newpage
		}
		models.GetDb().Create(&novel)
	})

	//获取url
	coll.OnHTML(`a[href]`,func(e *colly.HTMLElement){
		link := e.Attr("href")
		//匹配书库
		reg := regexp.MustCompile(`^(/shuku/[1-9]{6}/)`)
		res := reg.FindAllString(link,-1)
		//存在返回数据
		if len(res)>0{
			//存入redis中
			newLink :=  dlink +link
			//判断执行存储的redis中是否存在如果存在不存储当前url
			Isset := clientredis.ToIsset(newLink)

			//如果不存在继续抓取当前链接
			if !Isset{
				//连接存入redis中
				log.Println(newLink)
				err := clientredis.DoLpush(newLink)
				if err!=nil{
					log.Println(err)
				}
			}

		}
	})

	coll.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	coll.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	//获取这个应该是配饰redisstorage使用 用于接收返回数据然后存储
	coll.OnResponse(func(r *colly.Response) {
		log.Println(coll.Cookies(r.Request.URL.String()))
	})

	novel.Url = dlink
	q.AddURL(dlink)
	q.Run(coll)
	for {
		dlink,_ := clientredis.DoRpop()
		if len(dlink)!=0{
			novel.Url = dlink
			q.AddURL(dlink)
			q.Run(coll)
			time.Sleep(time.Second*2)
		}

	}

}