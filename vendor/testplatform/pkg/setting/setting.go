package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File
	// RUN_MODE
	RunMode string
	// server
	HTTPPort int
	ReadTimeout time.Duration
	WriteTimeout time.Duration
	// app
	PageSize int
	JwtSecret string
	// redis
	RedisHost string
	RedisPassword string
	RedisMaxIdle int
	RedisMaxActive int
	RedisIdleTimeout time.Duration
)

func init() {
	var err error
	Cfg, err = ini.Load("vendor/testplatform/conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

// 解析RUN_MODE
func LoadBase() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

// 解析server
func LoadServer() {
	sec, err := Cfg.GetSection("server")
	if err != nil {
		log.Fatalf("Fail to get section 'server': %v", err)
	}

	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")

	HTTPPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeout = time.Duration(sec.Key("READ_TIME").MustInt(60)) * time.Second
	WriteTimeout = time.Duration(sec.Key("WRITE_TIMEOUT").MustInt(60)) * time.Second
}

// 解析app
func LoadApp() {
	sec, err := Cfg.GetSection("app")
	if err != nil {
		log.Fatalf("Fail to get section 'app': %v", err)
	}

	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}

// 解析redis
func LoadRedis() {
	sec, err := Cfg.GetSection("redis")
	if err != nil {
		log.Fatalf("Fail to get section 'redis': %v", err)
	}
	
	RedisHost = sec.Key("Host").MustString("172.17.3.185:6379")
	RedisPassword = sec.Key("Password").MustString("")
	RedisMaxIdle = sec.Key("MaxIdle").MustInt(30)
	RedisMaxActive = sec.Key("MaxActive").MustInt(30)
	RedisIdleTimeout = time.Duration(sec.Key("IdleTimeout").MustInt(200)) * time.Second
}