package main

import (
	"cipher/handler"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("views/templates/*.html")

	r.GET("/", handler.IndexHandler)
	r.POST("encrypted/", handler.EncryptionPostHandler)
	r.POST("decrypted/", handler.DecryptionPostHandler)

	r.GET("api/encryption/", handler.EncryptionApiHandler)
	r.GET("api/decryption/", handler.DecryptionApiHandler)

	r.Run()
}
