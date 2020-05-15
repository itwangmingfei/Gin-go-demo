package config

import (
	"encoding/json"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
)

type Server struct {
	Listen     string `json:"listen"`
	JwtSecret  string `json:"jwt_secret"`
	CookieSalt string `json:"cookie_salt"`
	Debug      bool   `json:"debug"`
}

type Mysql struct {
	Host         string `json:"host"`
	Port         string `json:"port"`
	Db           string `json:"db"`
	User         string `json:"user"`
	Passwd       string `json:"passwd"`
	MaxIdleConns int    `json:"max_idle_conns"`
	MaxOpenConns int    `json:"max_o_pen_conns"`
}
type Redis struct {
	Host   string `json:"host"`
	Port   string `json:"port"`
	Passwd string `json:"passwd"`
	Db     int    `json:"db"`
}
type Config struct {
	Server Server `json:"server"`
	Mysql  Mysql  `json:"mysql"`
	Redis  Redis  `json:"redis"`
}

var cfg Config

func initConfig(path string) {

	path = path + "/config/config.json"
	b, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("cant not load config from: %s", path)
	}

	if err = json.Unmarshal(b, &cfg); err != nil {
		log.Fatalf("Unmarshal failed : %s", err.Error())
	}
}
func Load(dir string) {
	initConfig(dir)
}
func GetConfig() *Config {
	return &cfg
}
