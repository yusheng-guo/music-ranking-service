package main

import (
	"time"

	"github.com/yushengguo557/music-ranking-service/crawler"
	"github.com/yushengguo557/music-ranking-service/dao"
	"github.com/yushengguo557/music-ranking-service/global"
	"github.com/yushengguo557/music-ranking-service/router"
)

func init() {
	// 初始化 mysql 数据库
	if err := global.InitDB(); err != nil {
		panic(err)
	}
}

// data := "data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAoAAAAKCAQAAAAnOwc2AAAAD0lEQVR4AWM4bYwJh7QgANLSYzmZH3ISAAAAAElFTkSuQmCC" // 这是一个示例数据，实际上应该从爬取结果中获取
// b64data := data[strings.IndexByte(data, ',')+1:]                                                                                 // 去掉"data:image/png;base64,"前缀
// img, err := base64.StdEncoding.DecodeString(b64data)                                                                             // 解码base64字符串
// if err != nil {
// 	panic(err)
// }
// err = ioutil.WriteFile("image.png", img, 0644) // 将图片保存为本地文件
// if err != nil {
// 	panic(err)
// }

func main() {
	// 关闭数据库连接
	defer global.DB.Close()

	// 定时爬虫
	ticker := time.NewTicker(24 * time.Hour)
	defer ticker.Stop()

	go func() {
		for range ticker.C {
			go func() {
				now := time.Now()
				if now.Hour() == 0 && now.Minute() == 0 && now.Second() == 0 {
					// 清空 songs 数据表
					d := dao.NewDao()
					_, err := d.Exec("TRUNCATE TABLE songs")
					if err != nil {
						panic(err.Error())
					}

					// 爬取数据
					if err = crawler.Run(); err != nil {
						panic(err)
					}
				}
			}()
		}
	}()

	// 路由
	r := router.NewRouter()
	if err := r.Run(); err != nil { // 监听并在 0.0.0.0:8080 上启动服务
		panic(err)
	}
}
