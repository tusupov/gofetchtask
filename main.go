package main

import (
	"flag"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/tusupov/gofetchtask/handle"
	"github.com/tusupov/gofetchtask/middleware"
)

var port = flag.Int("port", 8080, "Server port")

func init() {
	flag.Parse()
}

func main() {

	r := mux.NewRouter()

	// Handle function
	r.HandleFunc("/add", handle.Task)
	r.HandleFunc("/list", handle.List)
	r.HandleFunc("/delete", handle.Delete)

	// Middleware
	middleware.SetLogger(os.Stderr)
	r.Use(middleware.Panic, middleware.AccessLog)

	// Start server
	log.Printf("Listening port [%d] ...", *port)
	if err := http.ListenAndServe(":"+strconv.Itoa(*port), r); err != nil {
		log.Fatal(err)
	}

}
