package messaging_service

import (
	"context"
	"go_im/im/message"
	"go_im/im/messaging"
	"go_im/protobuff/gen/pb_rpc"
	"go_im/service/rpc"
)

type Server struct {
	*rpc.BaseServer
}

func NewServer(options *rpc.ServerOptions) *Server {
	s := &Server{
		BaseServer: rpc.NewBaseServer(options),
	}
	s.Register(options.Name, s)
	return s
}

func (s *Server) HandleMessage(ctx context.Context, request *pb_rpc.MessagingHandleRequest, replay *pb_rpc.Response) error {
	m := &message.Message{CommMessage: request.GetMessage()}
	messaging.HandleMessage(request.GetId(), request.GetDevice(), m)
	return nil
}