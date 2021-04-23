package main

import (
	"net/http"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func main() {
	e := echo.New()

	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// router
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/v1/number2kanji/:param", func(c echo.Context) error { return num2kanjiHandler(c) })
	e.Logger.Fatal(e.Start(":80"))
}

func num2kanjiHandler(c echo.Context) (err error) {
	param := c.Param("param")
	res, err := Kanji2number(param)
	if err != nil || res == "" {
		c.JSON(http.StatusNoContent, struct {
			Message string
		}{
			Message: "変換できません.",
		})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
