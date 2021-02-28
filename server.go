package main

import (
	"fmt"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	fmt.Printf("Starting service musicbrainz")
	initializeRoutes(e)
	e.Logger.Fatal(e.Start(":1337"))
}

func initializeRoutes(e *echo.Echo) {
	e.GET("/", Index)
	e.GET("/artists/:mbid", GetArtist)
}
