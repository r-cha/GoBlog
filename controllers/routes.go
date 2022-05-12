package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type routes struct {
	router *gin.Engine
}

func NewRoutes() routes {
	r := routes{
		router: gin.Default(),
	}

	v1 := r.router.Group("/v1")

	r.addPostsController(v1)

	return r
}

func (r routes) Run(addr ...string) error {
	return r.router.Run()
}

func (r routes) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	r.router.ServeHTTP(w, req)
}
