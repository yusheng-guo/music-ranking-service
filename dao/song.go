package dao

import (
	"database/sql"
	"fmt"

	"github.com/yushengguo557/music-ranking-service/global"
	"github.com/yushengguo557/music-ranking-service/model"
)

// Insert 插入数据
func (d *Dao) InsertSong(s *model.Song) (err error) {
	// 使用 colly 数据库
	_, err = global.DB.Exec("USE colly")
	if err != nil {
		// panic(err.Error())
		panic(fmt.Errorf("new dao, err: %w", err))
	}
	// SQL语句
	sqlStr := "INSERT INTO songs (name, singer, duration, link, cover, tag, platform) VALUES (?, ?, ?, ?, ?, ?, ?)"
	// 执行SQL语句
	_, err = d.DB.Exec(sqlStr, s.Name, s.Singer, s.Duration, s.Link, s.Cover, s.Tag, s.Platform)
	if err != nil {
		return fmt.Errorf("insert song, err: %w", err)
	}
	// 获取插入数据的id
	// id, err := ret.LastInsertId()
	// if err != nil {
	// 	return fmt.Errorf("get id failed, err:%w", err)
	// }
	// fmt.Println("insert success, id:", id)
	return nil
}

// QueryOne 查询单条数据
func (d *Dao) QueryOneSong(db *sql.DB) (err error) {
	// 使用 colly 数据库
	_, err = global.DB.Exec("USE colly")
	if err != nil {
		// panic(err.Error())
		panic(fmt.Errorf("new dao, err: %w", err))
	}
	//SQL语句
	sqlStr := "select name, singer, duration, link, cover, tag, platform from songs where id=?"
	var s model.Song
	err = d.DB.QueryRow(sqlStr, 1).Scan(&s.Name, &s.Singer, &s.Duration, &s.Link, &s.Cover, &s.Tag, &s.Platform)
	if err != nil {
		return fmt.Errorf("query failed, err:%w", err)
	}
	fmt.Println("query success:", s)
	return nil
}

// 查询多条数据
func (d *Dao) QueryAllSongs() ([]model.Song, error) {
	// 使用 colly 数据库
	_, err := global.DB.Exec("USE colly")
	if err != nil {
		// panic(err.Error())
		panic(fmt.Errorf("new dao, err: %w", err))
	}
	// SQL语句
	// sqlStr := "select name, singer, duration, link, cover, tag, platform from songs where id=?"
	// SELECT * FROM songs LIMIT 100
	sqlStr := "SELECT name, singer, duration, link, cover, tag, platform FROM songs LIMIT 100"
	rows, err := d.DB.Query(sqlStr)
	if err != nil {
		return nil, fmt.Errorf("query songs, err:%w", err)
	}

	//关闭Rows对象，释放连接资源
	defer rows.Close()

	var songs []model.Song
	//循环遍历Rows对象，获取每一行的数据
	for rows.Next() {
		var s model.Song
		err = rows.Scan(&s.Name, &s.Singer, &s.Duration, &s.Link, &s.Cover, &s.Tag, &s.Platform)
		if err != nil {
			return nil, fmt.Errorf("scan song, err: %w", err)
		}
		songs = append(songs, s)
	}
	fmt.Println("query success:", songs)
	return songs, nil
}

func (d *Dao) QueryMultiSongs(tag string, platform model.Platform) ([]*model.Song, error) {
	// 使用 colly 数据库
	_, err := global.DB.Exec("USE colly")
	if err != nil {
		// panic(err.Error())
		return nil, fmt.Errorf("select colly database, err: %w", err)
	}
	// SQL语句
	sqlStr := "SELECT * FROM songs WHERE tag = ? AND platform = ? LIMIT 0,100"
	rows, err := d.DB.Query(sqlStr, tag, platform)
	if err != nil {
		return nil, fmt.Errorf("query songs, err:%w", err)
	}
	//关闭Rows对象，释放连接资源
	defer rows.Close()

	var songs []*model.Song
	//循环遍历Rows对象，获取每一行的数据
	for rows.Next() {
		var s model.Song
		err = rows.Scan(&s.ID, &s.Name, &s.Singer, &s.Duration, &s.Link, &s.Cover, &s.Tag, &s.Platform)
		if err != nil {
			return nil, fmt.Errorf("scan song, err: %w", err)
		}
		songs = append(songs, &s)
	}

	fmt.Println("query success:", songs)
	return songs, nil
}
