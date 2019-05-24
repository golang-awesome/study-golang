package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	g := gin.New()

	middlewares := []gin.HandlerFunc{}

	g.Use(middlewares...)

	fmt.Println(http.ListenAndServe(":8080", g).Error())
}
