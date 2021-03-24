package controller

import (
	"encoding/json"
	"net/http"

	"github.com/kazuki0924/go-mux/entity"
	"github.com/kazuki0924/go-mux/errors"
	"github.com/kazuki0924/go-mux/service"
)

var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPosts(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

type controller struct{}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var post entity.Post
	bodyErr := json.NewDecoder(req.Body).Decode(&post)
	if bodyErr != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error creating the posts"})
		return
	}

	validationErr := postService.Validate(&post)
	if validationErr != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: validationErr.Error()})
		return
	}

	result, createErr := postService.Create(&post)
	if createErr != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}

	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
