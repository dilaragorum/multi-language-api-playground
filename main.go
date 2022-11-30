package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
	"sync"
)

var Container = ProductContainer{
	mu:                 sync.Mutex{},
	LangMapForProducts: map[string]Products{},
}

func main() {
	e := echo.New()

	Container.FillProductDetailsMap()

	e.GET("/users", GetProducts)
	e.Logger.Fatal(e.Start(":3000"))
}

func GetProducts(c echo.Context) error {
	lang := c.Request().Header.Get("Accept-Language")
	return c.JSON(http.StatusOK, Container.LangMapForProducts[strings.ToLower(lang)])
}
