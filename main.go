package main

import (
	"github.com/cerveraaxel/nginx_versions/scrapper"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/nginx-versions", scrapper.GetVersions)
	r.Run(":8080")
}
