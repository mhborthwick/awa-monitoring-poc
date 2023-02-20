package scraper

import (
	"github.com/gocolly/colly"
)

type HoverSelector struct {
	Container string
	Name      string
	Status    string
}

type HoverEntry struct {
	Name   string
	Status string
}

func GetHoverStatus(c *colly.Collector) []HoverEntry {

	var items []HoverEntry

	selector := HoverSelector{
		Container: "#statusio_components .component",
		Name:      ".component_name",
		Status:    ".component-status",
	}

	c.OnHTML(selector.Container, func(h *colly.HTMLElement) {
		item := HoverEntry{
			Name:   h.ChildText(selector.Name),
			Status: h.ChildText(selector.Status),
		}
		items = append(items, item)
	})

	c.Visit("https://hoverstatus.com/")

	return items
}
