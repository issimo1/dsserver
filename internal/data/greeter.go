package data

import (
	"context"
	"fmt"

	"github.com/issimo1/dsserver/internal/biz"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
)

type greeterRepo struct {
	data *Data
	log  *log.Helper
}

// NewGreeterRepo .
func NewGreeterRepo(data *Data, logger log.Logger) biz.GreeterRepo {
	return &greeterRepo{
		data: data,
		log:  log.NewHelper(logger),
	}
}

func (r *greeterRepo) Save(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	_, err := r.data.db.Article.
		Create().
		SetID(g.Id).
		SetContent(g.Hello).
		SetTitle(g.Hello + fmt.Sprintf("%d", g.Id)).
		Save(ctx)
	return g, err
}

func (r *greeterRepo) Update(ctx context.Context, g *biz.Greeter) (*biz.Greeter, error) {
	return g, nil
}

func (r *greeterRepo) FindByID(ctx context.Context, i int64) (*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListByHello(ctx context.Context, str string) ([]*biz.Greeter, error) {
	return nil, nil
}

func (r *greeterRepo) ListAll(ctx context.Context) ([]*biz.Greeter, error) {
	ps, err := r.data.db.Article.Query().All(ctx)
	if err != nil {
		return nil, err
	}
	rv := make([]*biz.Greeter, 0)
	for _, p := range ps {
		rv = append(rv, &biz.Greeter{
			Hello: fmt.Sprintf("%s 的内容为%s，id=%d", p.Title, p.Content, p.ID),
		})
	}
	return rv, nil
}

func (r *greeterRepo) GetHelloLike(ctx context.Context, str string) (int64, error) {
	save := r.data.redis.Get(ctx, str)
	rv, err := save.Int64()
	if err == redis.Nil {
		return 0, nil
	}
	return rv, err
}
