package main

import (
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"testingTask/internal/bootstrap"
	"testingTask/pkg/handlers"
	"testingTask/pkg/links"
)

func main() {
	//mode := flag.String("mode", "memory", "")
	//flag.Parse()

	mode := os.Getenv("WORKMODE")

	var linksRepo links.LinksRepo

	if mode == "memory" {
		linksRepo = bootstrap.InitMemoryRepo()
		log.Println("starting on memory")
	} else if mode == "db" {
		linksRepo = bootstrap.InitSQLRepo()
		log.Println("staring on postgresql db")
	} else {
		log.Fatalf("wrong mode error")
	}

	mux := http.NewServeMux()

	linksHandler := handlers.LinksHandler{
		LinksRepo: linksRepo,
	}

	mux.HandleFunc("/create", linksHandler.FromLongToShort) // POST method
	mux.HandleFunc("/get", linksHandler.FromShortToLong)    // GET method

	log.Println("starting at 8080")
	http.ListenAndServe(":8080", mux)

}
