package main

import (
	"fmt"

	"github.com/gocolly/colly"
	"github.com/mhborthwick/awa-monitoring/internal/scraper"
)

const (
	klaviyo = "status.klaviyo.com"
	hover   = "hoverstatus.com"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains(klaviyo, hover),
	)
	var entries = scraper.GetKlaviyoStatus(c)
	fmt.Println(entries)
}
