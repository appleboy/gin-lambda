# gin-lambda

running golang using gin framework in AWS Lambda &amp; API Gateway

## Sample code

see the [main.go](./main.go) in master branch.

```go
package main

import (
	"log"
	"net/http"
	"os"

	"github.com/apex/gateway"
	"github.com/gin-gonic/gin"
)

func helloHandler(c *gin.Context) {
	name := c.Param("name")
	c.String(http.StatusOK, "Hello %s", name)
}

func welcomeHandler(c *gin.Context) {
	c.String(http.StatusOK, "Hello World from Go")
}

func rootHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"text": "Welcome to gin lambda server.",
	})
}

func routerEngine() *gin.Engine {
	// set server mode
	gin.SetMode(gin.DebugMode)

	r := gin.New()

	// Global middleware
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	r.GET("/welcome", welcomeHandler)
	r.GET("/user/:name", helloHandler)
	r.GET("/", rootHandler)

	return r
}

func main() {
	addr := ":" + os.Getenv("PORT")
	log.Fatal(gateway.ListenAndServe(addr, routerEngine()))
}
```

## Build and Upload

Build binary

```sh
$ GOOS=linux go build -o main .
$ zip deployment.zip main
```

Upload the `deployment.zip` to AWS Lambda using [drone-lambda](https://github.com/appleboy/drone-lambda) command.

```
$ AWS_ACCESS_KEY_ID=xxxx \ 
  AWS_SECRET_ACCESS_KEY=xxx \
  drone-lambda --region ap-southeast-1 \
  --function-name function_name \
  --zip-file deployment.zip
```

## AWS Policy

Add the following AWS policy if you want to integrate with CI/CD tools like Jenkins, GitLab Ci or Drone.

```json
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "s3:PutObject",
        "iam:ListRoles",
        "lambda:UpdateFunctionCode",
        "lambda:CreateFunction"
      ],
      "Resource": "arn:aws:logs:*:*:*"
    }
  ]
}
```
