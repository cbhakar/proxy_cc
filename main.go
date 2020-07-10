package main

import (
	"fmt"
	"github.com/labstack/echo"
	"net/http"
	"time"
)
var (
	authEndpoint = "http://127.0.0.1:9991/auth?q=%s"
	userEndpoint = "http://127.0.0.1:9992/user/profile?q=%s"
	microserviceNameEndpoint = "http://127.0.0.1:9992/microservice/name"
)

func main() {
	e := echo.New()
	e.GET("/user/profile",userProfileRouter)
	e.GET("/microservice/name",microserviceNameRouter)

	e.Logger.Fatal(e.Start(":9999"))
}


func userProfileRouter(c echo.Context) error {
	u := c.Request().Header.Get("Username")
	if len(u) < 1 {
		return c.JSON(http.StatusBadRequest, map[string]string{"resp":"header 'Username' can't be nil"})
	}
	client := http.Client{
		Timeout:       time.Duration(5* time.Second),
	}
	resp, err := client.Get(fmt.Sprintf(authEndpoint,u))
	if err != nil{
		return c.JSON(http.StatusInternalServerError, map[string]string{"resp":"error authenticating user"})
	}
	if resp.StatusCode != 200{
		return c.JSON(http.StatusUnauthorized, map[string]string{"resp":"user not authorized"})
	}
	c.Redirect(http.StatusPermanentRedirect, fmt.Sprintf(userEndpoint, u))
	return nil
}


func microserviceNameRouter(c echo.Context) error {
	c.Redirect(http.StatusPermanentRedirect, microserviceNameEndpoint)
	return nil
}

