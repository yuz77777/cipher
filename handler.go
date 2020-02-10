package main

import (
	"github.com/gin-gonic/gin"
)

func indexGet(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}

func encryptionGet(c *gin.Context) {
	text := c.Query("text")
	text = encryption(text)
	m := map[string]string{"encrypted_text": text}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, m)
}

func decryptionGet(c *gin.Context) {
	text := c.Query("text")
	text = decryption(text)
	m := map[string]string{"decrypted_text": text}
	c.Header("Access-Control-Allow-Origin", "*")
	c.JSON(200, m)
}
