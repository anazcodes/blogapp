package cache

import (
	"context"
	"errors"
	"sync"
	"time"

	"github.com/anazcodes/blogapp/internal/business/blogbus"
)

var (
	ErrCacheInMaxCap = errors.New("cache capacity is full")
	ErrItemNotFound  = errors.New("item  not found")
)

func NewCache(capacity int) Cache {
	return &cache{
		blogs:    make(map[uint64]blogbus.BlogPost, capacity),
		RWMutex:  &sync.RWMutex{},
		capacity: capacity,
	}
}

type serial uint64

func (s *serial) Inc() {
	*s++
}

func (s *serial) ID() uint64 {
	return uint64(*s)
}

type cache struct {
	serial
	blogs    map[uint64]blogbus.BlogPost
	capacity int // Maximum number of items can store.
	*sync.RWMutex
}

type Cache interface {
	AddBlogPost(ctx context.Context, abp blogbus.AddBlogPost) (uint64, error)
	BlogPost(ctx context.Context, id uint64) (blogbus.BlogPost, error)
	BlogPosts(ctx context.Context) []blogbus.BlogPost
	DeleteBlogPost(ctx context.Context, id uint64) (uint64, error)
	UpdateBlogPost(ctx context.Context, id uint64, ubp blogbus.UpdateBlogPost) (uint64, error)
}

func (c *cache) AddBlogPost(ctx context.Context, abp blogbus.AddBlogPost) (uint64, error) {
	c.Lock()
	defer c.Unlock()

	if len(c.blogs) == c.capacity {
		return 0, ErrCacheInMaxCap
	}

	c.Inc()

	id := c.ID()
	now := time.Now()

	blog := blogbus.BlogPost{
		ID:          id,
		Title:       abp.Title,
		Description: abp.Description,
		Body:        abp.Body,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	c.blogs[id] = blog

	return id, nil
}

func (c *cache) BlogPost(ctx context.Context, id uint64) (blogbus.BlogPost, error) {
	c.RLock()
	defer c.RUnlock()

	bp, ok := c.blogs[id]
	if !ok {
		return blogbus.BlogPost{}, ErrItemNotFound
	}

	return bp, nil
}

func (c *cache) BlogPosts(ctx context.Context) []blogbus.BlogPost {
	c.RLock()
	defer c.RUnlock()

	bps := make([]blogbus.BlogPost, 0, len(c.blogs))

	for _, bp := range c.blogs {
		bps = append(bps, bp)
	}

	return bps
}

func (c *cache) DeleteBlogPost(ctx context.Context, id uint64) (uint64, error) {
	c.Lock()
	defer c.Unlock()

	_, ok := c.blogs[id]
	if !ok {
		return 0, ErrItemNotFound
	}

	delete(c.blogs, id)

	return id, nil
}

func (c *cache) UpdateBlogPost(ctx context.Context, id uint64, abp blogbus.UpdateBlogPost) (uint64, error) {
	c.Lock()
	defer c.Unlock()

	bp, ok := c.blogs[id]
	if !ok {
		return 0, ErrItemNotFound
	}

	if abp.Title != "" {
		bp.Title = abp.Title
	}
	if abp.Description != "" {
		bp.Description = abp.Description
	}
	if abp.Body != "" {
		bp.Body = abp.Body
	}

	bp.UpdatedAt = time.Now()
	c.blogs[id] = bp

	return id, nil
}
