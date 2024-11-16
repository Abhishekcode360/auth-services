package jwtauthController

import (
	"auth-services/utils"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
	"sync"
)

var mu sync.Mutex

func AuthenticateJWT(c *gin.Context) {
	outputmap := make(map[string]interface{})
	authtoken := c.GetHeader("Authorization")
	authtoken = strings.TrimPrefix(authtoken, "Bearer ")
	if authtoken == "" {
		outputmap["status"] = http.StatusBadRequest
		outputmap["data"] = struct{}{}
		outputmap["message"] = "token missing"
	} else {
		mu.Lock()
		validation := utils.RevokedTokens[authtoken]
		mu.Unlock()
		if validation {
			outputmap["status"] = http.StatusUnauthorized
			outputmap["data"] = struct{}{}
			outputmap["message"] = "token revoked"
		} else {
			token, err := jwt.Parse(authtoken, func(token *jwt.Token) (interface{}, error) {
				return utils.JwtKey, nil
			})
			if err != nil || !token.Valid {
				outputmap["status"] = http.StatusBadRequest
				outputmap["data"] = struct{}{}
				outputmap["message"] = "token invalid"
			} else {
				outputmap["status"] = http.StatusOK
				outputmap["data"] = struct{}{}
				outputmap["message"] = "token is correct"
			}
		}
	}
	c.JSON(outputmap["status"].(int), outputmap)
}
