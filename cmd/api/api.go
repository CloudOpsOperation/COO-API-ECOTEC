package api

import (
	"database/sql"
	"log"
	"net/http"

	treeinfo "github.com/CloudOpsOperation/COO-API-ECOTEC/service/treeInfo"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

type API struct {
	address string
	db      *sql.DB
}

func NewAPI(address string, db *sql.DB) *API {
	return &API{
		address: address,
		db:      db,
	}
}

func (a *API) Start() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	storage := treeinfo.NewStorage(a.db)
	treeInfo := treeinfo.NewTreeInfo(storage)
	treeInfo.RegisterRoutes(subrouter)

	log.Print("Starting API on ", a.address)

	// Configura el middleware CORS
	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type", "Authorization"}),
	)

	return http.ListenAndServe(a.address, corsMiddleware(router))
}
