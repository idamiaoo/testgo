package example

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"

	"testing"

	"github.com/bohler/sd"
	"github.com/micro/go-micro/registry"
	"github.com/micro/go-plugins/registry/etcdv3"
	"github.com/micro/misc/lib/addr"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"

	etcdsd "github.com/bohler/lib/etcd-sd/grpc"
	etcdclient "github.com/coreos/etcd/clientv3"
)

type server struct {
}

func (s *server) SayHello(ctx context.Context, in *Req) (*Rsp, error) {
	out := new(Rsp)
	if in.GetName() == "" {
		out.Code = 500
		log.Println("nil name")
		return out, fmt.Errorf("nil name")
	}
	out.Resp = "hello " + in.GetName()
	log.Println(out.Resp)
	return out, nil
}

func (s *server) Crazy(stream ServiceTest_CrazyServer) error {
	log.Println(stream.Context().Deadline)
	go func() {
		for {
			r, err := stream.Recv()
			if err != nil {
				log.Println(err)
				return
			}
			out := new(Rsp)
			out.Resp = "hello " + r.GetName()
			stream.Send(out)
		}
	}()
	return nil
}

func TestServer(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Println(err)
		return
	}
	gs := grpc.NewServer()
	//RegisterServiceTestServer(gs, s)
	go gs.Serve(lis)
	s := new(server)
	RegisterServiceTestServer(gs, s)
	log.Println(lis.Addr().String())

	var host string
	var port int
	parts := strings.Split(lis.Addr().String(), ":")
	if len(parts) > 1 {
		host = strings.Join(parts[:len(parts)-1], ":")
		port, _ = strconv.Atoi(parts[len(parts)-1])
	} else {
		host = parts[0]
	}

	address, err := addr.Extract(host)
	// register service
	node := &registry.Node{
		Id:       "test" + "-" + "abcd",
		Address:  address,
		Port:     port,
		Metadata: nil,
	}
	log.Println(node.Address, node.Port)
	service := &registry.Service{
		Name:      "test",
		Version:   "1",
		Nodes:     []*registry.Node{node},
		Endpoints: nil,
	}

	re := etcdv3.NewRegistry()
	//log.Println(re.String())
	//registry.Register(service, registry.RegisterTTL(5*time.Second))
	re.Register(service)
	select {}
}

func TestSDServer(t *testing.T) {
	sd.Service = sd.NewService(
		sd.Name("test"),
		sd.ID(viper.GetString("1")),
		sd.AfterStart(func() error { //注册服务
			RegisterServiceTestServer(sd.Service.Server().Server(), &server{})
			return nil
		}),
	)
	if err := sd.Service.Run(); err != nil {
		log.Fatal(err)
	}

}

func TestSDServerV2(t *testing.T) {
	lis, err := net.Listen("tcp", ":0")
	if err != nil {
		log.Println(err)
		return
	}
	gs := grpc.NewServer()
	//RegisterServiceTestServer(gs, s)
	go gs.Serve(lis)
	s := new(server)
	RegisterServiceTestServer(gs, s)
	log.Println(lis.Addr().String())
	var host string
	var port int
	parts := strings.Split(lis.Addr().String(), ":")
	if len(parts) > 1 {
		host = strings.Join(parts[:len(parts)-1], ":")
		port, _ = strconv.Atoi(parts[len(parts)-1])
	} else {
		host = parts[0]
	}

	address, err := addr.Extract(host)
	addr := fmt.Sprintf("%s:%d", address, port)
	etcdsd.RegisterService(etcdclient.Config{
		Endpoints:   []string{"127.0.0.1:2379"},
		DialTimeout: 3 * time.Second,
	}, "test", addr)
	select {}
}
