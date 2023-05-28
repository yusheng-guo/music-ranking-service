package dao

import (
	"sync"

	"github.com/yushengguo557/music-ranking-service/db"
	"github.com/yushengguo557/music-ranking-service/global"
)

// Dao 数据结构 数据访问层
type Dao struct {
	*db.Database
	Lock sync.RWMutex
}

// NewDao 新建 Dao 实例化
func NewDao() *Dao {
	return &Dao{
		global.DB,
		sync.RWMutex{},
	}
}
