package main

import (
	"flag"
	"fmt"
	"log"

	"chainproxy/internal/config"
	"chainproxy/internal/handler"
	"chainproxy/internal/svc"

	"github.com/tal-tech/go-zero/core/conf"
	"github.com/tal-tech/go-zero/rest"
)

var configFile = flag.String("f", "etc/chainproxy-api.yaml", "the config file")

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c)

	ctx := svc.NewServiceContext(c)
	server := rest.MustNewServer(c.RestConf)
	defer server.Stop()

	handler.RegisterHandlers(server, ctx)

	fmt.Printf("Starting server at %s:%d...\n", c.Host, c.Port)
	server.Start()
}
