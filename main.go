package main

import (
	"net"
	"platform-data/config"
	"platform-data/handler"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	"github.com/xh-polaris/gopkg/kitex/middleware"
	logx "github.com/xh-polaris/gopkg/util/log"
	data "github.com/xh-polaris/service-idl-gen-go/kitex_gen/platform/data/dataservice"
)

func main() {

	config.Init()

	klog.SetLogger(logx.NewKlogLogger())

	addr, err := net.ResolveTCPAddr("tcp", config.Get().ListenOn)

	if err != nil {
		panic(err)
	}
	svr := data.NewServer(
		handler.NewHandler(),
		server.WithServiceAddr(addr),
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: config.Get().Name}),
		server.WithMiddleware(middleware.LogMiddleware(config.Get().Name)),
	)

	err = svr.Run()

	if err != nil {
		logx.Error(err.Error())
	}
}
