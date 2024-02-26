package entities

type Catalog struct {
	Id          int64
	Bean        string `validate:"required"`
	Description string `validate:"required"`
	Price       string `validate:"required"`
}
