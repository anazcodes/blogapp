package blogrepo

import (
	"context"
	"fmt"

	"github.com/anazcodes/blogapp/internal/business/blogbus"
	"github.com/anazcodes/blogapp/internal/repository/blogrepo/cache"
)

type repo struct {
	cache cache.Cache
}

func NewRepository(capacity int) *repo {
	return &repo{
		cache: cache.NewCache(capacity),
	}
}

func (r *repo) AddBlogPost(ctx context.Context, abp blogbus.AddBlogPost) (uint64, error) {
	id, err := r.cache.AddBlogPost(ctx, abp)
	if err != nil {
		return 0, fmt.Errorf("query: %w", err)
	}

	return id, nil
}

func (r *repo) BlogPost(ctx context.Context, id uint64) (blogbus.BlogPost, error) {
	bp, err := r.cache.BlogPost(ctx, id)
	if err != nil {
		return blogbus.BlogPost{}, fmt.Errorf("query: %w", err)
	}

	return bp, nil
}
func (r *repo) BlogPosts(ctx context.Context) []blogbus.BlogPost {
	bps := r.cache.BlogPosts(ctx)

	return bps
}

func (r *repo) DeleteBlogPost(ctx context.Context, id uint64) (uint64, error) {
	id, err := r.cache.DeleteBlogPost(ctx, id)
	if err != nil {
		return 0, fmt.Errorf("query: %w", err)
	}

	return id, nil
}

func (r *repo) UpdateBlogPost(ctx context.Context, id uint64, ubp blogbus.UpdateBlogPost) (uint64, error) {
	id, err := r.cache.UpdateBlogPost(ctx, id, ubp)
	if err != nil {
		return 0, fmt.Errorf("query: %w", err)
	}

	return id, nil
}
