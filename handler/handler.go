package handler

import (
	decryption "cipher/internal"
	encryption "cipher/internal"
	"github.com/gin-gonic/gin"
)

func IndexHandler(c *gin.Context)  {
	c.HTML(200, "index.html", gin.H{})
}

func EncryptedPostHandler(c *gin.Context)  {
	c.Request.ParseForm()
	encryptText := c.Request.Form["encrypt-text"]
	encryptedText := encryption.Encryption(encryptText[0])

	c.HTML(200, "encrypted.html", gin.H{
		"encryptedText" : encryptedText,
	})
}

func DecryptionPostHandler(c *gin.Context)  {
	c.Request.ParseForm()
	decryptText := c.Request.Form["decrypt-text"]
	decryptedText := decryption.Decryption(decryptText[0])

	c.HTML(200, "decrypted.html", gin.H{
		"decryptedText": decryptedText,
	})
}
