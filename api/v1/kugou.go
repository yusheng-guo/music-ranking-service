package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/music-ranking-service/dao"
	"github.com/yushengguo557/music-ranking-service/model"
	"github.com/yushengguo557/music-ranking-service/model/response"
)

func GetKugouSoaring(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("酷狗飙升榜", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if len(songs) > 0 {
		response.OKWithDataAndMessage(c, songs, "success")
		return
	}
	response.OKWithMessage(c, "len(songs)=0")
}

func GetKugouTop500(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("酷狗TOP500", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if len(songs) > 0 {
		response.OKWithDataAndMessage(c, songs, "success")
		return
	}
	response.OKWithMessage(c, "len(songs)=0")
}

func GetKugouPopular(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("蜂鸟流行音乐榜", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if len(songs) > 0 {
		response.OKWithDataAndMessage(c, songs, "success")
		return
	}
	response.OKWithMessage(c, "len(songs)=0")
}

func GetKugouDouyin(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("抖音热歌榜", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if len(songs) > 0 {
		response.OKWithDataAndMessage(c, songs, "success")
		return
	}
	response.OKWithMessage(c, "len(songs)=0")
}

func GetKugouKuaishou(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("快手热歌榜", model.KugouMusic)
	s.Lock.RUnlock()
	if err != nil {
		response.FailWithMessage(c, err.Error())
		return
	}
	if len(songs) > 0 {
		response.OKWithDataAndMessage(c, songs, "success")
		return
	}
	response.OKWithMessage(c, "len(songs)=0")
}
