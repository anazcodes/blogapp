package blogapp

import (
	"log"

	_ "github.com/anazcodes/blogapp/docs"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/requestid"

	"github.com/gofiber/swagger"
)

//	@title			Blog Post CRUD
//	@version		1.0
//	@description	A Blog Post CRUD API in Go powered by Fiber framework.

//	@contact.name	Anaz Ibinu Rasheed
//	@contact.url	https://www.linkedin.com/in/anaz-ibinu-rasheed-a2b461253/
//	@contact.email	anazibinurasheed@gmail.com

// @license.name	Apache 2.0
// @license.url	http://www.apache.org/licenses/LICENSE-2.0.html
func (b *app) register(app *fiber.App) {
	app.Use(requestid.New())
	app.Use(logger.New(logger.Config{
		Format: "${ip}:${port} | ${locals:requestid} | ${status} | ${method} | ${path}​\n",
	}))

	app.Get("/swagger/*", swagger.HandlerDefault)

	router := app.Group("/api/blog-post")

	router.Post("", b.AddBlogPost)
	router.Get("", b.BlogPosts)
	router.Delete("/:id", b.DeleteBlogPost)
	router.Get("/:id", b.BlogPost)
	router.Patch("/:id", b.UpdateBlogPost)
}

func (b *app) Serve() {
	err := b.fbr.Listen(":" + b.port)
	if err != nil {
		log.Fatalln("Listen returned with an error: ", err)
	}

	log.Fatalln("Listening has been closed")
}
