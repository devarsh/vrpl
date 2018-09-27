package main

/*
import (
	"github.com/go-chi/chi"
	"log"
	"net/http"
)

func main() {
	router := chi.NewRouter()
	router.Route("/v1", func(r chi.Router) {
		r.Get("/abc/{id}", myHandler)
		r.Post("/", myHandler)
		r.Put("/", myHandler)
		r.MethodNotAllowed(nandler)
	})
	log.Println("Started Listening on :8083")
	log.Fatal(http.ListenAndServe(":8083", router))
}

func myHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is dummy handler"))
}

func nandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("This is no  handler"))
}
*/
