package crawler

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

// GetKuGouRankUrls Find all kugou rank links
func (c *Crawler) GetKugouRankUrls() error {
	c.OnHTML("div.pc_temp_side > div > ul > li > a", func(e *colly.HTMLElement) {
		// e.Request.Visit(e.Attr("href"))
		url := e.Attr("href")
		if strings.HasPrefix(url, "https://www.kugou.com/yy/rank/home/") {
			title := e.Attr("title")
			fmt.Println(title, url)
		}
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://www.kugou.com/yy/html/rank.html")
	return nil
}
