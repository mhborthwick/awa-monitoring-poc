package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains("status.klaviyo.com"),
	)
	c.OnHTML(".page-status", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText(".status"))
	})

	c.Visit("https://status.klaviyo.com/")
}
