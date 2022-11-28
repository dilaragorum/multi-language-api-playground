package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"strings"
)

var (
	TRProducts []Product
	ENProducts []Product
)

func init() {
	TRProducts = append(TRProducts, Product{
		Id:          1,
		Description: "Kırmızı Yırtmaçlı Saten Elbise",
		Color:       "Kırmızı",
		Category:    "Elbise",
	})
	ENProducts = append(ENProducts, Product{
		Id:          1,
		Description: "Red Slit Satin Dress",
		Color:       "Red",
		Category:    "Dress",
	})
}

func main() {
	e := echo.New()

	e.GET("/users", GetProducts)
	e.Logger.Fatal(e.Start(":3000"))
}

func GetProducts(c echo.Context) error {
	lang := c.Request().Header.Get("Accept-Language")
	if strings.ToLower(lang) == "tr" {
		return c.JSON(http.StatusOK, TRProducts)
	}

	return c.JSON(http.StatusOK, ENProducts)
}
