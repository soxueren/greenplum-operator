package server

import (
	"fmt"
	"storageagent/pkg/setting"
	"strings"
	"time"

    "log"
	"github.com/chilts/sid"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	consul "github.com/micro/go-micro/registry/consul"
	web "github.com/micro/go-web"
)

func Start() web.Service {

	if setting.ServerSetting.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	hosts := strings.Split(setting.ConsulSetting.Hosts, ",")
	consl := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = hosts
	})

	service := web.NewService(
		web.Id(fmt.Sprintf("%s-%s", setting.ServerSetting.Name, sid.Id())),
		web.Name(setting.ServerSetting.Name),
		web.Version(setting.ServerSetting.Version),
		web.Address(setting.ServerSetting.Addr),
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
	)

	log.Printf("Storage Server Listening on :%v", setting.ServerSetting.Addr)

	service.Init(func(o *web.Options) {
		o.Registry = consl
	})

	return service
}

func RegistryRouter(service web.Service, router *gin.Engine) {
	service.Handle("/", router)
	mainLog.Printf("Storage Server Dispatch Root Path")
	// Run server
	if err := service.Run(); err != nil {
		mainLog.Fatal(err)
	}
}
