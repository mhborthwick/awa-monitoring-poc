package main

import (
	"fmt"

	"github.com/gocolly/colly"
)

const (
	klaviyo = "status.klaviyo.com"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains(klaviyo),
	)
	c.OnHTML(".page-status", func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText(".status"))
	})

	c.Visit("https://status.klaviyo.com/")
}
