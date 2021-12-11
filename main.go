package main

import (
	"net/http"

	"github.com/famasoon/gowhois/whois"
	"github.com/labstack/echo"
)

type whoisInfo struct {
	Domain      string `json:"domain"`
	WhoisResult string `json:"result"`
}

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "OK")
	})

	e.GET("/api/v1/whois/:domain", getWhoisResult)
	e.Logger.Fatal(e.Start(":8080"))
}

func getWhoisResult(c echo.Context) error {
	domain := c.Param("domain")
	result, err := whois.Whois(domain)
	if err != nil {
		return err
	}

	return c.JSON(http.StatusOK, whoisInfo{
		Domain:      domain,
		WhoisResult: result,
	})
}
