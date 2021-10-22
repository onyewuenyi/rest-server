package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/onyewuenyi/rest-server\configuration"
)
C:\Users\Charles A Onyewuenyi\dev\go-projs\rest-server\configuration
type App struct {
	Router *mux.Router
	DB     *sql.DB
}

func (a *App) Initialize() {
	cfg := configuration
	cfg.Database.conn_str()

	var err error
	a.DB, err = sql.Open("postgres", cfg.Database.conn_str())
	if err != nil {
		log.Fatal(err)
	}

	a.Router = mux.NewRouter()

}

func main() {
	mux := http.NewServeMux()

	// ctx := context.Background()
	// db, err := sql.Open("postgres", "postgres://postgres:postgres@localhost/tasks?sslmode=disable")
	// if err != nil {
	// 	panic(err)
	// }
	// defer db.Close()
	// taskstore = taskstore.New(db)
	// tasks, err := taskstore.GetAll(ctx)
}
