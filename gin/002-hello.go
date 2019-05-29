package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"os"
)

var File = struct{}{}

func main() {
	g := gin.Default()
	g.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello World")
	})

	// route param
	g.GET("/user/:name", func(c *gin.Context) {
		name := c.Param("name")
		c.String(http.StatusOK, "Hello %s", name)
	})
	g.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		c.String(http.StatusOK, name+" is "+action)
	})

	// query string
	g.GET("/user", func(c *gin.Context) {
		fistname := c.DefaultQuery("firstname", "Guest")
		lastname := c.Query("lastname")
		c.String(http.StatusOK, "hello %s %s", fistname, lastname)
	})

	// form
	// application/x-www-form-urlencoded
	g.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")
		c.JSON(http.StatusOK, gin.H{
			"status": gin.H{
				"status_code": http.StatusOK,
				"status":      http.StatusText(http.StatusOK),
			},
			"message": message,
			"nick":    nick,
		})
	})

	// form & put
	g.PUT("/put", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")
		fmt.Printf("id: %s; page: %s; name: %s; message: %s \n", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"satus_code": http.StatusOK,
		})
	})

	// form
	// multipart/form-data
	// file uploading
	g.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		fmt.Println(name)
		file, header, err := c.Request.FormFile("file")
		if err != nil {
			c.String(http.StatusBadRequest, "Bad request: %v", err)
			return
		}

		filename := header.Filename
		out, err := os.Create(filename)
		if err != nil {
			log.Fatal(err)
		}
		defer out.Close()
		_, err = io.Copy(out, file)
		if err != nil {
			log.Fatal(err)
		}
		c.String(http.StatusOK, "upload successfully")
	})

	// multi files uploading
	g.POST("/upload/multi", func(c *gin.Context) {
		err := c.Request.ParseMultipartForm(200000)
		if err != nil {
			log.Fatal(err)
		}

		formdata := c.Request.MultipartForm
		files := formdata.File["file"]
		for i, _ := range files {
			fmt.Println("processing: " + files[i].Filename)
			file, err := files[i].Open()
			if err != nil {
				log.Fatal(err)
			}
			defer file.Close()

			out, err := os.Create(files[i].Filename)
			if err != nil {
				log.Fatal(err)
			}
			defer out.Close()

			_, err = io.Copy(out, file)
			if err != nil {
				log.Fatal(err)
			}
		}
		c.String(http.StatusCreated, "upload successfully")
	})

	// body
	_ = g.Run(":8080")
}
