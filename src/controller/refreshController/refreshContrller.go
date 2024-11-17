package refreshController

import (
	"auth-services/utils"
	"net/http"
	"strings"
	"sync"

	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
)

var mu sync.Mutex

func Refresh(c *gin.Context) {
	outputmap := make(map[string]interface{})
	authtoken := c.GetHeader("Authorization")
	if authtoken == "" {
		outputmap["status"] = http.StatusBadRequest
		outputmap["data"] = struct{}{}
		outputmap["message"] = "token missing"
	} else {
		authtoken = strings.TrimPrefix(authtoken, "Bearer ")
		token, err := jwt.Parse(authtoken, func(token *jwt.Token) (interface{}, error) {
			return utils.JwtKey, nil
		})
		if err != nil || !token.Valid {
			outputmap["status"] = http.StatusBadRequest
			outputmap["data"] = struct{}{}
			outputmap["message"] = "token invalid"
		} else {
			mu.Lock()
			_, tokenExist := utils.RevokedTokens[authtoken]
			mu.Unlock()
			if !tokenExist {
				outputmap["status"] = http.StatusUnauthorized
				outputmap["data"] = struct{}{}
				outputmap["message"] = "token not exist"
			} else {
				mu.Lock()
				utils.RevokedTokens[authtoken] = false
				mu.Unlock()
				outputmap["status"] = http.StatusOK
				outputmap["data"] = struct{}{}
				outputmap["message"] = "token is refreshed"

			}
		}
	}
	c.JSON(outputmap["status"].(int), outputmap)
}
