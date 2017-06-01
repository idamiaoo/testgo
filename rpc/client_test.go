package example

import (
	"fmt"
	"log"
	"strconv"
	"sync"
	"testing"
	"time"

	etcdsd "github.com/bohler/lib/etcd-sd/grpc"
	"github.com/bohler/sd"
	"github.com/micro/go-micro/selector"
	"github.com/micro/go-plugins/registry/etcdv3"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

func TestClient(t *testing.T) {
	se := selector.NewSelector(selector.Registry(etcdv3.NewRegistry()))
	//se := selector.NewSelector()
	next, err := se.Select("test")
	if err != nil {
		log.Println(err)
		return
	}
	node, err := next()
	if err != nil {
		log.Println(err)
		return
	}
	addr := fmt.Sprintf("%s:%d", node.Address, node.Port)
	log.Println(addr)
	conn, err := grpc.Dial(addr, grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}

	cc := NewServiceTestClient(conn)
	out, err := cc.SayHello(context.Background(), &Req{Name: "holly"})
	if err != nil {
		log.Println(err)
	}
	log.Println(out.Resp)
}

func TestStream(t *testing.T) {
	conn, err := grpc.Dial("127.0.0.1:61206", grpc.WithInsecure())
	if err != nil {
		log.Println(err)
		return
	}
	ctx, _ := context.WithTimeout(context.Background(), 0)
	cc := NewServiceTestClient(conn)
	stream, err := cc.Crazy(ctx)
	if err != nil {
		log.Println(err)
		return
	}

	go func() {
		for {
			r, err := stream.Recv()
			if err != nil {
				t.Fatal(err)
			}
			t.Log(r.GetResp())
		}
	}()

	i := 0
	for {
		<-time.After(1 * time.Second)
		if i > 10 {
			return
		}
		stream.Send(&Req{strconv.Itoa(i)})
		i++
	}
}

func sayHello() error {
	cc, err := etcdsd.GetService("test")
	if err != nil {
		log.Println(err)
		return err
	}
	defer cc.Close()
	cli := NewServiceTestClient(cc)
	once := sync.Once{}
	sf := func() {
		stream, err := cli.Crazy(context.Background())
		if err != nil {
			log.Println(err)
			return
		}

		go func() {
			if err := stream.Send(&Req{Name: "longzl"}); err != nil {
				log.Println("send", err)
				return
			}
			<-time.After(1 * time.Second)
		}()
		for {
			r, err := stream.Recv()
			if err != nil {
				log.Println("recv", err)
				return
			}

			log.Println(r.GetResp())
		}
	}
	for {
		<-time.After(1 * time.Second)
		rsp, err := cli.SayHello(context.Background(), &Req{Name: "bohler"})
		if err != nil {
			log.Println(err)
			continue
		}
		once.Do(sf)
		log.Println(rsp.GetResp())
	}
}

func TestSDClientV2(t *testing.T) {

	sayHello()
	/*
		cc, err := grpc.Dial("192.168.11.9:56643", grpc.WithInsecure())
		if err != nil {
			log.Println(err)
			return
		}
		defer cc.Close()
		rsp, err := NewServiceTestClient(cc).SayHello(context.Background(), &Req{Name: "bohler"})
		if err != nil {
			log.Println(err)
			return
		}
		log.Println(rsp.GetResp())
	*/
}

func TestDial(t *testing.T) {
	_, err := sd.Service.Selector().Select("lobby")
	if err != nil {
		log.Println(err)
		return
	}
}
