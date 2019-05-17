package main

import (
	"github.com/devarsh/vrpl/auth"
	"log"
	"net/http"
	"time"

	"github.com/devarsh/vrpl/master"
	"github.com/devarsh/vrpl/misc"
	"github.com/devarsh/vrpl/util"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"
	"github.com/go-chi/render"
	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

func init() {
	db, err := gorm.Open("mysql", "root:root@/mycooldb?parseTime=true&charset=utf8")
	if err != nil {
		panic(err)
	}
	db.SingularTable(true)
	db.Set("gorm:association_autoupdate", false)
	db.Set("gorm:association_autocreate", false)
	DB = db.Debug()
	util.CreateTables(DB)
}

func Routes() *chi.Mux {
	router := chi.NewRouter()
	cors := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "OPTIONS", "DELETE"},
		AllowedHeaders:   []string{"Accept", "Content-Type", "Authorization"},
		MaxAge:           300,
		AllowCredentials: true,
	})
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		cors.Handler,
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.Recoverer,
	)
	am := auth.NewAuthManager("HS256", "This_is_jwt_key", time.Duration(time.Minute*5), "VRPL_APP", DB)
	mm := misc.NewMiscManager(DB)
	mmaster := master.NewMasterManager(DB)

	router.Route("/v1", func(r chi.Router) {
		r.Mount("/api/auth", am.Routers())
		r.Mount("/api/misc", mm.Routers())
		r.Mount("/api/master", mmaster.Routers())
	})
	return router
}

func main() {
	router := Routes()
	log.Println("Started Listening on :8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}

func SecureRoute(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Secure Route"))
}
