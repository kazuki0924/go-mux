package main

import (
	"fmt"
	"net/http"

	"github.com/kazuki0924/go-mux/controller"
	"github.com/kazuki0924/go-mux/infrastructure"
)

var (
	httpRouter     infrastructure.Router     = infrastructure.NewMuxRouter()
	postController controller.PostController = controller.NewPostController()
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
