package models

import (
	"database/sql"

	"github.com/jeypc/go-auth/config"
	"github.com/jeypc/go-auth/entities"
)

type CatalogModel struct {
	conn *sql.DB
}

func NewCatalogModel() *CatalogModel {
	conn, err := config.DBConn()
	if err != nil {
		panic(err)
	}

	return &CatalogModel{
		conn: conn,
	}
}

func (c *CatalogModel) FindAll() ([]entities.Catalog, error) {
	rows, err := c.conn.Query("select id, bean, description, price from catalogs")

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var catalogs []entities.Catalog

	for rows.Next() {
		var catalog entities.Catalog
		rows.Scan(&catalog.Id, &catalog.Bean, &catalog.Description, &catalog.Price)
		catalogs = append(catalogs, catalog)
	}

	return catalogs, nil
}

func (c *CatalogModel) Find(id int64, catalog *entities.Catalog) error {

	return c.conn.QueryRow("select * from catalog where id = ?", id).Scan(
		&catalog.Id,
		&catalog.Bean,
		&catalog.Description,
		&catalog.Price)
}
