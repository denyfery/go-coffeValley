package models

import (
	"database/sql"
	"fmt"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
)

type FileModel struct {
	conn *sql.DB
}

func NewFileModel() *FileModel {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &FileModel{
		conn: conn,
	}
}

func (f *FileModel) FindAll() ([]entities.File, error) {
	rows, err := f.conn.Query("select id, title, file, author from files")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var files []entities.File

	for rows.Next() {
		var file entities.File
		rows.Scan(&file.Id, &file.Title, &file.File, &file.Author)
		files = append(files, file)
	}

	return files, nil
}

func (f *FileModel) Find(id int64, file *entities.File) error {

	return f.conn.QueryRow("select * from file where id = ?", id).Scan(
		&file.Id,
		&file.Title,
		&file.File,
		&file.Author)
}

func (f *FileModel) Create(file entities.File) bool {

	result, err := f.conn.Exec("insert into files (title, file, author) values(?,?,?)",
		file.Title, file.File, file.Author)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}
