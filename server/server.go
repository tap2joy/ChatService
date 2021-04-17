package server

import (
	"context"

	"github.com/tap2joy/ChatService/dao"
	"github.com/tap2joy/ChatService/services"
	pb_common "github.com/tap2joy/Protocols/go/common"
	pb "github.com/tap2joy/Protocols/go/grpc/chat"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Server struct {
}

func (*Server) SendMessage(ctx context.Context, req *pb.SendMessageRequest) (*pb.SendMessageResponse, error) {
	senderName := req.SenderName
	channel := req.Channel
	content := req.Content
	isSystem := req.System

	retStr, err := services.GetChatMgr().SendMessage(channel, senderName, content, isSystem)
	if err != nil {
		return nil, err
	}

	resp := new(pb.SendMessageResponse)
	resp.Result = retStr
	return resp, nil
}

func (*Server) GetChatLog(ctx context.Context, req *pb.GetChatLogRequest) (*pb.GetChatLogResponse, error) {
	channelId := req.Channel

	channel, err := dao.GetChannel(channelId)
	if err != nil {
		return nil, err
	}

	if channel == nil {
		return nil, status.Errorf(codes.Code(pb_common.ErrorCode_CHANNEL_NOT_EXIST_ERROR), "channel not exist")
	}

	logs, err := services.GetChatMgr().GetChatLog(channelId)
	if err != nil {
		return nil, err
	}

	resp := new(pb.GetChatLogResponse)
	for _, v := range logs {
		resp.Logs = append(resp.Logs, &pb.ChatLogInfo{SenderName: v.SenderName, Timestamp: uint64(v.CreateAt.Unix()), Content: v.Content})
	}

	return resp, nil
}

func (*Server) GetChannelList(ctx context.Context, req *pb.GetChannelListRequest) (*pb.GetChannelListResponse, error) {
	channels, err := dao.GetChannelList(0, 100)
	if err != nil {
		return nil, err
	}

	resp := new(pb.GetChannelListResponse)
	for _, v := range channels {
		resp.List = append(resp.List, &pb.ChannelInfo{Id: v.Id, Desc: v.Desc})
	}

	return resp, nil
}
