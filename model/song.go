package model

// Song 音乐结构体
type Song struct {
	ID       int    `db:"id"`       // id
	Name     string `db:"name"`     // 名称
	Singer   string `db:"singer"`   // 作者
	Duration string `db:"duration"` // 时长
	Link     string `db:"link"`     // 地址
	Cover    string `db:"cover"`    // 封面地址
}

// NewSong 创建一首歌曲
func NewSong(name, singer, duration, link, cover string) *Song {
	return &Song{
		Name:     name,
		Singer:   singer,
		Duration: duration,
		Link:     link,
		Cover:    cover,
	}
}
