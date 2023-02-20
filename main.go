package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/gocolly/colly"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"github.com/joho/godotenv"
)

const (
	klaviyo = "status.klaviyo.com"
	hover   = "hoverstatus.com"
)

type KlaviyoSelector struct {
	Container string
	Name      string
	Status    string
}

type HoverSelector struct {
	Container string
	Name      string
	Status    string
}

type Item struct {
	Service string
	Name    string
	Status  string
}

type DataPoint struct {
	Measurement string
	Tags        map[string]string
	Fields      map[string]interface{}
	Time        time.Time
}

func main() {
	// Load env variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	token := os.Getenv("DOCKER_INFLUXDB_INIT_ADMIN_TOKEN")
	org := os.Getenv("DOCKER_INFLUXDB_INIT_ORG")
	bucket := os.Getenv("DOCKER_INFLUXDB_INIT_BUCKET")

	// Initialize colly - klaviyo
	c := colly.NewCollector(
		colly.AllowedDomains(klaviyo),
	)

	// Initialize colly - hover
	c1 := colly.NewCollector(
		colly.AllowedDomains(hover),
	)

	// Initialize items slice
	var items []Item

	// Scrape Data - klaviyo
	selector := KlaviyoSelector{
		Container: ".components-container .component-inner-container",
		Name:      ".name",
		Status:    ".component-status",
	}

	c.OnHTML(selector.Container, func(h *colly.HTMLElement) {
		item := Item{
			Service: "Klaviyo",
			Name:    h.ChildText(selector.Name),
			Status:  h.ChildText(selector.Status),
		}
		items = append(items, item)
	})

	c.Visit("https://status.klaviyo.com/")

	// Scrape data - hover
	selector1 := HoverSelector{
		Container: "#statusio_components .component",
		Name:      ".component_name",
		Status:    ".component-status",
	}

	c1.OnHTML(selector1.Container, func(h *colly.HTMLElement) {
		item := Item{
			Service: "Hover",
			Name:    h.ChildText(selector1.Name),
			Status:  h.ChildText(selector1.Status),
		}
		items = append(items, item)
	})

	c1.Visit("https://hoverstatus.com/")

	// Create data points - klaviyo, hover
	var dataPoints []DataPoint

	for _, i := range items {
		dataPoints = append(dataPoints, DataPoint{
			Measurement: "status",
			Tags:        map[string]string{"service": i.Service},
			Fields:      map[string]interface{}{i.Name: i.Status},
			Time:        time.Now(),
		})
	}

	// Validate dataPoints looks correct
	m, err := json.MarshalIndent(dataPoints, "", "  ")

	if err != nil {
		fmt.Println("Error marshaling Person to JSON:", err)
		return
	}

	fmt.Println(string(m))

	// Connect to InfluxDB
	client := influxdb2.NewClient("http://localhost:8086", token)
	writeAPI := client.WriteAPIBlocking(org, bucket)

	for _, dp := range dataPoints {
		p := influxdb2.NewPoint(dp.Measurement, dp.Tags, dp.Fields, dp.Time)
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			fmt.Printf("Error writing point to InfluxDB: %v\n", err)
		}
	}

	fmt.Println("Data written to InfluxDB")

	// Close client
	client.Close()
}
