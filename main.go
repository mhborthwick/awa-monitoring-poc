package main

import (
	"github.com/gocolly/colly"
	"github.com/mhborthwick/awa-monitoring/internal/scraper"
)

const (
	klaviyo = "status.klaviyo.com"
)

func main() {
	c := colly.NewCollector(
		colly.AllowedDomains(klaviyo),
	)
	scraper.GetKlaviyoStatus(c)
}
