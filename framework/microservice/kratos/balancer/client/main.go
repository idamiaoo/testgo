package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"testgo/framework/microservice/kratos/balancer/helloworld"
	"testgo/framework/microservice/kratos/balancer/sessionkeep"
	"testgo/framework/microservice/kratos/balancer/sessionkeep/session"

	"github.com/go-kratos/etcd/registry"
	"github.com/go-kratos/kratos/v2/transport/grpc"
	"github.com/go-redis/redis/v8"
	clientv3 "go.etcd.io/etcd/client/v3"
	ggrpc "google.golang.org/grpc"
	"google.golang.org/grpc/balancer"
	"google.golang.org/grpc/metadata"
)

func main() {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints: []string{"127.0.0.1:2379"},
	})
	if err != nil {
		panic(err)
	}

	rds := redis.NewClient(&redis.Options{
		Addr: "127.0.0.1:6379",
	})

	balancer.Register(sessionkeep.NewBuilder(session.NewSessionMgr(rds)))

	r := registry.New(cli)
	conn, err := grpc.DialInsecure(
		context.Background(),
		grpc.WithEndpoint("discovery:///helloworld"),
		grpc.WithDiscovery(r),
		grpc.WithOptions(ggrpc.WithBalancerName(sessionkeep.Name)),
	)
	if err != nil {
		log.Fatal(err)
	}
	md := metadata.New(nil)
	// session.MDWithSessionID(md, "carl")
	ctx := context.Background()
	ctx = metadata.NewOutgoingContext(ctx, md)
	client := helloworld.NewGreeterClient(conn)
	var i int
	for i < 10 {
		reply, err := client.SayHello(ctx, &helloworld.HelloRequest{Name: fmt.Sprintf("kratos %d", i)})
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("[grpc] SayHello %+v\n", reply)
		i++
		time.Sleep(1 * time.Second)
	}
}
