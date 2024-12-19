package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

/*func main() {
	db, err := sqlConn()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("db connected")

	albums, err := albumsByArtist("John Coltrane", db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("albums found: %v\n", albums)

	alb, err := albumsByID(2, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("album found: %v\n", alb)

	albID, err := albumAdd(Album{
		Title:  "The Modern Sound of Betty Carter",
		Artist: "Betty Carter",
		Price:  49.99,
	}, db)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("ID of added album: %v\n", albID)

	defer db.Close()
}*/

func SqlConn() (*sql.DB, error) {
	var db *sql.DB

	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}

	cfg := mysql.Config{
		User:   os.Getenv("DBUSER"),
		Passwd: os.Getenv("DBPASS"),
		Net:    "tcp",
		Addr:   "127.0.0.1:3306",
		DBName: "db",
	}

	db, err = sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	if err := db.Ping(); err != nil {
		log.Fatal(err)
	}

	return db, nil
}

func AlbumsByArtist(name string, db *sql.DB) ([]Album, error) {
	var albums []Album

	rows, err := db.Query("select * from album where artist = ?", name)
	if err != nil {
		return nil, fmt.Errorf("albumByArtist %q:%v", name, err)
	}
	defer rows.Close()

	for rows.Next() {
		var alb Album

		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
			return nil, fmt.Errorf("albumByArtist %q:%v", name, err)
		}
		albums = append(albums, alb)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
	}

	return albums, nil
}

func AlbumsByID(id byte, db *sql.DB) (Album, error) {
	var alb Album

	row := db.QueryRow("select * from album where id = ?", id)
	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
		if err == sql.ErrNoRows {
			return alb, fmt.Errorf("albumsById %d: no such album", id)
		}
		return alb, fmt.Errorf("albumsById %d: %v", id, err)
	}
	return alb, nil
}

func AlbumAdd(alb Album, db *sql.DB) (int64, error) {
	result, err := db.Exec("insert into album (title, artist, price) values (?, ?, ?)",
		alb.Title, alb.Artist, alb.Price)
	if err != nil {
		return 0, fmt.Errorf("albumAdd: %v", err)
	}

	id, err := result.LastInsertId()
	if err != nil {
		return 0, fmt.Errorf("albumAdd: %v", err)
	}

	return id, nil
}
