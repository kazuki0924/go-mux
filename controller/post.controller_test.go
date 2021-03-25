package controller

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kazuki0924/go-mux/entity"
	"github.com/kazuki0924/go-mux/repository"
	"github.com/kazuki0924/go-mux/service"
	"github.com/stretchr/testify/assert"
)

const (
	ID    int64  = 123
	TITLE string = "Title 11"
	TEXT  string = "Text 11"
)

var (
	postRepo       repository.PostRepository = repository.NewSQLiteRepository()
	postSrv        service.PostService       = service.NewPostService(postRepo)
	postController PostController            = NewPostController(postSrv)
)

func TestAddPost(t *testing.T) {
	//Create a new HTTP POST request
	var jsonReq = []byte(`{"title": "` + TITLE + `", "text": "` + TEXT + `"}`)
	req, _ := http.NewRequest("POST", "/posts", bytes.NewBuffer(jsonReq))

	//Assign HTTP Handler function (controller AddPost function)
	handler := http.HandlerFunc(postController.AddPost)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var post entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&post)

	// Assert HTTP response
	assert.NotNil(t, post.ID)
	assert.Equal(t, TITLE, post.Title)
	assert.Equal(t, TEXT, post.Text)

	// Clean up database
	cleanUp(&post)
}

func TestGetPosts(t *testing.T) {
	// Insert new post
	setup()

	// Create a GET HTTP request
	req, _ := http.NewRequest("GET", "/posts", nil)

	//Assign HTTP Handler function (controller GetPosts function)
	handler := http.HandlerFunc(postController.GetPosts)

	//Record HTTP Response (httptest)
	response := httptest.NewRecorder()

	//Dispatch the HTTP request
	handler.ServeHTTP(response, req)

	//Add Assertions on the HTTP Status code and the response
	status := response.Code

	if status != http.StatusOK {
		t.Errorf("Handler returned a wrong status code: got %v want %v", status, http.StatusOK)
	}

	// Decode the HTTP response
	var posts []entity.Post
	json.NewDecoder(io.Reader(response.Body)).Decode(&posts)

	// Assert HTTP response
	assert.NotNil(t, posts[0].ID)
	assert.Equal(t, TITLE, posts[0].Title)
	assert.Equal(t, TEXT, posts[0].Text)

	// Clean up database
	cleanUp(&posts[0])
}

func setup() {
	var post entity.Post = entity.Post{
		ID:    ID,
		Title: TITLE,
		Text:  TEXT,
	}
	postRepo.Save(&post)
}

func cleanUp(post *entity.Post) {
	postRepo.Delete(post)
}
