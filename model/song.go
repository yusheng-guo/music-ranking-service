package model

// Song 音乐结构体
type Song struct {
	Name     string // 名称
	Singer   string // 作者
	Duration string // 时长
	Link     string // 地址
}

// NewSong 创建一首歌曲
func NewSong(name, singer, duration, link string) *Song {
	return &Song{
		Name:     name,
		Singer:   singer,
		Duration: duration,
		Link:     link,
	}
}
