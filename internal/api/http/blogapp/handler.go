package blogapp

import (
	"context"
	"time"

	"github.com/anazcodes/blog-crud-api/internal/business/blogbus"
	errs "github.com/anazcodes/blog-crud-api/internal/errs/blogapperr"

	"github.com/anazcodes/blog-crud-api/pkg/request"
	"github.com/gofiber/fiber/v2"
)

type app struct {
	port     string
	business blogbus.Business
	fbr      *fiber.App
}

type App interface {
	Serve()

	// Shutdown stops listening and gracefully shutdowns the server.
	Shutdown(ctx context.Context) error
	Fiber() *fiber.App
}

func NewApp(port string, bus blogbus.Business) App {
	app := &app{
		port:     port,
		business: bus,
		fbr:      fiber.New(),
	}
	app.register(app.fbr)

	return app
}

func (a *app) Fiber() *fiber.App {
	return a.fbr
}

func (a *app) Shutdown(ctx context.Context) error {
	return a.fbr.ShutdownWithContext(ctx)
}

// @Summary		Add Blog Post
// @Description	Creates a new Blog Post entry to the system and returns the Blog Post's ID.
// @Tags			Blog Post
// @Accept			json
// @Produce		json
// @Param			body	body		AddBlogPost							true	"Payload"
// @Success		201		{object}	request.Response{data=BlogPostID}	"Success"
// @Failure		422		{object}	request.Response					"Failed to save, blog storage capacity reached"
// @Failure		404		{object}	request.Response					"Referenced resource does not found in the system"
// @Failure		400		{object}	request.Response					"Failed to bind JSON"
// @Failure		400		{object}	request.Response					"Failed to bind query"
// @Failure		400		{object}	request.Response					"Failed to bind path param"
// @Failure		500		{object}	request.Response					"Failed to process your request"
// @Router			/api/blog-post [post]
func (a *app) AddBlogPost(c *fiber.Ctx) error {
	body := new(AddBlogPost)
	return request.Handle(c, body, 15*time.Second, errs.Response,
		func(ctx context.Context, req any) (any, error) {
			body := req.(*AddBlogPost)
			c.Status(fiber.StatusCreated)
			return a.business.AddBlogPost(ctx, toBusAddBlogPost(body))
		})
}

// @Summary		Blog Posts
// @Description	Retrieves all the available Blog Posts.
// @Tags			Blog Post
// @Accept			json
// @Produce		json
// @Success		200	{object}	request.Response{data=[]BlogPost}	"Success"
// @Failure		422	{object}	request.Response					"Failed to save, blog storage capacity reached"
// @Failure		404	{object}	request.Response					"Referenced resource does not found in the system"
// @Failure		400	{object}	request.Response					"Failed to bind JSON"
// @Failure		400	{object}	request.Response					"Failed to bind query"
// @Failure		400	{object}	request.Response					"Failed to bind path param"
// @Failure		500	{object}	request.Response					"Failed to process your request"
// @Router			/api/blog-post [get]
func (a *app) BlogPosts(c *fiber.Ctx) error {
	return request.Handle(c, nil, 15*time.Second, errs.Response,
		func(ctx context.Context, req any) (any, error) {
			bps := a.business.BlogPosts(ctx)
			return toBlogPosts(bps), nil
		})
}

// @Summary		Blog Post
// @Description	Retrieves single Blog Post belongs to the provided ID.
// @Tags			Blog Post
// @Accept			json
// @Produce		json
// @Param			id	path		int								true	"Blog Post ID"
// @Success		200	{object}	request.Response{data=BlogPost}	"Success"
// @Failure		422	{object}	request.Response				"Failed to save, blog storage capacity reached"
// @Failure		404	{object}	request.Response				"Referenced resource does not found in the system"
// @Failure		400	{object}	request.Response				"Failed to bind JSON"
// @Failure		400	{object}	request.Response				"Failed to bind query"
// @Failure		400	{object}	request.Response				"Failed to bind path param"
// @Failure		500	{object}	request.Response				"Failed to process your request"
// @Router			/api/blog-post/{id} [get]
func (a *app) BlogPost(c *fiber.Ctx) error {
	body := new(BlogPostID)
	return request.Handle(c, body, 15*time.Second, errs.Response,
		func(ctx context.Context, req any) (any, error) {
			body := req.(*BlogPostID)
			bp, err := a.business.BlogPost(ctx, body.ID)
			return toBlogPost(bp), err
		})
}

// @Summary		Delete Blog Post
// @Description	Deletes Blog Post in the given ID.
// @Tags			Blog Post
// @Accept			json
// @Produce		json
// @Param			id	path		int									true	"Blog Post ID"
// @Success		200	{object}	request.Response{data=BlogPostID}	"Success"
// @Failure		422	{object}	request.Response					"Failed to save, blog storage capacity reached"
// @Failure		404	{object}	request.Response					"Referenced resource does not found in the system"
// @Failure		400	{object}	request.Response					"Failed to bind JSON"
// @Failure		400	{object}	request.Response					"Failed to bind query"
// @Failure		400	{object}	request.Response					"Failed to bind path param"
// @Failure		500	{object}	request.Response					"Failed to process your request"
// @Router			/api/blog-post/{id} [delete]
func (a *app) DeleteBlogPost(c *fiber.Ctx) error {
	body := new(BlogPostID)
	return request.Handle(c, body, 15*time.Second, errs.Response,
		func(ctx context.Context, req any) (any, error) {
			body := req.(*BlogPostID)
			return a.business.DeleteBlogPost(ctx, body.ID)
		})
}

// @Summary		Update Blog Post
// @Description	Updates Blog Post with the given data.
// @Tags			Blog Post
// @Accept			json
// @Produce		json
// @Param			id		path		int									true	"Blog Post ID"
// @Param			body	body		UpdateBlogPost						true	"Payload"
// @Success		200		{object}	request.Response{data=BlogPostID}	"Success"
// @Failure		422		{object}	request.Response					"Failed to save, blog storage capacity reached"
// @Failure		404		{object}	request.Response					"Referenced resource does not found in the system"
// @Failure		400		{object}	request.Response					"Failed to bind JSON"
// @Failure		400		{object}	request.Response					"Failed to bind query"
// @Failure		400		{object}	request.Response					"Failed to bind path param"
// @Failure		500		{object}	request.Response					"Failed to process your request"
// @Router			/api/blog-post/{id} [patch]
func (a *app) UpdateBlogPost(c *fiber.Ctx) error {
	body := new(UpdateBlogPost)
	return request.Handle(c, body, 15*time.Second, errs.Response,
		func(ctx context.Context, req any) (any, error) {
			body := req.(*UpdateBlogPost)
			return a.business.UpdateBlogPost(ctx, body.ID, toBusUpdateBlogPost(body))
		})
}
