package infrastructure

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/kazuki0924/go-mux/infrastructure/middleware"
)

type muxRouter struct {
}

var (
	muxDispatcher = mux.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) GET(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("GET")
}

func (*muxRouter) POST(uri string, f func(w http.ResponseWriter, req *http.Request)) {
	muxDispatcher.HandleFunc(uri, f).Methods("POST")
}

func (*muxRouter) SERVE(port string) {
	fmt.Printf("Mux HTTP server running on port %v", port)
	muxDispatcher.Use(middleware.MuxCORS)
	http.ListenAndServe(":"+port, muxDispatcher)
}
