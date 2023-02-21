package scraper

import (
	"github.com/gocolly/colly"
)

type Selector struct {
	Container string
	Name      string
	Status    string
}

type Item struct {
	Service string
	Name    string
	Status  string
}

func ScrapeData(
	c *colly.Collector,
	selector Selector,
	service string,
	url string,
) []Item {
	var items []Item
	c.OnHTML(selector.Container, func(h *colly.HTMLElement) {
		item := Item{
			Service: service,
			Name:    h.ChildText(selector.Name),
			Status:  h.ChildText(selector.Status),
		}
		items = append(items, item)
	})
	c.Visit(url)
	return items
}
