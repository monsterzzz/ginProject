package setting

import (
	"gopkg.in/ini.v1"
	"log"
	"time"
)

var (
	Cfg *ini.File
	RunMode string
	HttpPort int
	ReadTimeOut time.Duration
	WriteTimeOuT time.Duration
	PageSize int
	JwtSecret string
	err error
)


func init(){
	Cfg,err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("ini error, %v",err)
	}

	LoadBase()
	LoadServer()
	LoadApp()
}

func LoadBase(){
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
}

func LoadServer(){
	sec,err := Cfg.GetSection("server")
	if err != nil{
		log.Fatalf("error section, %v",err)
	}
	HttpPort = sec.Key("HTTP_PORT").MustInt(8000)
	ReadTimeOut = time.Duration(sec.Key("READ_TIMEOUT").MustDuration(60)) * time.Second
	WriteTimeOuT = time.Duration(sec.Key("WRITE_TIMEOUT").MustDuration(60)) * time.Second
}

func LoadApp(){
	sec,err := Cfg.GetSection("app")
	if err != nil{
		log.Fatalf("error section, %v",err)
	}
	JwtSecret = sec.Key("JWT_SECRET").MustString("!@)*#)!@U#@*!@!)")
	PageSize = sec.Key("PAGE_SIZE").MustInt(10)
}