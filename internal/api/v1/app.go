package app

import (
	"database/sql"

	"github.com/gorilla/mux"
)

type App struct {
	Router *mux.Router
	DB     *sql.DB
}

// func (a *App) Initialize(config){
// 	db:=
// }
