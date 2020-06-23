package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//to get and use the data from the database, we need to define a struct
type Post struct {
	ID    int
	Title string
}

func main() {
	db, err := sql.Open("mysql", "root:Sagarmatha1@tcp(127.0.0.1:3306)/gosql")

	if err != nil {
		panic(err.Error())
	}

	defer db.Close()

	// insert, err := db.Query("INSERT INTO posts(id, title) VALUES (4, 'our post')")
	// if err != nil {
	// 	panic(err.Error())
	// }

	// defer insert.Close()
	posts, err := db.Query("SELECT id, title FROM posts")
	if err != nil {
		panic(err.Error())
	}

	for posts.Next() {
		var post Post
		err := posts.Scan(&post.ID, &post.Title)
		if err != nil {
			panic(err.Error())
		}
		fmt.Println(post)
	}
}
