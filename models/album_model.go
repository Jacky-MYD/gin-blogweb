package models

import (
	"gin/blogWeb_gin/utils"
	db "gin/blogWeb_gin/database"
)

type Album struct {
	Id         int
	Filepath   string
	Filename   string
	Status     int
	CreateTime int64
}

//-------插入图片---------------
func (album * Album) InsertAlbum() (int64, error) {
	return utils.ModifyDB("insert into album(id,filepath,filename,status,createtime)values(?,?,?,?,?)",
		album.Id, album.Filepath, album.Filename, album.Status, album.CreateTime)
}

//--------查询图片----------
func FindAllAlbums() ([]Album, error) {
	rows, err := db.QueryDB("select id,filepath,filename,status,createtime from album")
	if err != nil {
		return nil, err
	}
	var albums []Album
	for rows.Next() {
		id := 0
		filepath := ""
		filename := ""
		status := 0
		var createtime int64
		createtime = 0
		rows.Scan(&id, &filepath, &filename, &status, &createtime)
		album := Album{id, filepath, filename, status, createtime}
		albums = append(albums, album)
	}
	return albums, nil
}
