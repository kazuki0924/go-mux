package main

import (
	"os"

	"github.com/kazuki0924/go-mux/cache"
	"github.com/kazuki0924/go-mux/controller"
	router "github.com/kazuki0924/go-mux/infrastructure/router"
	"github.com/kazuki0924/go-mux/repository"
	"github.com/kazuki0924/go-mux/service"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postRepository repository.PostRepository = repository.NewDynamoDBRepository()
	postService    service.PostService       = service.NewPostService(postRepository)
	postCache      cache.PostCache           = cache.NewRedisCache("localhost:6379", 0, 10)
	postController controller.PostController = controller.NewPostController(postService, postCache)
)

func main() {
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/post", postController.AddPost)
	httpRouter.GET("/posts/{id}", postController.GetPostByID)

	httpRouter.SERVE(os.Getenv("PORT"))
}
