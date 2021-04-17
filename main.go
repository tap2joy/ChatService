package main

import (
	"fmt"
	"log"
	"net"

	"go.elastic.co/apm/module/apmgrpc"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"
	"google.golang.org/grpc/reflection"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_validator "github.com/grpc-ecosystem/go-grpc-middleware/validator"
	_ "github.com/lib/pq"

	"github.com/tap2joy/ChatService/server"
	"github.com/tap2joy/ChatService/services"
	"github.com/tap2joy/ChatService/utils"
	pb "github.com/tap2joy/Protocols/go/grpc/chat"
)

func main() {
	go InitService()

	host := utils.GetString("app", "host")
	port := utils.GetInt("app", "port")
	lis, err := net.Listen("tcp", ":9101")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer(grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
		apmgrpc.NewUnaryServerInterceptor(apmgrpc.WithRecovery()),
		grpc_validator.UnaryServerInterceptor())))

	pb.RegisterChatServiceServer(s, &server.Server{})
	grpc_health_v1.RegisterHealthServer(s, &server.HealthServer{})
	reflection.Register(s)
	fmt.Printf("ChatService start at host %s, port %d\n", host, port)
	s.Serve(lis)
}

func InitService() {
	services.GetChatMgr().InitChannels()

	// 注册服务
	localAddress := utils.GetLocalAddress()
	err := services.RegisterChatService(localAddress)
	if err != nil {
		fmt.Printf("register chat service failed, err = %v\n", err)
	} else {
		fmt.Println("register chat service success")

		// 启动定时心跳，1秒一次
		utils.StartTimer(1, "2021-01-01 19:14:30", "", func() {
			services.SendHeartBeat()
		})
		select {}
	}
}
