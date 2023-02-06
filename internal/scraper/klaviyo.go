package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type klaviyoSelector struct {
	pageStatus string
	status     string
}

func GetKlaviyoStatus(c *colly.Collector) {
	selector := klaviyoSelector{
		pageStatus: ".page-status",
		status:     ".status",
	}
	c.OnHTML(selector.pageStatus, func(h *colly.HTMLElement) {
		fmt.Println(h.ChildText(selector.status))
	})
	c.Visit("https://status.klaviyo.com/")
}
