package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	account_http "github.com/cpartogi/izyai/internal/menu/handler/http"
	account_service "github.com/cpartogi/izyai/internal/menu/service"
	account_store "github.com/cpartogi/izyai/internal/menu/store"
	"github.com/cpartogi/izyai/pkg/common/config"

	db "github.com/cpartogi/izyai/pkg/common/lib/pgsql"

	"github.com/joho/godotenv"
	"github.com/julienschmidt/httprouter"
	"github.com/rs/cors"
)

func main() {
	errEnv := godotenv.Load()
	if errEnv != nil {
		log.Println(errEnv)
	}

	// to change the flags on the default logger
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// init config
	err := config.Init()
	if err != nil {
		log.Fatalln(err)
	}
	cfg := config.Get()

	// init postgres
	dbConn, err := db.Init(cfg.DB.Postgres)
	if err != nil {
		log.Fatalln(err)
	}

	defer dbConn.Close()

	// init router
	router := httprouter.New()

	//cors
	handler := cors.AllowAll().Handler(router)
	// accounts
	{
		store := account_store.NewStore(db.GetDB())
		accountSvc := account_service.NewService(store)

		account_http.NewHandler(router, accountSvc)
	}

	// start serve
	port := cfg.API.Port
	serve(port, handler)
}

func serve(port int, h http.Handler) {
	log.Println("starting server on:", port)

	s := &http.Server{
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		Addr:         fmt.Sprintf(":%d", port),
		Handler:      h,
	}

	log.Fatal(s.ListenAndServe())
}
