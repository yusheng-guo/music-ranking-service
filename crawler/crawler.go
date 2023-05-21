package crawler

import (
	"fmt"

	"github.com/gocolly/colly"
)

// Crawler 爬行器 爬取网页
type Crawler struct {
	*colly.Collector
}

// NewCrawler 实例化 Crawler
func NewCrawler() *Crawler {
	// c := colly.NewCollector()
	// c.UserAgent =
	// 	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"
	// c.MaxDepth = 1
	// c.AllowedDomains = []string{"y.qq.com"}
	c := &Crawler{
		Collector: colly.NewCollector(
			colly.UserAgent("Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/113.0.0.0 Safari/537.36"),
			colly.MaxDepth(1),
			// colly.AllowedDomains("y.qq.com"),
			colly.Async(true),
		),
	}
	// 限制 goroutine 数量 为 7
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 7})
	return c
}

// Run 运行 实例化的 爬取器 Crawler
func (c *Crawler) Run() (err error) {
	// QQ Music
	// if err = c.CrawlerQQMusic(); err != nil {
	// 	return fmt.Errorf("%w", err)
	// }
	// if err = c.CrawlerNetEaseMusic(); err != nil {
	// 	return fmt.Errorf("%w", err)
	// }
	// if err = c.CrawlerKuGouMusic(); err != nil {
	// 	return fmt.Errorf("%w", err)
	// }
	if err = c.GetKugouRankUrls(); err != nil {
		return fmt.Errorf("%w", err)
	}
	return nil
}

// Run 运行爬虫程序
func Run() (err error) {
	// 实例化 Crawler
	c := NewCrawler()

	// 运行
	c.Run()

	// Wait until threads are finished
	c.Wait()
	return nil
}

// Find and visit all links
// c.OnHTML("a", func(e *colly.HTMLElement) {
// 	e.Request.Visit(e.Attr("href"))
// })

// c.OnHTML("a", func(e *colly.HTMLElement) {
// 	e.Request.Visit(e.Attr("href"))
// })

// c.OnRequest(func(r *colly.Request) {
// 	fmt.Println("Visiting", r.URL)
// })

// c.Visit("http://go-colly.org/")
