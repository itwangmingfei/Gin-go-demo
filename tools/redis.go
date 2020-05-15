package tools

import (
	"fmt"
	"gin/config"
	"github.com/monnand/goredis"
	log "github.com/sirupsen/logrus"
)

type GoRedis struct {
	*goredis.Client
}

var Client GoRedis

const (
	DO_QUEUE = "DoQueue" //记录需要执行的数据
	TO_QUEUE = "Toqueue" //处理过的数据放这里
)

//链接redis
func InitRedis(cfg config.Redis) {
	var cliect  goredis.Client
	cliect.Addr = fmt.Sprintf("%s:%s", cfg.Host, cfg.Port)
	cliect.Password = cfg.Passwd
	cliect.Db = cfg.Db
	pong, err := cliect.Ping()
	if err != nil {
		log.Fatalf("cant ping redis %s\n", err.Error())
	}
	Client = GoRedis{&cliect}
	log.Println("ping redis:", pong)
}

//存入数据头部存入数据
func (r GoRedis) DoLpush(value string) error {
	err := Client.Lpush(DO_QUEUE, []byte(value))
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//读取数据尾部获取数据
func (r GoRedis) DoRpop() (string, error) {
	val, err := Client.Rpop(DO_QUEUE)
	if err != nil {
		return "", err
	}
	return string(val), err
}

//获取当前key存入的数量
func (r GoRedis) DoLen() int {
	nums, err := Client.Llen(DO_QUEUE)
	if err != nil {
		return 0
	}
	return nums
}

//存入集合中------------------------------------
func (r GoRedis) ToSadd(value string) bool {
	res, _ := Client.Sadd(TO_QUEUE, []byte(value))
	return res
}

//获取当前value是否需要存
func (r GoRedis) ToIsset(value string) bool {
	res, _ := Client.Sismember(TO_QUEUE, []byte(value))
	return res
}

//---------------------------------------------------
