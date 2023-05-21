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
			// colly.Async(true),
		),
	}
	// 限制 goroutine 数量 为 7
	// c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 7})
	return c
}

// Run 运行 实例化的 爬取器 Crawler
func (c *Crawler) Run() (err error) {
	// get urls
	var urls map[string]string
	// 酷狗
	// if urls, err = c.CrawlKugouRankUrls(); err != nil {
	// 	return fmt.Errorf("get kugou urls -- %w", err)
	// }
	// if err = c.CrawlKuGouMusic(urls); err != nil {
	// 	return fmt.Errorf("crawl kugou -- %w", err)
	// }

	// 咪咕
	if urls, err = c.CrawlMiguRankUrls(); err != nil {
		return fmt.Errorf("get migu urls -- %w", err)
	}
	if err = c.CrawlMiguMusic(urls); err != nil {
		return fmt.Errorf("crawl migu -- %w", err)
	}

	// QQ
	// if urls, err = c.CrawlQQRankUrls(); err != nil {
	// 	return fmt.Errorf("get qq urls -- %w", err)
	// }
	// if err = c.CrawlQQMusic(urls); err != nil {
	// 	return fmt.Errorf("crawl qq -- %w", err)
	// }

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})
	return nil
}

// Run 运行爬虫程序
func Run() (err error) {
	// 实例化 Crawler
	c := NewCrawler()

	// 运行
	if err = c.Run(); err != nil {
		panic(err)
	}

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
