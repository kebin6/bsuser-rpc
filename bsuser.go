package main

import (
	"flag"
	"fmt"

	"github.com/kebin6/bsuser-rpc/internal/config"
	"github.com/kebin6/bsuser-rpc/internal/server"
	"github.com/kebin6/bsuser-rpc/internal/svc"
	"github.com/kebin6/bsuser-rpc/types/bsuser"

	"github.com/zeromicro/go-zero/core/conf"
	"github.com/zeromicro/go-zero/core/service"
	"github.com/zeromicro/go-zero/zrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

var configFile = flag.String("f", "etc/bsuser.yaml", "the config file")

func main() {
	flag.Parse()

	var c config.Config
	conf.MustLoad(*configFile, &c, conf.UseEnv())
	ctx := svc.NewServiceContext(c)

	s := zrpc.MustNewServer(c.RpcServerConf, func(grpcServer *grpc.Server) {
		bsuser.RegisterBsuserServer(grpcServer, server.NewBsuserServer(ctx))

		if c.Mode == service.DevMode || c.Mode == service.TestMode {
			reflection.Register(grpcServer)
		}
	})
	defer s.Stop()

	fmt.Printf("Starting rpc server at %s...\n", c.ListenOn)
	s.Start()
}
