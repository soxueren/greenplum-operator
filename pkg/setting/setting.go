package setting

import (
	"log"
	"time"

	"github.com/go-ini/ini"
)

type App struct {
	JwtSecret string
	PageSize  int

	RuntimeRootPath  string
	DownloadRootPath string
	ImageSavePath    string
	ImageMaxSize     int
	ImageAllowExts   []string

	ExportSavePath string
	QrCodeSavePath string
	FontSavePath   string

	LogSavePath string
	LogSaveName string
	LogFileExt  string
	TimeFormat  string
}

var AppSetting = &App{}

type Server struct {
	Name         string
	Addr         string
	Version      string
	ReadTimeout  time.Duration
	WriteTimeout time.Duration
	RunMode      string
	RootDir      string
}

var ServerSetting = &Server{}

type Consul struct {
	Hosts    string
	AclToken string
}

var ConsulSetting = &Consul{}

var cfg *ini.File

func Setup() {
	var err error
	cfg, err = ini.Load("conf/app.ini")
	if err != nil {
		log.Fatalf("Fail to parse 'conf/app.ini': %v", err)
	}

	mapTo("app", AppSetting)
	mapTo("consul", ConsulSetting)
	mapTo("server", ServerSetting)	

	AppSetting.ImageMaxSize = AppSetting.ImageMaxSize * 1024 * 1024
	ServerSetting.ReadTimeout = ServerSetting.ReadTimeout * time.Second
	ServerSetting.WriteTimeout = ServerSetting.ReadTimeout * time.Second
}

func mapTo(section string, v interface{}) {
	err := cfg.Section(section).MapTo(v)
	if err != nil {
		log.Fatalf("Cfg.MapTo RedisSetting err: %v", err)
	}
}
