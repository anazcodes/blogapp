package blogapp

import (
	"time"

	"github.com/anazcodes/blog-crud-api/internal/business/blogbus"
)

type BlogPost struct {
	ID          uint64    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	Body        string    `json:"body"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}

func toBlogPosts(bps []blogbus.BlogPost) []BlogPost {
	out := make([]BlogPost, len(bps))

	for i, bp := range bps {
		out[i] = toBlogPost(bp)
	}

	return out
}

func toBlogPost(bp blogbus.BlogPost) BlogPost {
	return BlogPost{
		ID:          bp.ID,
		Title:       bp.Title,
		Description: bp.Description,
		Body:        bp.Body,
		CreatedAt:   bp.CreatedAt,
		UpdatedAt:   bp.UpdatedAt,
	}
}

type AddBlogPost struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func toBusAddBlogPost(abp *AddBlogPost) blogbus.AddBlogPost {
	return blogbus.AddBlogPost{
		Title:       abp.Title,
		Description: abp.Description,
		Body:        abp.Body,
	}
}

type BlogPostID struct {
	ID uint64 `json:"id" uri:"id"`
}

type UpdateBlogPost struct {
	ID          uint64 `json:"-" uri:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Body        string `json:"body"`
}

func toBusUpdateBlogPost(ubp *UpdateBlogPost) blogbus.UpdateBlogPost {
	return blogbus.UpdateBlogPost{
		Title:       ubp.Title,
		Description: ubp.Description,
		Body:        ubp.Body,
	}
}
