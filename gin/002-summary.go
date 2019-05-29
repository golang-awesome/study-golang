package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/labstack/gommon/log"
	"io"
	"net/http"
	"os"
	"time"
)

type User struct {
	Username string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	Age      int    `form:"age" json:"age"` // age default to 0 if missing or malformed
}

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

	// render html
	g.LoadHTMLGlob("templates/*")
	g.GET("/upload.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "upload.html", gin.H{})
	})

	// body
	// model binding
	g.POST("/login", func(c *gin.Context) {
		var user User
		var err error
		contentType := c.Request.Header.Get("Content-Type")
		switch contentType {
		case "application/json":
			err = c.BindJSON(&user)
		case "application/x-www-form-urlencoded":
			err = c.MustBindWith(&user, binding.Form)
		}

		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}

		c.JSON(http.StatusOK, gin.H{
			"user":     user.Username,
			"password": user.Password,
			"age":      user.Age,
		})
	})

	// auto content-type binding
	// form data or json
	g.POST("/login2", func(c *gin.Context) {
		var user User
		err := c.Bind(&user)
		if err != nil {
			fmt.Println(err)
			log.Fatal(err)
		}
		c.JSON(http.StatusOK, gin.H{
			"username": user.Username,
			"password": user.Password,
			"age":      user.Age,
		})
	})

	// c.String
	// c.JSON
	// c.HTMl
	// c.XMl
	g.GET("/xml", func(c *gin.Context) {
		contentType := c.DefaultQuery("content_type", "json")
		if contentType == "json" {
			c.JSON(http.StatusOK, gin.H{
				"user":     "zhenglai",
				"password": "P@ssw0rd",
			})
		} else if contentType == "xml" {
			c.XML(http.StatusOK, gin.H{
				"user":     "zhenglai",
				"password": "P@ssw0rd",
			})
		}
	})

	// redirect
	g.GET("/redirect/google", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "https//www.google.com")
	})

	// router group
	// router group middlewares
	v1 := g.Group("/v1", MiddleWare())
	v1.GET("/login", func(c *gin.Context) {
		c.String(http.StatusOK, "v1 login")
	})

	v2 := g.Group("v2")
	// register middleware for specific route
	v2.GET("/login", MiddleWare(), func(c *gin.Context) {
		c.String(http.StatusOK, "v2 login")
	})

	g.Use(MiddleWare()) // take effect afterwards
	{ // convention only
		g.GET("/middleware", func(c *gin.Context) {
			request := c.MustGet("request").(string)
			req, _ := c.Get("request")
			c.JSON(http.StatusOK, gin.H{
				"middleware_request": request,
				"request":            req,
			})
		})
	}

	// middleware
	// - logging
	// - error handler
	// - auth & authorize
	g.GET("/auth/signin", func(c *gin.Context) {
		cookie := &http.Cookie{
			Name:     "session_id",
			Value:    "1234",
			Path:     "/",
			HttpOnly: true,
		}
		http.SetCookie(c.Writer, cookie)
		c.String(http.StatusOK, "Login successful")
	})

	g.GET("/home", AuthMiddleWare(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"data": "home",
		})
	})

	// async goroutine
	g.GET("/sync", func(c *gin.Context) {
		time.Sleep(5 * time.Second)
		log.Printf("Done in path: " + c.Request.URL.Path)
	})

	g.GET("/async", func(c *gin.Context) {
		cCp := c.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			log.Printf("Done in path: " + cCp.Request.URL.Path)
		}()
	})

	// _ = g.Run(":8080")

	// listen & serve on 0.0.0.0:8080
	// g.Run()

	// or with http as below
	// _ = http.ListenAndServe(":8080", g)
	s := &http.Server{
		Addr:           ":8080",
		Handler:        g,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	_ = s.ListenAndServe()

	// supervisor to manage restart & stop
}

func AuthMiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Request.Cookie("session_id"); err == nil {
			value := cookie.Value
			fmt.Println(value)
			if value == "1234" {
				c.Next()
				return
			}
		}
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Unauthorized",
		})
		c.Abort()
		return
	}
}

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		fmt.Println("before middleware")
		c.Set("request", "client_request")
		c.Next()
		fmt.Println("after middleware")
	}
}
