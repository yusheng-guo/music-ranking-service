package main

import (
	"github.com/yushengguo557/music-ranking/crawler"
	"github.com/yushengguo557/music-ranking/global"
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

	// 爬虫
	if err := crawler.Run(); err != nil {
		panic(err)
	}
}
