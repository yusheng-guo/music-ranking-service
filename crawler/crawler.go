package crawler

import (
	"fmt"
	"log"

	"github.com/gocolly/colly"
	"github.com/yushengguo557/music-ranking/dao"
	"github.com/yushengguo557/music-ranking/model"
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
	if err = c.CrawlerQQMusic(); err != nil {
		return fmt.Errorf("%w", err)
	}
	if err = c.CrawlerNetEaseMusic(); err != nil {
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

// CrawlerQQMusic 爬取 QQ 音乐
func (c *Crawler) CrawlerQQMusic() error {
	d := dao.NewDao()

	c.OnHTML(".songlist__item", func(e *colly.HTMLElement) {
		name := e.ChildText(".songlist__songname_txt")
		singer := e.ChildText(".songlist__artist a")
		duration := e.ChildText(".songlist__time")
		link_addr := e.ChildAttr("div.songlist__songname > span > a:nth-child(2)", "href")
		link := fmt.Sprintf("%s%s%s", "https://", e.Request.URL.Host, link_addr)
		s := model.NewSong(name, singer, duration, link, "")
		fmt.Println(s)
		d.Lock.Lock()
		if err := d.InsertSong(s); err != nil {
			// panic(err)
			log.Println(err)
		}
		d.Lock.Unlock()
	})

	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.Visit("https://y.qq.com/n/ryqq/toplist/4")
	c.Visit("https://y.qq.com/n/ryqq/toplist/62")
	c.Visit("https://y.qq.com/n/ryqq/toplist/26")
	c.Visit("https://y.qq.com/n/ryqq/toplist/27")
	c.Visit("https://y.qq.com/n/ryqq/toplist/52")
	c.Visit("https://y.qq.com/n/ryqq/toplist/67")
	return nil
}

func (c *Crawler) CrawlerNetEaseMusic() error {
	// d := dao.NewDao()

	// c.OnHTML(".songlist__item", func(e *colly.HTMLElement) {
	// 	name := e.ChildText(".songlist__songname_txt")
	// 	singer := e.ChildText(".songlist__artist a")
	// 	duration := e.ChildText(".songlist__time")
	// 	link_addr := e.ChildAttr("div.songlist__songname > span > a:nth-child(2)", "href")
	// 	link := fmt.Sprintf("%s%s%s", "https://", e.Request.URL.Host, link_addr)
	// 	s := model.NewSong(name, singer, duration, link, "")
	// 	fmt.Println(s)
	// 	if err := d.InsertSong(s); err != nil {
	// 		// panic(err)
	// 		log.Println(err)
	// 	}
	// })

	// c.OnRequest(func(r *colly.Request) {
	// 	fmt.Println("Visiting", r.URL)
	// })
	return nil
}
