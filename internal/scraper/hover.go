package scraper

import (
	"fmt"

	"github.com/gocolly/colly"
)

type hoverSelector struct {
	status string
}

func GetHoverStatus(c *colly.Collector) {
	selector := hoverSelector{
		status: "#statusio_components .component:nth-child(2) .component-status",
	}
	c.OnHTML(selector.status, func(h *colly.HTMLElement) {
		fmt.Println(h.Text)
	})
	c.Visit("https://hoverstatus.com/")
}
