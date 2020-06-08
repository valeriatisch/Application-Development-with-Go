package main

import (
	"github.com/gin-gonic/gin"
	"newsfeeder/httpd/handler"
	"newsfeeder/platform/newsfeed"
)

func main() {

	feed := newsfeed.New()

	r := gin.Default()

	r.GET("/ping", handler.PingGet())
	r.GET("/newsfeed", handler.NewsfeedGet(feed))
	r.POST("/newsfeed", handler.NewsfeedPost(feed))

	_ = r.Run()

}
