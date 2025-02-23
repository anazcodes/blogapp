package blogbus

import (
	"context"
	"fmt"
)

type business struct {
	repo Repo
}

type Repo interface {
	AddBlogPost(ctx context.Context, abp AddBlogPost) (uint64, error)
	BlogPost(ctx context.Context, id uint64) (BlogPost, error)
	BlogPosts(ctx context.Context) []BlogPost
	DeleteBlogPost(ctx context.Context, id uint64) (uint64, error)
	UpdateBlogPost(ctx context.Context, id uint64, ubp UpdateBlogPost) (uint64, error)
}

type Business interface {
	AddBlogPost(ctx context.Context, abp AddBlogPost) (ID, error)
	BlogPost(ctx context.Context, id uint64) (BlogPost, error)
	BlogPosts(ctx context.Context) []BlogPost
	DeleteBlogPost(ctx context.Context, id uint64) (ID, error)
	UpdateBlogPost(ctx context.Context, id uint64, ubp UpdateBlogPost) (ID, error)
}

func NewBusiness(repo Repo) Business {
	return &business{
		repo: repo,
	}
}

func (b *business) AddBlogPost(ctx context.Context, abp AddBlogPost) (ID, error) {
	id, err := b.repo.AddBlogPost(ctx, abp)
	if err != nil {
		return nil, fmt.Errorf("repo.addblogpost: %w", err)
	}

	return ToID(id), nil
}

func (b *business) BlogPosts(ctx context.Context) []BlogPost {
	v := b.repo.BlogPosts(ctx)
	return v
}

func (b *business) BlogPost(ctx context.Context, id uint64) (BlogPost, error) {
	bp, err := b.repo.BlogPost(ctx, id)
	if err != nil {
		return BlogPost{}, fmt.Errorf("repo.blogpost: %w id: %d", err, id)
	}

	return bp, nil
}

func (b *business) DeleteBlogPost(ctx context.Context, id uint64) (ID, error) {
	id, err := b.repo.DeleteBlogPost(ctx, id)
	if err != nil {
		return nil, fmt.Errorf("repo.deleteblogpost: %w id: %d", err, id)
	}

	return ToID(id), nil
}

func (b *business) UpdateBlogPost(ctx context.Context, id uint64, ubp UpdateBlogPost) (ID, error) {
	id, err := b.repo.UpdateBlogPost(ctx, id, ubp)
	if err != nil {
		return nil, fmt.Errorf("repo.updateblogpost: %w id: %d", err, id)
	}

	return ToID(id), nil
}
