package utils

import (
	"github.com/gin-gonic/gin"
)

func ExtractPrams(c *gin.Context) map[string]string {
	params := make(map[string]string)

	if c.Request.Method == "GET" {
		for key, value := range c.Request.URL.Query() {
			if len(value) > 0 {
				params[key] = value[0]
			} else {
				params[key] = ""
			}
		}
	}

	if c.Request.Method == "POST" {
		c.Request.ParseForm()
		for key, value := range c.Request.PostForm {
			if len(value) > 0 {
				params[key] = value[0]
			} else {
				params[key] = ""
			}
		}
	}
	return params
}
