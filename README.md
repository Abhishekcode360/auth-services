Authentication services in GIN 

There are 5 micro-serivces in this authentication service which serve different purposes. 

To start the services : 
Redirect to /src folder that contains main.go.
Run the go run main.go. This will start the service.  

Now the service is available on the port 3001 like : 

Starting the server....
[GIN-debug] [WARNING] Running in "debug" mode. Switch to "release" mode in production.
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /signup                   --> auth-services/controller/signupController.SignUp (1 handlers)
[GIN-debug] POST   /signin                   --> auth-services/controller/signInController.SignIn (1 handlers)
[GIN-debug] GET    /tokenauth                --> auth-services/controller/jwtauthController.AuthenticateJWT (1 handlers)
[GIN-debug] POST   /revoketoken              --> auth-services/controller/revokeController.Revoke (1 handlers)
[GIN-debug] POST   /refresh                  --> auth-services/controller/refreshController.Refresh (1 handlers)
[GIN-debug] [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
Please check https://pkg.go.dev/github.com/gin-gonic/gin#readme-don-t-trust-all-proxies for details.
[GIN-debug] Listening and serving HTTP on :3001

How to use the services : 

User Creation : This service is used for user creation. 
  Checks : 
    If the user is already created.
    If the params are provided or not.
  
  URL - http://localhost:3001/signup
  Request Params - Pass these params in x-www-form-urlencoded 
    email:user@example.com
    password:password@123

	Sample Output : { "data": {}, "message": "user succesfully created", "status": 200 }	

User Signup : This service is used for JWT token generation. 
  Checks : 
    If the params are provided or not.
    If the user exists in the system.
    If the password of the user is correct or not.
    Error with token generation.
  
  URL - http://localhost:3001/signin
  Request Params - Pass these params in x-www-form-urlencoded 
    email:user@example.com
    password:password@123
	
	Sample output : 
	{ "data": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzE4MzI3MzgsInVzZXJlbWFpbCI6InVzZXJAZXhhbXBsZS5jb20ifQ.mSQPvAP6L42uzdbBWhAhzUzKmJUWI97PUHycKyIkc4w", "message": "token generated", "status": 200 }
	
JWT token authorization : This service is used for JWT token authentication.
	Checks: 
    If the token is provided or not.
    If the token is expired or not.

  URL - http://localhost:3001/tokenauth
  Request Params - Pass these params in Authorization.
    Bearer Token : ‘Pass the JWT token here.’
	
	Sample output : { "data": {}, "message": "token is correct", "status": 200 }

Revoke token service : This service is use to restrict the JWT token.
  Checks: 
    If the token is provided or not.
  
  URL - http://localhost:3001/revoketoken
  Request Params - Pass these params in Authorization.
    Bearer Token : ‘Pass the JWT token here.’
	
	Sample output : { "data": {}, "message": "token revoked", "status": 200 }

Refresh token :  This refresh is used to refresh the revoked JWT token until it expires.
  Checks: 
    If the token is provided or not.
    If the token is expired or not.
  
  URL - http://localhost:3001/refresh
  Request Params - Pass these params in Authorization.
    Bearer Token : ‘Pass the JWT token here.’
  	
	Sample output: { "data": {}, "message": "token is refreshed", "status": 200 }