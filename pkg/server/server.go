package server

import (
	"fmt"
	"github.com/soxueren/greenplum-operator/pkg/setting"
	"strings"
	"log"
	"time"
	"github.com/chilts/sid"
	"github.com/gin-gonic/gin"
	"github.com/micro/go-micro/registry"
	consul "github.com/micro/go-micro/registry/consul"	
	"github.com/micro/go-micro/web"
)


func Start() web.Service {

	if setting.ServerSetting.RunMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	hosts := strings.Split(setting.ConsulSetting.Hosts, ",")
	consl := consul.NewRegistry(func(o *registry.Options) {
		o.Addrs = hosts
	})

	log.Printf("Agent Server Listening on :%v", setting.ServerSetting.Addr)

	service := web.NewService(
		web.Id(fmt.Sprintf("%s-%s", setting.ServerSetting.Name, sid.Id())),
		web.Name(setting.ServerSetting.Name),
		web.Version(setting.ServerSetting.Version),
		web.Address(setting.ServerSetting.Addr),		
		web.RegisterTTL(time.Second*30),
		web.RegisterInterval(time.Second*10),
	)	

	service.Init(func(o *web.Options) {
		o.Registry = consl
	})

	return service
}

func RegistryRouter(service web.Service,router *gin.Engine) {
	log.Printf("Agent Server Dispatch Root Path")
	// Run server
	if err := service.Run(); err != nil {	
		log.Printf("Agent Server error")	
		log.Fatal(err)
	}
}