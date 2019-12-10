package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("resources/templates/*.html")
	r.Static("/assets", "./resources/assets")
	r.GET("/", indexGet)
	r.GET("/api/encryption", encryptionGet)
	r.GET("/api/decryption", decryptionGet)
	r.Run()
}
