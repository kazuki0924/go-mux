package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	"github.com/kazuki0924/go-mux/entity"
	"github.com/kazuki0924/go-mux/repository"
)

// type Post struct {
// 	Id    int    `json:"id"`
// 	Title string `json:"title"`
// 	Text  string `json:"text"`
// }

var (
	// posts []entity.Post
	repo repository.PostRepository = repository.NewPostRepository()
)

// func init() {
// 	posts = []entity.Post{{ID: 1, Title: "Title 1", Text: "Text 1"}}
// }

func GetPosts(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the posts"`))
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	var post entity.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"`))
		return
	}

	post.ID = rand.Int63()
	// posts = append(posts, post)
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	// result, err := json.Marshal(posts)
	// resp.Write(result)
	json.NewEncoder(resp).Encode(post)
}
