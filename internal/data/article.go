package data

import (
	"context"
	"github.com/go-kratos/kratos/v2/log"
	"github.com/issimo1/dsserver/internal/biz"
	"time"
)

type articleRepo struct {
	data *Data
	log  *log.Helper
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (u *articleRepo) CreateArticle(ctx context.Context, art *biz.Article) (*biz.Article, error) {
	_, err := u.data.db.Article.
		Create().
		SetID(art.ID).
		SetContent(art.Content).
		SetTitle(art.Title).
		Save(ctx)
	if err != nil {
		u.log.Errorf("create article %d error: %s", art.ID, err)
	}
	return art, nil
}

func (u *articleRepo) UpdateArticle(ctx context.Context, art *biz.Article) (*biz.Article, error) {
	err := u.data.db.Article.
		UpdateOneID(art.ID).
		SetUpdatedAt(time.Now()).
		SetTitle(art.Title).
		SetContent(art.Content).
		Exec(ctx)
	if err != nil {
		u.log.Errorf("update article %d error: %s", art.ID, err)
	}
	return art, err
}

func (u *articleRepo) DeleteArticle(ctx context.Context, id int64) (bool, error) {
	err := u.data.db.Article.DeleteOneID(id).Exec(ctx)
	if err != nil {
		u.log.Errorf("delete article %d error: %s", id, err)
		return false, err
	}
	return true, nil
}

func (u *articleRepo) GetArticle(ctx context.Context) ([]*biz.Article, error) {
	artcList := make([]*biz.Article, 0)
	a, err := u.data.db.Article.Get(ctx, 1)
	if err != nil {
		u.log.Errorf("get article error: %s", err)
		return nil, err
	}
	artcList = append(artcList, &biz.Article{
		ID:      a.ID,
		Content: a.Content,
		Title:   a.Title,
	})
	return artcList, nil
}

func (u *articleRepo) ListById(ctx context.Context, id int64) (*biz.Article, error) {
	art, err := u.data.db.Article.Get(ctx, id)
	if err != nil {
		u.log.Errorf("get article error: %s", err)
		return nil, err
	}
	return &biz.Article{
		ID:      art.ID,
		Content: art.Content,
		Title:   art.Title,
	}, nil
}
