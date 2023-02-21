package scraper

import (
	"github.com/gocolly/colly"
)

const (
	klaviyo string = "status.klaviyo.com"
	hover   string = "hoverstatus.com"
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

func getSelectors(service string) Selector {
	var selectors Selector
	if service == "Klaviyo" {
		selectors = Selector{
			Container: ".components-container .component-inner-container",
			Name:      ".name",
			Status:    ".component-status",
		}
	}
	if service == "Hover" {
		selectors = Selector{
			Container: "#statusio_components .component",
			Name:      ".component_name",
			Status:    ".component-status",
		}
	}
	return selectors
}

func ScrapeData(
	service string,
	url string,
) []Item {
	c := colly.NewCollector(
		colly.AllowedDomains(klaviyo, hover),
	)
	selectors := getSelectors(service)
	var items []Item
	c.OnHTML(selectors.Container, func(h *colly.HTMLElement) {
		item := Item{
			Service: service,
			Name:    h.ChildText(selectors.Name),
			Status:  h.ChildText(selectors.Status),
		}
		items = append(items, item)
	})
	c.Visit(url)
	return items
}
