package grpc

import (
	"net"
	"smartcard/config"
	api "smartcard/pkg/grpc/api"
	log "smartcard/pkg/logging"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	api.UnimplementedScannerSmartCardServer
}

func Run() {
	server := grpc.NewServer()

	grpcServer := &GRPCServer{}
	api.RegisterScannerSmartCardServer(server, grpcServer)

	listen, err := net.Listen("tcp", config.Cfg.GRPC_LISTEN_PORT)
	if err != nil {
		log.Logrus.Fatalf("Error establishing tcp connection on port = %v, error = %v", config.Cfg.GRPC_LISTEN_PORT, err)
	}
	log.Logrus.Debug("GRPC Server ready to accept request on port = ", config.Cfg.GRPC_LISTEN_PORT)
	err = server.Serve(listen)
	if err != nil {
		log.Logrus.Fatal(err)
	}
}
