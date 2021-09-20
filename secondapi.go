package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"log"
	"os"

	"github.com/gocolly/colly"
)

type product struct {
	Name        string `json:"name"`
	ImageUrl    string `json:"imageURL"`
	Description string `json:"description"`
	Price       string `json:"price"`
	Totalreview string `json:"totalReviews"`
}
type detail struct {
	Url     string  `json:"url"`
	Product product `json:"product"`
}

func main() {
	fname := "data.csv"
	file, err := os.Create(fname)
	if err != nil {
		log.Fatalf("couldnot create file")
		return
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	url := os.Args[1]
	prod := Firstapi(url)
	var d detail
	d.Url = url
	d.Product = prod

	writer.Write([]string{
		"url: " + d.Url + "\n" +
			"product:{" +
			"\n\t\tname: " + d.Product.Name +
			"\n\t\timageURL: " + d.Product.ImageUrl +
			"\n\t\tdescription: " + d.Product.Totalreview +
			"\n\t\tprice: " + d.Product.Price +
			"\n\t\ttotalReviews: " + d.Product.Description + "\n\t}",
	})

	b, err := json.Marshal(prod)
	if err != nil {
		return
	}
	fmt.Println(string(b))
}
func Firstapi(link string) product {

	c := colly.NewCollector(
		colly.AllowedDomains("amazon.com", "www.amazon.com"),
	)

	var p product
	c.OnHTML("#ppd", func(e *colly.HTMLElement) {

		p.Name = e.ChildText("#productTitle")
		p.Price = e.ChildText("#priceblock_ourprice")
		p.Totalreview = e.ChildText("#acrCustomerReviewText")
		p.ImageUrl = e.ChildAttr("img", "src")
	})
	c.OnHTML("#productDescription", func(e *colly.HTMLElement) {

		p.Description = e.ChildText("p")
	})
	c.Visit(link)
	return p
}
