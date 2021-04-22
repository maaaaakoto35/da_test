package main

import (
	"net/http"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()

	// router
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})
	e.GET("/v1/number2kanji/:param", func(c echo.Context) error { return num2kanjiHandler(c) })
	e.Logger.Fatal(e.Start(":1323"))
}

func num2kanjiHandler(c echo.Context) (err error) {
	param := c.Param("param")
	res, err := Kanji2number(param)
	if err != nil {
		c.JSON(http.StatusInternalServerError, struct {
			Message string
		}{
			Message: "変換できません.",
		})
		return
	}

	c.JSON(http.StatusOK, res)
	return
}
