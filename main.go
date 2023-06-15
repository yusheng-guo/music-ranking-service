package main

import (
	"github.com/yushengguo557/music-ranking-service/crawler"
	"github.com/yushengguo557/music-ranking-service/global"
	"github.com/yushengguo557/music-ranking-service/router"
)

func init() {
	// 初始化 mysql 数据库
	if err := global.InitDB(); err != nil {
		panic(err)
	}
	// 初始化日志
	global.InitLog()
}

func main() {
	// 关闭数据库连接
	defer global.DB.Close()

	// // 定时爬虫
	// ticker := time.NewTicker(24 * time.Hour)
	// defer ticker.Stop()

	// go func() {
	// 	for range ticker.C {
	// 		go func() {
	// 			now := time.Now()
	// 			if now.Hour() == 0 && now.Minute() == 0 && now.Second() == 0 {
	// 				// 清空 songs 数据表
	// 				d := dao.NewDao()
	// 				_, err := d.Exec("TRUNCATE TABLE songs")
	// 				if err != nil {
	// 					panic(err.Error())
	// 				}

	// 				// 爬取数据
	// 				if err = crawler.Run(); err != nil {
	// 					panic(err)
	// 				}
	// 			}
	// 		}()
	// 	}
	// }()

	if err := crawler.Run(); err != nil {
		panic(err)
	}

	global.MRS_LOG.INFO("New Router....")
	// 路由
	r := router.NewRouter()
	if err := r.Run(); err != nil { // 监听并在 0.0.0.0:8080 上启动服务
		panic(err)
	}
}
