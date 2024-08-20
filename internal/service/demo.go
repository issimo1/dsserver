package service

import (
	"context"
	"github.com/issimo1/dsserver/internal/biz"

	pb "github.com/issimo1/dsserver/api/thanos"
)

type DemoService struct {
	pb.UnimplementedDemoServer

	uc *biz.GreeterUsecase
}

func NewDemoService(uc *biz.GreeterUsecase) *DemoService {
	return &DemoService{
		uc: uc,
	}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	return &pb.CreateDemoReply{}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
	return &pb.UpdateDemoReply{}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
	return &pb.DeleteDemoReply{}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	return &pb.GetDemoReply{}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
	return &pb.ListDemoReply{}, nil
}
func (s *DemoService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	g, err := s.uc.CreateGreeter(ctx, &biz.Greeter{Hello: req.Name})
	if err != nil {
		return nil, err
	}
	return &pb.HelloReply{
		Message: "Hello " + g.Hello,
	}, nil
}
