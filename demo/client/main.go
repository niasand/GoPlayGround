package main

import (
	pb "demo/demo_proto"
	"log"
	"os"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address     = "localhost:10000"
	defaultName = "Zhiwei Yang"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())

	if err != nil {
		log.Fatal("did not connect:%v", err)
	}

	defer conn.Close()
	c := pb.NewHelloClient(conn)

	name := defaultName

	if len(os.Args) > 1 {
		name = os.Args[1]
	}

	r, err := c.SendHelpInfoWithGetMethod(context.Background(), &pb.DemoProtoHelloRequest{Name: name})
	if err != nil {
		log.Fatal("打不了招呼，错误:%v", err)
	}
	log.Printf("客户端 get方法:%s", r.Message)

	r, err = c.SendInfoWithPostMethod(context.Background(), &pb.DemoProtoHelloRequest{Name: name})

	if err != nil {
		log.Fatal("不能打招呼啦，;%v", err)
	}

	log.Printf("客户端发请求:%s", r.Message)
}
