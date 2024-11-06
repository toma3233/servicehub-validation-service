// Auto generated. Don't modify.
package demoserver

import (
	"fmt"
	log "log/slog"
	"net"
	"os"

	pb "<<apiModule .envInformation.goModuleNamePrefix .serviceInput.directoryName>>/v1"
	"<<serverModule .envInformation.goModuleNamePrefix .serviceInput.directoryName>>/internal/logattrs"
	"github.com/Azure/aks-middleware/interceptor"
	"google.golang.org/grpc"
)

func Serve(options Options) {
	logger := log.New(log.NewTextHandler(os.Stdout, nil).WithAttrs(logattrs.GetAttrs()))
	if options.JsonLog {
		logger = log.New(log.NewJSONHandler(os.Stdout, nil).WithAttrs(logattrs.GetAttrs()))
	}

	log.SetDefault(logger)

	apiServer := NewServer()
	apiServer.init(options)

	gRpcServer := grpc.NewServer(grpc.ChainUnaryInterceptor(
		interceptor.DefaultServerInterceptors(interceptor.GetServerInterceptorLogOptions(logger, logattrs.GetAttrs()))...,
	))
	pb.Register<<.serviceInput.serviceName>>Server(gRpcServer, apiServer)

	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", options.Port))
	if err != nil {
		// slog does not have "Fatal" method
		logger.Error("failed to listen: " + err.Error())
		os.Exit(1)
	}
	logger.Info(fmt.Sprintf("demoserver listening at %s", listener.Addr().String()))
	if err := gRpcServer.Serve(listener); err != nil {
		logger.Error("failed to serve: " + err.Error())
		os.Exit(1)
	}
}
