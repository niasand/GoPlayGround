package main

import (
	"log"
	"net"

	pb "demo/demo_proto"

	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const (
	port = ":10000"
)

type server struct{}

func (s *server) SendHelpInfoWithGetMethod(ctx context.Context, in *pb.DemoProtoHelloRequest) (*pb.DemoProtoHelloResponse, error) {
	return &pb.DemoProtoHelloResponse{Message: "接收GET方法：" + in.Name}, nil
}

func (s *server) SendInfoWithPostMethod(ctx context.Context, in *pb.DemoProtoHelloRequest) (*pb.DemoProtoHelloResponse, error) {
	return &pb.DemoProtoHelloResponse{Message: "接收POST方法" + in.Name}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("监听端口失败: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server{})
	reflection.Register(s)
	if err := s.Serve(lis); err != nil {
		log.Fatal("启动服务失败:%v", err)
	}
}
