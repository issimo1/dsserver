package service

import (
	"context"
	"github.com/issimo1/dsserver/internal/biz"

	pb "github.com/issimo1/dsserver/api/thanos"
)

type DemoService struct {
	pb.UnimplementedDemoServer

	ac *biz.ArticleUseCase
	uc *biz.GreeterUsecase
}

func NewDemoService(uc *biz.GreeterUsecase) *DemoService {
	return &DemoService{
		uc: uc,
	}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	_, err := s.ac.CreateArticle(ctx, biz.Article{
		ID:      req.Id,
		Comment: req.Comment,
		Content: req.Content,
	})
	if err != nil {
		return &pb.CreateDemoReply{
			Ok: false,
		}, err
	}
	return &pb.CreateDemoReply{
		Ok: true,
	}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
	_, err := s.ac.UpdateArticle(ctx, biz.Article{
		ID:      req.Id,
		Comment: req.Comment,
		Content: req.Content,
	})
	if err != nil {
		return &pb.UpdateDemoReply{
			Ok: false,
		}, err
	}
	return &pb.UpdateDemoReply{
		Ok: true,
	}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
	_, err := s.ac.DeleteArticle(ctx, biz.Article{
		ID: req.Id,
	})
	if err != nil {
		return &pb.DeleteDemoReply{
			Ok: false,
		}, err
	}
	return &pb.DeleteDemoReply{
		Ok: true,
	}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	_, err := s.ac.CreateArticle(ctx, biz.Article{})
	if err != nil {
		return &pb.GetDemoReply{
			Ok: false,
		}, err
	}
	return &pb.GetDemoReply{
		Ok: true,
	}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
	_, err := s.ac.CreateArticle(ctx, biz.Article{
		ID: req.Id,
	})
	if err != nil {
		return &pb.ListDemoReply{
			Ok: false,
		}, err
	}
	return &pb.ListDemoReply{
		Ok: true,
	}, nil
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
