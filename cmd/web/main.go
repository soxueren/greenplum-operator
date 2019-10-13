package main

import(
	"flag"
	"github.com/soxueren/greenplum-operator/pkg/routers"
	"github.com/soxueren/greenplum-operator/pkg/server"
	"github.com/DeanThompson/ginpprof"
	
)

func main() {   

	srv := server.Start()
	router := routers.InitRouter()
	ginpprof.Wrap(router)
	server.RegistryRouter(srv, router)	
}