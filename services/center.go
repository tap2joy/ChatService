package services

import (
	"context"
	"time"

	"github.com/tap2joy/ChatService/utils"
	protocols "github.com/tap2joy/Protocols/go/grpc/center"
	"google.golang.org/grpc"
)

func RegisterChatService(chatAddress string) error {
	address := utils.GetString("grpc", "grpc_center_address")

	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := protocols.NewCenterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.RegisterService(ctx, &protocols.RegisterServiceRequest{Type: "chat", Address: chatAddress})
	if err != nil {
		return err
	}

	return nil
}

func GetAllOnlineUsers(channelId uint32) (*protocols.GetOnlineUsersResponse, error) {
	address := utils.GetString("grpc", "grpc_center_address")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := protocols.NewCenterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.GetOnlineUsers(ctx, &protocols.GetOnlineUsersRequest{Channel: channelId})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func GetUserOnlineTime(name string) (*protocols.GetUserOnlineTimeResponse, error) {
	address := utils.GetString("grpc", "grpc_center_address")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := protocols.NewCenterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	resp, err := c.GetUserOnlineTime(ctx, &protocols.GetUserOnlineTimeRequest{Name: name})
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func SendHeartBeat() error {
	address := utils.GetString("grpc", "grpc_center_address")

	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		return err
	}
	defer conn.Close()
	c := protocols.NewCenterServiceClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	_, err = c.HeartBeat(ctx, &protocols.HeartBeatRequest{Type: "chat", Address: utils.GetLocalAddress()})
	if err != nil {
		return err
	}

	return nil
}
