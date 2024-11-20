# **Authentication service in GIN** 

There are 5 micro-serivces in this authentication service which serve different purposes. 

To start the auth service : 

**Method 1 : Docker**

1. Switch to '**master**' branch.
2. Redirect to /docker folder that contains dockercompose.yml.
3. Run the command - 'docker compose -f dockercompose.yml up' or 'sudo docker compose -f dockercompose.yml up'. 

**Method 2 : Main.go**

Switch to '**master**' branch.
Redirect to /src folder that contains main.go.
Run the '**go run main.go**'.

Both of these methods will start service on the port 3001.

**How to use the services :**

1. **_SignUp Service_** : This service is used for user creation.

Curl : 
```bash
curl --location 'http://localhost:3001/signup' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'email=auth@example.com' \
--data-urlencode 'password=password'
```

```
Sample Output : { "data": {}, "message": "user succesfully created", "status": 200 }	
```

TestCase :

1. If one of the params is not provided.
```
Sample Output : {"data":{},"message":"please provide the params","status":400}	
```

2. If the user is already created.
```
Sample Output : {"data":{},"message":"user already exist","status":400}	
```

2. **_SignIn Service_** : This service is used for JWT token generation. 
Curl :
```bash
curl --location 'http://localhost:3001/signin' \
--header 'Content-Type: application/x-www-form-urlencoded' \
--data-urlencode 'email=auth@example.com' \
--data-urlencode 'password=password'
```

```
Sample output : {"data":"eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzIxMzU1NTYsInVzZXJlbWFpbCI6ImF1dGhAZXhhbXBsZS5jb20ifQ.ailyMfXvfaT2vLKLpKql15kluxaPoogYEv8xrO5qTtk","message":"token generated","status":200}
```

Testcases : 

1. If one of the params not provided.
```
{"data":{},"message":"please provide the params","status":400}
```

2. If the password is incorrect.
```
{"data":{},"message":"password is incorrect","status":401}
```

3. **_JWT Authorization Service_** : This service is used for JWT token authentication.
Curl :
```bash
curl --location --request POST 'http://localhost:3001/tokenauth' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzIxMzYzNDQsInVzZXJlbWFpbCI6ImF1dGhAZXhhbXBsZS5jb20ifQ.10Nv_0mjYvDo1hJPapK2vGip26UyJvbdALzaV_SSsnw'
```
Provide the JWT token generated in signin service here.

```
Sample output : {"data":{},"message":"token is correct","status":200}
```

Testcases : 

1. If token is not provided.
```
{"data":{},"message":"token missing","status":400}
```

2. If the token is invalid or expired after 4 minutes.
```
{"data":{},"message":"token invalid","status":400}
```

3. If the token is revoked.
```
{"data":{},"message":"token revoked","status":401}
```

4. **_Revoke Token Service_** : This service is use to restrict the JWT token.
Curl :
```bash
curl --location --request POST 'http://localhost:3001/revoketoken' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzIxMzIzMjgsInVzZXJlbWFpbCI6ImF1dGhAZXhhbXBsZS5jb20ifQ.TgZO_Cj2bb8yRsU1iffQJdq3MOPjxbsQk-H3rnaMql4'
```
Provide the JWT token generated in signin service here.

```
Sample output : {"data":{},"message":"token revoked","status":200}
```

Testcases : 

1. If token is not provided.
```
{"data":{},"message":"token missing","status":401}
```

5. **_Refresh Token Service_** :  This refresh is used to refresh the revoked JWT token until it expires.
Curl :
```bash
curl --location --request POST 'http://localhost:3001/refresh' \
--header 'Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MzIxMzIzMjgsInVzZXJlbWFpbCI6ImF1dGhAZXhhbXBsZS5jb20ifQ.TgZO_Cj2bb8yRsU1iffQJdq3MOPjxbsQk-H3rnaMql4'
```
Provide the JWT token generated in signin service here.

```
Sample output : {"data":{},"message":"token revoked","status":200}
```

Testcases : 

1. If token is not provided.
```
{"data":{},"message":"token missing","status":401}
```

2. If the token is invalid then refresh will not happen.
```
{"data":{},"message":"token invalid","status":400}
```
