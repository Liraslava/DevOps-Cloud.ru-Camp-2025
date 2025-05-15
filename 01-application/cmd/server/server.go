package main

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
	"strings"

	"github.com/labstack/echo/v4"
)

type ViewData struct {
	Hostname string
	Ip       string
	Author   string
}

func main() {

	author := os.Getenv("$AUTHOR")
	_ = author
	e := echo.New()
	e.HideBanner = true

	hostname, err := os.Hostname()
	if err != nil {
		panic(err)
	}

	file, err := os.Open("templates/index.html")
	if err != nil {
		panic(err)
	}

	buff := strings.Builder{}

	_, err = io.Copy(&buff, file)
	if err != nil {
		panic(err)
	}
	t, err := template.New("index.html").Parse(buff.String())
	if err != nil {
		panic(err)
	}

	fmt.Println(hostname)

	e.GET("/", func(c echo.Context) error {
		data := ViewData{
			Ip:       c.RealIP(),
			Hostname: hostname,
			Author:   author,
		}
		buff := strings.Builder{}
		err := t.Execute(&buff, data)
		if err != nil {
			return err
		}
		return c.HTML(http.StatusOK, buff.String())
	})

	e.Logger.Fatal(e.Start(":8000"))

}
