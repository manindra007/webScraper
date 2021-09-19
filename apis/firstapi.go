package apis

import (
	"fmt"

	"github.com/gocolly/colly"
)

type product struct {
	name        string
	imageUrl    string
	description string
	price       string
	totalreview string
}

func Firstapi(link string) (product, error) {
	// fmt.Println("hello world")
	// fname := "data.csv"
	// file, err := os.Create(fname)
	// if err != nil {
	// 	log.Fatalf("couldnot create file")
	// 	return nil, err
	// }
	// defer file.Close()
	// writer := csv.NewWriter(file)
	// defer writer.Flush()
	c := colly.NewCollector(
		colly.AllowedDomains("amazon.com", "www.amazon.com"),
	)
	// c := colly.NewCollector{
	// 	colly.AllowedDomains("amazon.com", "www.amazon.com"),
	// }

	var p product
	c.OnHTML("#ppd", func(e *colly.HTMLElement) {
		// factID, err := strconv.Atoi(element.Attr("id"))
		// if err != nil {
		// 	fmt.Println("couldnt do this")
		// }
		// fmt.Println(factID)
		p.name = e.ChildText("#productTitle")
		p.price = e.ChildText("#priceblock_ourprice")
		p.totalreview = e.ChildText("#acrCustomerReviewText")
		p.imageUrl = e.ChildAttr("img", "src")

		// e.Request.Visit(e.Attr("h2"))
	})
	c.OnHTML("#productDescription", func(e *colly.HTMLElement) {

		p.description = e.ChildText("b")
	})
	c.Visit(link)
	// c.OnRequest(func(request *colly.Request) {
	// 	fmt.Println("visiting", request.URL.String())
	// })
	return p, nil
}
func main() {
	fmt.Println(Firstapi("https://www.amazon.com/HP-Micro-Edge-Flagship-i7-1160G7-Accessories/dp/B09D7HD3C9/ref=sr_1_3?dchild=1&keywords=laptop&qid=1632045838&sr=8-3"))
}
