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
	"github.com/mhborthwick/awa-monitoring/internal/scraper"
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

	// Assign env variables
	token := os.Getenv("DOCKER_INFLUXDB_INIT_ADMIN_TOKEN")
	org := os.Getenv("DOCKER_INFLUXDB_INIT_ORG")
	bucket := os.Getenv("DOCKER_INFLUXDB_INIT_BUCKET")

	// Initialize colly - Klaviyo
	c := colly.NewCollector(
		colly.AllowedDomains(klaviyo),
	)

	// Initialize colly - Hover
	c1 := colly.NewCollector(
		colly.AllowedDomains(hover),
	)

	// Set selectors - Klaviyo
	selector := scraper.Selector{
		Container: ".components-container .component-inner-container",
		Name:      ".name",
		Status:    ".component-status",
	}

	// Set selectors - Hover
	selector1 := scraper.Selector{
		Container: "#statusio_components .component",
		Name:      ".component_name",
		Status:    ".component-status",
	}

	// Initialize items slice
	var items []scraper.Item

	// Scrape data - Klaviyo, Hover
	items = append(items, scraper.ScrapeData(c, selector, "Klaviyo", "https://status.klaviyo.com/")...)
	items = append(items, scraper.ScrapeData(c1, selector1, "Hover", "https://hoverstatus.com/")...)

	// Create data points - Klaviyo, Hover
	var dataPoints []DataPoint

	for _, i := range items {
		dataPoints = append(dataPoints, DataPoint{
			Measurement: "status",
			Tags:        map[string]string{"service": i.Service},
			Fields:      map[string]interface{}{i.Name: i.Status},
			Time:        time.Now(),
		})
	}

	// Validate dataPoints
	res, err := json.MarshalIndent(dataPoints, "", "  ")

	if err != nil {
		fmt.Println("Error marshaling Person to JSON:", err)
		return
	}

	fmt.Println(string(res))

	// Connect to InfluxDB
	client := influxdb2.NewClient("http://localhost:8086", token)
	writeAPI := client.WriteAPIBlocking(org, bucket)

	// Write to InfluxDB
	for _, dp := range dataPoints {
		p := influxdb2.NewPoint(dp.Measurement, dp.Tags, dp.Fields, dp.Time)
		err := writeAPI.WritePoint(context.Background(), p)
		if err != nil {
			fmt.Printf("Error writing point to InfluxDB: %v\n", err)
		}
	}

	// Close client
	client.Close()
}
