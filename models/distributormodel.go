package models

import (
	"database/sql"
	"fmt"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
)

type DistributorModel struct {
	conn *sql.DB
}

func NewDistributorModel() *DistributorModel {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &DistributorModel{
		conn: conn,
	}
}

func (d *DistributorModel) FindAll() ([]entities.Distributor, error) {
	rows, err := d.conn.Query("select id, name, city, region, country, phone, email from distributors")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var distributors []entities.Distributor

	for rows.Next() {
		var distributor entities.Distributor
		rows.Scan(&distributor.Id,
			&distributor.Name,
			&distributor.City,
			&distributor.Region,
			&distributor.Country,
			&distributor.Phone,
			&distributor.Email,
		)
		distributors = append(distributors, distributor)
	}

	return distributors, nil
}

func (d *DistributorModel) Find(id int64, distributor *entities.Distributor) error {

	return d.conn.QueryRow("select * from distributors where id = ?", id).Scan(
		&distributor.Id,
		&distributor.Name,
		&distributor.City,
		&distributor.Region,
		&distributor.Country,
		&distributor.Phone,
		&distributor.Email,
	)
}

func (d *DistributorModel) Create(distributor entities.Distributor) bool {

	result, err := d.conn.Exec("insert into distributors (name, city, region, country, phone, email) values(?,?,?,?,?,?)",
		distributor.Name, distributor.City, distributor.Region, distributor.Country, distributor.Phone, distributor.Email)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (d *DistributorModel) Update(distributor entities.Distributor) error {

	_, err := d.conn.Exec(
		"UPDATE distributors SET name = ?, city = ?, country = ?, region = ?, phone = ?, email = ? WHERE id = ?",
		distributor.Name,
		distributor.City,
		distributor.Country,
		distributor.Region,
		distributor.Phone,
		distributor.Email,
		distributor.Id)

	if err != nil {
		return err
	}

	return nil
}

func (p *DistributorModel) Delete(id int64) {
	p.conn.Exec("delete from distributors where id = ?", id)
}
