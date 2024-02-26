package models

import (
	"database/sql"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
)

type UserModel struct {
	db *sql.DB
}

func NewUserModel() *UserModel {
	conn, err := config.DBConn()

	if err != nil {
		panic(err)
	}

	return &UserModel{
		db: conn,
	}
}

func (u UserModel) FindAll() ([]entities.User, error) {
	rows, err := u.db.Query("select id, nama_lengkap, email, username, password from users")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []entities.User

	for rows.Next() {
		var user entities.User
		rows.Scan(&user.Id, &user.NamaLengkap, &user.Email, &user.Username, &user.Password)
		users = append(users, user)
	}

	return users, nil

}

func (u UserModel) Find(id int64) (entities.User, error) {
	row := u.db.QueryRow("select id, nama_lengkap, email, username, password from users where id = ?", id)

	var user entities.User

	err := row.Scan(&user.Id, &user.NamaLengkap, &user.Email, &user.Username, &user.Password)

	if err != nil {
		return user, err
	}

	return user, nil
}

func (u UserModel) Where(user *entities.User, fieldName, fieldValue string) error {

	row, err := u.db.Query("select id, nama_lengkap, email, username, password from users where "+fieldName+" = ? limit 1", fieldValue)

	if err != nil {
		return err
	}

	defer row.Close()

	for row.Next() {
		row.Scan(&user.Id, &user.NamaLengkap, &user.Email, &user.Username, &user.Password)
	}

	return nil
}

func (u UserModel) Create(user entities.User) (int64, error) {

	result, err := u.db.Exec("insert into users (nama_lengkap, email, username, password) values(?,?,?,?)",
		user.NamaLengkap, user.Email, user.Username, user.Password)

	if err != nil {
		return 0, err
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId, nil

}

func (p *UserModel) Update(user entities.User) error {

	_, err := p.db.Exec(
		"update user set nama_lengkap = ?, email = ?, username = ?, password = ? where id = ?",
		user.NamaLengkap, user.Email, user.Username, user.Password, user.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *UserModel) Delete(id int64) {
	p.db.Exec("delete from user where id = ?", id)
}
