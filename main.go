package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/onyewuenyi/rest-server/configuration"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	var err error
	cfg, err := configuration.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	a.DB, err = sql.Open("postgres", cfg.Database.Conn_str())
	if err != nil {
		log.Fatal(err)
	}
	a.Router = mux.NewRouter()

	// a.initializeRoutes()
}

func (a *App) Run(addr string) {}

func main() {
	x := os.Getenv("pwd")
	cfg := fmt.Sprintf("%s/configuration", x)
	fmt.Println(cfg, x)
	a := App{}
	a.Initialize()
	// a.Run(":8080")
	// ctx := context.Background()
	// db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/tasks?sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// taskstore = taskstore.New(db)
	// tasks, err := taskstore.GetAll(ctx)
}
