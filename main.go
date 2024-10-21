package main

import (
	"github.com/labstack/echo"
	"log"
	"os"
	"scrapper/scrapper"
	"strings"
)

const fileName string = "jobs.csv"

func handleHome(c echo.Context) error {
	return c.File("home.html")
}

func handleScrape(c echo.Context) error {
	defer func() {
		err := os.Remove(fileName)
		if err != nil {
			log.Fatal(err)
		}
	}()

	cleanTerm := strings.ToLower(scrapper.CleanString(c.FormValue("term")))
	scrapper.Scrape(cleanTerm)
	return c.Attachment(fileName, fileName)
}

func main() {
	e := echo.New()
	e.GET("/", handleHome)
	e.POST("/scrape", handleScrape)
	e.Logger.Fatal(e.Start(":8080"))
}
