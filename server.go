package main

import (
	"fmt"
	"net/http"

	"github.com/kazuki0924/go-mux/controller"
	router "github.com/kazuki0924/go-mux/infrastructure/router"
	"github.com/kazuki0924/go-mux/repository"
	"github.com/kazuki0924/go-mux/service"
)

var (
	httpRouter     router.Router             = router.NewMuxRouter()
	postRepository repository.PostRepository = repository.NewFirestoreRepository("posts")
	postService    service.PostService       = service.NewPostService(postRepository)
	postController controller.PostController = controller.NewPostController(postService)
)

func main() {
	// router := mux.NewRouter()
	const port string = ":8000"

	httpRouter.GET("/", func(resp http.ResponseWriter, req *http.Request) {
		fmt.Fprintln(resp, "up and running...")
	})
	httpRouter.GET("/posts", postController.GetPosts)
	httpRouter.POST("/post", postController.AddPost)
	// log.Println("server listening on port", port)
	// log.Fatalln(
	// 	http.ListenAndServe(port, router),
	// )
	httpRouter.SERVE(port)
}
