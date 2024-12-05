/*func main() {

	e := echo.New()

		//Logger регистрирует HTTP-запросы.
	e.Use(middleware.Logger())

		//Recover восстанавливается после сбоев в любой точке цепочки
		//и передает управление централизованному HTTPErrorHandler.
	e.Use(middleware.Recover())

	// инициирует бд
	db, err := initStore()
	if err != nil {
		log.Fatalf("failed to initialise the store: %s", err)
	}
	defer db.Close()

	e.GET("/", func(c echo.Context) error {
		// отображает количество строк в таблице базы данных
		return rootHandler(db, c)
	})

	e.GET("/ping", func(c echo.Context) error {
		// просто показывает что статус ок. хз зачем
		return c.JSON(http.StatusOK, struct{ Status string }{Status: "пошел нахуй"})
	})

	e.POST("/send", func(c echo.Context) error {
		return sendHandler(db, c)
	})

	httpPort := os.Getenv("HTTP_PORT")
	if httpPort == "" {
		httpPort = "8000"
	}

	e.Logger.Fatal(e.Start(":" + httpPort))
}

type Message struct {
	Value string `json:"value"`
}

func initStore() (*sql.DB, error) {
	// инициирует бд

	pgConnString := fmt.Sprintf("host=%s port=%s dbname=%s user=%s password=%s sslmode=disable",
		os.Getenv("PGHOST"),
		os.Getenv("PGPORT"),
		os.Getenv("PGDATABASE"),
		os.Getenv("PGUSER"),
		os.Getenv("PGPASSWORD"),
	)

	var (
		db  *sql.DB
		err error
	)
	openDB := func() error {
		db, err = sql.Open("postgres", pgConnString)
		return err
	}

	// повторение запросов в бд. хз что это. каких запросов?
	err = backoff.Retry(openDB, backoff.NewExponentialBackOff())
	if err != nil {
		return nil, err
	}

	// что то типа приветствия когда впервые заходишь в бд через exec
	if _, err := db.Exec(
		"CREATE TABLE IF NOT EXISTS message (value TEXT PRIMARY KEY)"); err != nil {
		return nil, err
	}

	return db, nil
}

func rootHandler(db *sql.DB, c echo.Context) error {
	// отображает количество строк в таблице базы данных
	r, err := countRecords(db)
	if err != nil {
		return c.HTML(http.StatusInternalServerError, err.Error())
	}
	return c.HTML(http.StatusOK, fmt.Sprintf("Hello, Docker! (%d)\n", r))
}

func sendHandler(db *sql.DB, c echo.Context) error {
	// хуй знает нахуй это все надо но
	// к запуску встроенного сервера
	// это все равно отношения не имеет
	m := &Message{}

	if err := c.Bind(m); err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	err := crdb.ExecuteTx(context.Background(), db, nil,
		func(tx *sql.Tx) error {
			_, err := tx.Exec(
				"INSERT INTO message (value) VALUES ($1) ON CONFLICT (value) DO UPDATE SET value = excluded.value",
				m.Value,
			)
			if err != nil {
				return c.JSON(http.StatusInternalServerError, err)
			}
			return nil
		})

	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}

	return c.JSON(http.StatusOK, m)
}

func countRecords(db *sql.DB) (int, error) {
	//возвращает количество строк в таблице

	rows, err := db.Query("SELECT COUNT(*) FROM message")
	if err != nil {
		return 0, err
	}
	defer rows.Close()

	count := 0
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return 0, err
		}
		rows.Close()
	}

	return count, nil
}*/

// main.go

/*import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8000", nil)
}*/

package main

import (
	"fmt"
	"net/http"
)

/*import (
    "database/sql"
    "fmt"
    "log"
)

func main() {
    // Настроить строку подключения
    dsn := "root:root@tcp(127.0.0.1:3306)/mydb"
    // Открыть подключение к базе данных
    db, err := sql.Open("mysql", dsn)
    if err != nil {
        log.Fatal(err)
    }
    defer db.Close()
    // Проверить подключение
    if err := db.Ping(); err != nil {
        log.Fatal(err)
    }
    fmt.Println("Подключён к базе данных MySQL успешно!")
}*/

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, World!")
	})

	http.ListenAndServe(":8000", nil)
}
