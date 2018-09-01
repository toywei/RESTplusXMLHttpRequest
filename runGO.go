package main

import (
	"./myKey"
	"fmt"
	"os"

	"github.com/aliyun/aliyun-sts-go-sdk/sts"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"net/http"
)

func handleError(err error) {
	fmt.Println(err)
	os.Exit(-1)
}

const (
	accessKeyID     = myKey.AccessKeyID
	accessKeySecret = myKey.AccessKeySecret
	roleArn         = myKey.RoleArn
	sessionName     = myKey.SessionName
)

type ExceptionResp struct {
	Status  string
	Message string
}

func main() {
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		// AllowOrigins: []string{"https://labstack.com", "https://labstack.net"},
		AllowOrigins: []string{"*"},
		// AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowHeaders: []string{"*"},
	}))
	e.GET("/aliSTS/:uid", func(c echo.Context) error {
		uid := c.Param("uid")
		fmt.Println("LOG ", uid)
		stsClient := sts.NewClient(accessKeyID, accessKeySecret, roleArn, sessionName)
		resp, err := stsClient.AssumeRole(3600)
		if err != nil {
			handleError(err)
			r := &ExceptionResp{
				Status:  "-1",
				Message: "what?",
			}
			return c.JSON(http.StatusOK, r)
		}
		r := resp.Credentials
		return c.JSON(http.StatusOK, r)
	})
	e.Logger.Fatal(e.Start(":1324"))
}
