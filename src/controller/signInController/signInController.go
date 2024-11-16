package signInController

import (
	"auth-services/utils"
	"github.com/gin-gonic/gin"
	jwt "github.com/golang-jwt/jwt/v5"
	"net/http"
	"sync"
	"time"
)

var mu sync.Mutex

func SignIn(c *gin.Context) {
	inputParams := utils.ExtractPrams(c)
	outputmap := make(map[string]interface{})
	if inputParams["email"] == "" || inputParams["password"] == "" {
		outputmap["status"] = http.StatusBadRequest
		outputmap["data"] = struct{}{}
		outputmap["message"] = "please provide the params"
	} else {
		mu.Lock()
		password, userExist := utils.Users[inputParams["email"]]
		mu.Unlock()
		if userExist {
			if password == inputParams["password"] {
				token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
					"useremail": inputParams["email"],
					"exp":       time.Now().Add(4 * time.Minute).Unix(),
				})
				tokenString, err := token.SignedString(utils.JwtKey)
				if err != nil {
					outputmap["status"] = http.StatusBadRequest
					outputmap["data"] = struct{}{}
					outputmap["message"] = "error with token generation"
				} else {
					outputmap["status"] = http.StatusOK
					outputmap["data"] = tokenString
					outputmap["message"] = "token generated"
				}

			} else {
				outputmap["status"] = http.StatusUnauthorized
				outputmap["data"] = struct{}{}
				outputmap["message"] = "password is incorrect"
			}
		} else {
			outputmap["status"] = http.StatusBadRequest
			outputmap["data"] = struct{}{}
			outputmap["message"] = "user not exist"
		}
	}
	c.JSON(outputmap["status"].(int), outputmap)
}
