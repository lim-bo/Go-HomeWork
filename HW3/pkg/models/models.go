package models

type Book struct {
	Id       int
	Name     string
	Price    int
	GenreId  int
	AuthorId int
}

type Genre struct {
	Id   int
	Name string
}

type Author struct {
	Id   int
	Name string
}
