package signupController

import (
	"auth-services/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var mu sync.Mutex

func SignUp(c *gin.Context) {
	var userExist bool
	inputParams := utils.ExtractPrams(c)
	outputmap := make(map[string]interface{})
	if inputParams["email"] == "" || inputParams["password"] == "" {
		outputmap["status"] = http.StatusBadRequest
		outputmap["data"] = struct{}{}
		outputmap["message"] = "please provide the params"
	} else {
		mu.Lock()
		_, userExist = utils.Users[inputParams["email"]]
		mu.Unlock()
		if userExist {
			outputmap["status"] = http.StatusBadRequest
			outputmap["data"] = struct{}{}
			outputmap["message"] = "user already exist"
		} else {
			utils.Users[inputParams["email"]] = inputParams["password"]
			outputmap["status"] = http.StatusOK
			outputmap["data"] = struct{}{}
			outputmap["message"] = "user succesfully created"
		}
	}
	c.JSON(outputmap["status"].(int), outputmap)
}
