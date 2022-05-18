package main

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"testingTask/internal/bootstrap"
	"testingTask/pkg/handlers"
	"testingTask/pkg/links"
)

func main() {
	mode := flag.String("mode", "memory", "")
	flag.Parse()

	var linksRepo links.LinksRepo

	if *mode == "memory" {
		linksRepo = bootstrap.InitMemoryRepo()
	} else if *mode == "db" {
		linksRepo = bootstrap.InitSQLRepo()
	} else {
		log.Fatalf("frong mode error")
	}

	mux := http.NewServeMux()

	linksHandler := handlers.LinksHandler{
		LinksRepo: linksRepo,
	}

	mux.HandleFunc("/getShort", linksHandler.FromLongToShort) // POST method
	mux.HandleFunc("/getLong", linksHandler.FromShortToLong)  // GET method

	fmt.Println("starting at 8080")
	http.ListenAndServe(":8080", mux)

}
