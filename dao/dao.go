package dao

import (
	"sync"

	"github.com/yushengguo557/music-ranking/db"
	"github.com/yushengguo557/music-ranking/global"
)

// Dao 数据结构 数据访问层
type Dao struct {
	*db.Database
	Lock sync.RWMutex
}

// NewDao 新建 Dao 实例化
func NewDao() *Dao {
	d := &Dao{
		global.DB,
		sync.RWMutex{},
	}
	// 使用 colly 数据库
	_, err := global.DB.Exec("USE colly")
	if err != nil {
		panic(err.Error())
	}
	return d
}
