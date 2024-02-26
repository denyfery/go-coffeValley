package entities

type File struct {
	Id     int64
	Title  string `validate:"required"`
	File   string `validate:"required"`
	Author string `validate:"required"`
}
