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

func NewDemoService(uc *biz.GreeterUsecase, ac *biz.ArticleUseCase) *DemoService {
	return &DemoService{
		uc: uc,
		ac: ac,
	}
}

func (s *DemoService) CreateDemo(ctx context.Context, req *pb.CreateDemoRequest) (*pb.CreateDemoReply, error) {
	cdr := biz.Article{
		ID:      req.Id,
		Title:   req.Comment,
		Content: req.Content,
	}
	_, err := s.ac.CreateArticle(ctx, cdr)
	if err != nil {
		return nil, err
	}
	return &pb.CreateDemoReply{
		Results: []*pb.CreateDemoReply_Article{
			{
				Id:      req.Id,
				Comment: req.Title,
				Content: req.Content,
			}},
	}, nil
}
func (s *DemoService) UpdateDemo(ctx context.Context, req *pb.UpdateDemoRequest) (*pb.UpdateDemoReply, error) {
	_, err := s.ac.UpdateArticle(ctx, biz.Article{
		ID:      req.Id,
		Title:   req.Comment,
		Content: req.Content,
	})
	if err != nil {
		return nil, err
	}
	return &pb.UpdateDemoReply{
		Results: []*pb.UpdateDemoReply_Article{
			{
				Id:      req.Id,
				Comment: req.Comment,
				Content: req.Content,
			}},
	}, nil
}
func (s *DemoService) DeleteDemo(ctx context.Context, req *pb.DeleteDemoRequest) (*pb.DeleteDemoReply, error) {
	_, err := s.ac.DeleteArticle(ctx, biz.Article{
		ID: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.DeleteDemoReply{
		Results: []*pb.DeleteDemoReply_Article{
			{
				Id:      req.Id,
				Comment: "null",
				Content: "null",
			}},
	}, nil
}
func (s *DemoService) GetDemo(ctx context.Context, req *pb.GetDemoRequest) (*pb.GetDemoReply, error) {
	_, err := s.ac.GetArticle(ctx, biz.Article{})
	if err != nil {
		return nil, err
	}
	return &pb.GetDemoReply{
		Results: []*pb.GetDemoReply_Article{
			{
				Id:      -1,
				Comment: "null",
				Content: "null",
			}},
	}, nil
}
func (s *DemoService) ListDemo(ctx context.Context, req *pb.ListDemoRequest) (*pb.ListDemoReply, error) {
	_, err := s.ac.ListArticleById(ctx, biz.Article{
		ID: req.Id,
	})
	if err != nil {
		return nil, err
	}
	return &pb.ListDemoReply{
		Results: []*pb.ListDemoReply_Article{
			{
				Id:      req.Id,
				Comment: "null",
				Content: "null",
			}},
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
