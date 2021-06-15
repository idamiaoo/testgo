package main

import (
	"context"
	"fmt"
	"os"

	pb "testgo/framework/microservice/kratos/balancer/helloworld"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2"
	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-kratos/kratos/v2/middleware/logging"
	"github.com/go-kratos/kratos/v2/middleware/recovery"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-kratos/kratos/v2/transport/http"
	etcd "go.etcd.io/etcd/client/v3"
)

// go build -ldflags "-X main.Version=x.y.z"
var (
	// Name is the name of the compiled software.
	Name = "helloworld"
	// Version is the version of the compiled software.
	Version = "v1.0.0"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	if in.Name == "error" {
		return nil, errors.BadRequest("custom_error", fmt.Sprintf("invalid argument %s", in.Name))
	}
	if in.Name == "panic" {
		panic("grpc panic")
	}
	return &pb.HelloReply{Message: fmt.Sprintf("Hello %+v", in.Name)}, nil
}

func main() {
	logger := log.NewStdLogger(os.Stdout)

	log := log.NewHelper(logger)

	grpcSrv := grpc.NewServer(
		grpc.Address(":0"),
		grpc.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		))

	s := &server{}
	pb.RegisterGreeterServer(grpcSrv, s)

	cli, err := etcd.New(etcd.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		log.Error(err)
		return
	}

	r := registry.New(cli)

	httpSrv := http.NewServer(http.Address(":0"))
	httpSrv.HandlePrefix("/", pb.NewGreeterHandler(s,
		http.Middleware(
			recovery.Recovery(),
			logging.Server(logger),
		)),
	)

	app := kratos.New(
		kratos.Name(Name),
		kratos.Server(
			httpSrv,
			grpcSrv,
		),
		kratos.Registrar(r),
	)

	if err := app.Run(); err != nil {
		log.Error(err)
	}
}
