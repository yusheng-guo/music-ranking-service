package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/yushengguo557/music-ranking-service/dao"
	"github.com/yushengguo557/music-ranking-service/model"
	"github.com/yushengguo557/music-ranking-service/model/response"
)

func GetQQHotSong(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("热歌榜", model.QQMusic)
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

func GetQQNewSong(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("新歌榜", model.QQMusic)
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

func GetQQSoaring(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("飙升榜", model.QQMusic)
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

func GetQQPopular(c *gin.Context) {
	s := dao.NewDao()
	s.Lock.RLock()
	songs, err := s.QueryMultiSongs("流行指数榜", model.QQMusic)
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
