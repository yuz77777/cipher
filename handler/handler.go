package handler

import (
	"github.com/gin-gonic/gin"
	"strings"
)

func IndexHandler(c *gin.Context)  {
	c.HTML(200, "index.html", gin.H{})
}

func EncryptedPostHandler(c *gin.Context)  {
	c.Request.ParseForm()
	encryptText := c.Request.Form["encrypt-text"]
	encryptedText := encryption(encryptText[0])

	c.HTML(200, "encrypted.html", gin.H{
		"encryptedText" : encryptedText,
	})
}

func DecryptionPostHandler(c *gin.Context)  {
	c.Request.ParseForm()
	decryptText := c.Request.Form["decrypt-text"]
	decryptedText := decryption(decryptText[0])

	c.HTML(200, "decrypted.html", gin.H{
		"decryptedText": decryptedText,
	})
}

func encryption(s string) string {
	sl := strings.Split(s, "")
	len := len(sl)
	var cipher string

	for i := 0; i < len; i++ {
		switch sl[i] {
		case "a":
			sl[i] = "å"
		case "b":
			sl[i] = "∫"
		case "c":
			sl[i] = "ç"
		case "d":
			sl[i] = "∂"
		case "e":
			sl[i] = "´"
		case "f":
			sl[i] = "ƒ"
		case "g":
			sl[i] = "©"
		case "h":
			sl[i] = "˙"
		case "i":
			sl[i] = "ˆ"
		case "j":
			sl[i] = "∆"
		case "k":
			sl[i] = "˚"
		case "l":
			sl[i] = "¬"
		case "m":
			sl[i] = "µ"
		case "n":
			sl[i] = "Ø"
		case "o":
			sl[i] = "ø"
		case "p":
			sl[i] = "π"
		case "q":
			sl[i] = "œ"
		case "r":
			sl[i] = "®"
		case "s":
			sl[i] = "ß"
		case "t":
			sl[i] = "†"
		case "u":
			sl[i] = "¨"
		case "v":
			sl[i] = "√"
		case "w":
			sl[i] = "∑"
		case "x":
			sl[i] = "≈"
		case "y":
			sl[i] = `\`
		case "z":
			sl[i] = "Ω"
		case "1":
			sl[i] = "¡"
		case "2":
			sl[i] = "™"
		case "3":
			sl[i] = "£"
		case "4":
			sl[i] = "¢"
		case "5":
			sl[i] = "∞"
		case "6":
			sl[i] = "§"
		case "7":
			sl[i] = "¶"
		case "8":
			sl[i] = "•"
		case "9":
			sl[i] = "ª"
		case "0":
			sl[i] = "º"
		case "@":
			sl[i] = "€"
		case "K":
			sl[i] = ""
		case "#":
			sl[i] = "‹"
		case "%":
			sl[i] = "ﬁ"
		case "&":
			sl[i] = "‡"
		case "[":
			sl[i] = "”"
		case "]":
			sl[i] = "’"
		case "(":
			sl[i] = "·"
		case ")":
			sl[i] = "‚"
		case "?":
			sl[i] = "¿"
		case "!":
			sl[i] = "⁄"
		case "-":
			sl[i] = "—"
		case "+":
			sl[i] = "±"
		case "=":
			sl[i] = "≠"
		case "~":
			sl[i] = "`"
		case ",":
			sl[i] = "≤"
		case ".":
			sl[i] = "≥"
		}
		cipher += sl[i]
	}
	
	return cipher
}

func decryption(s string) string {
	sl := strings.Split(s, "")
	length := len(sl)
	var cipher string

	for i := 0; i < length; i++ {
		switch sl[i] {
		case "å":
			sl[i] = "a"
		case "∫":
			sl[i] = "b"
		case "ç":
			sl[i] = "c"
		case "∂":
			sl[i] = "d"
		case "´":
			sl[i] = "e"
		case "ƒ":
			sl[i] = "f"
		case "©":
			sl[i] = "g"
		case "˙":
			sl[i] = "h"
		case "ˆ":
			sl[i] = "i"
		case "∆":
			sl[i] = "j"
		case "˚":
			sl[i] = "k"
		case "¬":
			sl[i] = "l"
		case "µ":
			sl[i] = "m"
		case "ø":
			sl[i] = "o"
		case "π":
			sl[i] = "p"
		case "œ":
			sl[i] = "q"
		case "®":
			sl[i] = "r"
		case "ß":
			sl[i] = "s"
		case "†":
			sl[i] = "t"
		case "¨":
			sl[i] = "u"
		case "√":
			sl[i] = "v"
		case "∑":
			sl[i] = "w"
		case "≈":
			sl[i] = "x"
		case `\`:
			sl[i] = `y`
		case "Ω":
			sl[i] = "z"
		case "¡":
			sl[i] = "1"
		case "™":
			sl[i] = "2"
		case "£":
			sl[i] = "3"
		case "¢":
			sl[i] = "4"
		case "∞":
			sl[i] = "5"
		case "§":
			sl[i] = "6"
		case "¶":
			sl[i] = "7"
		case "•":
			sl[i] = "8"
		case "ª":
			sl[i] = "9"
		case "º":
			sl[i] = "0"
		case "€":
			sl[i] = "@"
		case "":
			sl[i] = "K"
		case "‹":
			sl[i] = "#"
		case "ﬁ":
			sl[i] = "%"
		case "‡":
			sl[i] = "&"
		case "”":
			sl[i] = "["
		case "’":
			sl[i] = "]"
		case "·":
			sl[i] = "("
		case "‚":
			sl[i] = ")"
		case "¿":
			sl[i] = "?"
		case "⁄":
			sl[i] = "!"
		case "—":
			sl[i] = "-"
		case "±":
			sl[i] = "+"
		case "≠":
			sl[i] = "="
		case "`":
			sl[i] = "~"
		case "≤":
			sl[i] = ","
		case "≥":
			sl[i] = "."
		case " ":
			sl[i] = "\n"
		case "Ø":
			sl[i] = "n"
		}
		cipher += sl[i]
	}

	return cipher
}
