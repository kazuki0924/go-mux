package cache

import "github.com/kazuki0924/go-mux/entity"

type PostCache interface {
	Set(key string, value *entity.Post)
	Get(key string) *entity.Post
}
