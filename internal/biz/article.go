package biz

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
)

type Article struct {
	ID      int64
	Content string
	Title   string
}

type ArticleRepo interface {
	CreateArticle(context.Context, *Article) (*Article, error)
	UpdateArticle(context.Context, *Article) (*Article, error)
	DeleteArticle(context.Context, int64) (bool, error)
	GetArticle(context.Context) ([]*Article, error)
	ListById(context.Context, int64) (*Article, error)
}

// ArticleUseCase case.
type ArticleUseCase struct {
	poll ArticleRepo
	log  *log.Helper
}

func NewArticleUsecase(poll ArticleRepo, logger log.Logger) *ArticleUseCase {
	return &ArticleUseCase{
		poll: poll,
		log:  log.NewHelper(logger),
	}
}

func (a *ArticleUseCase) CreateArticle(ctx context.Context, g Article) (*Article, error) {
	a.log.WithContext(ctx).Infof("CreateArticle: %v", g)
	return a.poll.CreateArticle(ctx, &g)
}

func (a *ArticleUseCase) UpdateArticle(ctx context.Context, g Article) (*Article, error) {
	a.log.WithContext(ctx).Infof("UpdateArticle: %v", g)
	return a.poll.UpdateArticle(ctx, &g)
}

func (a *ArticleUseCase) DeleteArticle(ctx context.Context, g Article) (bool, error) {
	a.log.WithContext(ctx).Infof("DeleteArticle: %v", g)
	return a.poll.DeleteArticle(ctx, g.ID)
}

func (a *ArticleUseCase) GetArticle(ctx context.Context, g Article) ([]*Article, error) {
	a.log.WithContext(ctx).Infof("GetArticle: %v", g)
	return a.poll.GetArticle(ctx)
}

func (a *ArticleUseCase) ListArticleById(ctx context.Context, g Article) (*Article, error) {
	a.log.WithContext(ctx).Infof("ListArticleById: %v", g)
	return a.poll.ListById(ctx, g.ID)
}
