package main

import (
	"context"
	"fmt"
	"log"

	"main.go/pkg/repository"
)

const connString = "postgres://limbo:10007@localhost:5432/books"

func main() {
	psql, err := repository.New(context.Background(), connString)
	if err != nil {
		log.Fatal(err)
	}
	genres, err := psql.ReadAuthors(context.Background(), 1, 100)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(genres)
	psql.Pool.Close()
}
