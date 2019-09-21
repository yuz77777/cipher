package main

import (
	"cipher/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/templates/*.html")

	r.GET("/", handler.IndexHandler)
	r.POST("encrypted/", handler.EncryptedPostHandler)
	r.POST("decrypted/", handler.DecryptionPostHandler)

	r.Run()
}
