package models

type Book struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	Price    int    `json:"price"`
	GenreId  int    `json:"genre_id"`
	AuthorId int    `json:"author_id"`
}

type Genre struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}
