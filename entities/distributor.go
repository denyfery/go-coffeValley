package entities

type Distributor struct {
	Id      int64
	Name    string `validate:"required"`
	City    string `validate:"required"`
	Region  string `validate:"required"`
	Country string `validate:"required"`
	Phone   string `validate:"required"`
	Email   string `validate:"required"`
}
