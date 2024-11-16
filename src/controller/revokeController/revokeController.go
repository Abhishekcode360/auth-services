package revokeController

import (
	"auth-services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"sync"
)

var mu sync.Mutex

func Revoke(c *gin.Context) {
	authtoken := c.GetHeader("Authorization")
	outputmap := make(map[string]interface{})
	if authtoken != "" {
		authtoken = strings.TrimPrefix(authtoken, "Bearer ")
		mu.Lock()
		utils.RevokedTokens[authtoken] = true
		mu.Unlock()
		outputmap["status"] = http.StatusOK
		outputmap["data"] = struct{}{}
		outputmap["message"] = "token revoked"
	} else {
		outputmap["status"] = http.StatusUnauthorized
		outputmap["data"] = struct{}{}
		outputmap["message"] = "token missing"
	}
	c.JSON(outputmap["status"].(int), outputmap)
}
