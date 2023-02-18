package scraper

import (
	"github.com/gocolly/colly"
)

type klaviyoSelector struct {
	container string
	name      string
	status    string
}

type Entry struct {
	name   string
	status string
}

func GetKlaviyoStatus(c *colly.Collector) []Entry {
	var items []Entry

	selector := klaviyoSelector{
		container: ".components-container .component-inner-container",
		name:      ".name",
		status:    ".component-status",
	}

	c.OnHTML(selector.container, func(h *colly.HTMLElement) {
		item := Entry{
			name:   h.ChildText(selector.name),
			status: h.ChildText(selector.status),
		}
		items = append(items, item)
	})

	c.Visit("https://status.klaviyo.com/")

	return items
}
