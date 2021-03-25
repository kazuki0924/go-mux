package main

import (
	"os"

	"github.com/kazuki0924/go-mux/controller"
	router "github.com/kazuki0924/go-mux/infrastructure/router"
	"github.com/kazuki0924/go-mux/repository"
	"github.com/kazuki0924/go-mux/service"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postRepository repository.PostRepository = repository.NewSQLiteRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/post", postController.AddPost)

	httpRouter.SERVE(os.Getenv("PORT"))
}
