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
	// get urls
	var urls map[string]string
	// 酷狗
	// if urls, err = c.CrawlKugouRankUrls(); err != nil {
	// 	return fmt.Errorf("get kugou urls -- %w", err)
	// }
	// urls = map[string]string{
	// 	"酷狗飙升榜":    "https://www.kugou.com/yy/rank/home/1-6666.html",
	// 	"酷狗TOP500": "https://www.kugou.com/yy/rank/home/1-8888.html",
	// 	"蜂鸟流行音乐榜":  "https://www.kugou.com/yy/rank/home/1-59703.html",
	// 	"抖音热歌榜":    "https://www.kugou.com/yy/rank/home/1-52144.html",
	// 	"快手热歌榜":    "https://www.kugou.com/yy/rank/home/1-52767.html",
	// }
	// if err = c.CrawlKuGouMusic(urls); err != nil {
	// 	return fmt.Errorf("crawl kugou -- %w", err)
	// }

	// 咪咕
	// if urls, err = c.CrawlMiguRankUrls(); err != nil {
	// 	return fmt.Errorf("get migu urls -- %w", err)
	// }

	// urls = map[string]string{
	// 	"尖叫新歌榜": "https://music.migu.cn/v3/music/top/jianjiao_newsong",
	// 	"尖叫热歌榜": "https://music.migu.cn/v3/music/top/jianjiao_hotsong",
	// 	"尖叫原创榜": "https://music.migu.cn/v3/music/top/jianjiao_original",
	// }
	// if err = c.CrawlMiguMusic(urls); err != nil {
	// 	return fmt.Errorf("crawl migu -- %w", err)
	// }

	// QQ
	// if urls, err = c.CrawlQQRankUrls(); err != nil {
	// 	return fmt.Errorf("get qq urls -- %w", err)
	// }
	urls = map[string]string{
		"飙升榜":   "https://y.qq.com/n/ryqq/toplist/62",
		"热歌榜":   "https://y.qq.com/n/ryqq/toplist/26",
		"新歌榜":   "https://y.qq.com/n/ryqq/toplist/27",
		"流行指数榜": "https://y.qq.com/n/ryqq/toplist/4",
	}
	if err = c.CrawlQQMusic(urls); err != nil {
		return fmt.Errorf("crawl qq -- %w", err)
	}

	// 网易云
	// if urls, err = c.CrawlNeteaseRankUrls(); err != nil {
	// 	return fmt.Errorf("get netease urls -- %w", err)
	// }
	// urls = map[string]string{
	// 	"飙升榜": "https://music.163.com/#/discover/toplist?id=19723756",
	// 	"新歌榜": "https://music.163.com/#/discover/toplist?id=3779629",
	// 	"原创榜": "https://music.163.com/#/discover/toplist?id=2884035",
	// 	"热歌榜": "https://music.163.com/#/discover/toplist?id=3778678",
	// }
	// if err = c.CrawlNetEaseMusic(urls); err != nil {
	// 	return fmt.Errorf("crawl netease -- %w", err)
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
