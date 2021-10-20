package repository

import (
	"database/sql"
	"log"
	"os"

	"github.com/suatacikel/golang-rest-api/entity"
)

type sqliteRepo struct{}

func NewSQLiteRepository() PostRepository {
	os.Remove("./posts.db")

	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sqlStmt := "create table post (id intager not null primary key, litle text, txt text); delete from posts;"

	_, err = db.Exec(sqlStmt)
	if err != nil {
		log.Printf("%q: %s/n", err, sqlStmt)
	}
	return &sqliteRepo{}
}

func (*sqliteRepo) Save(post *entity.Post) (*entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	stmt, err := tx.Prepare("insert into posts(id, title, txt) values (?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer stmt.Close()

	_, err = stmt.Exec(post.Id, post.Title, post.Text)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	err = tx.Commit()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return post, nil
}

func (*sqliteRepo) FindAll() ([]entity.Post, error) {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	rows, err := db.Query("select id, title, txt from posts")
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	defer rows.Close()

	var posts []entity.Post

	for rows.Next() {
		var id int
		var title string
		var text string
		err = rows.Scan(&id, &title, &text)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		post := entity.Post{
			Id:    id,
			Title: title,
			Text:  text,
		}
		posts = append(posts, post)
	}
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
		return nil, err
	}

	return posts, nil
}

func (*sqliteRepo) Delete(post *entity.Post) error {
	db, err := sql.Open("sqlite3", "./posts.db")
	if err != nil {
		log.Fatal(err)
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		log.Fatal(err)
		return err
	}
	stmt, err := tx.Prepare("delete from posts where id = ?")
	if err != nil {
		log.Fatal(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(post.Id)
	if err != nil {
		log.Fatal(err)
		return err
	}
	tx.Commit()
	return nil
}
