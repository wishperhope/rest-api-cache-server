package main

import (
	"log"
	"os"

	"github.com/allegro/bigcache/v2"
	"github.com/joho/godotenv"
	"github.com/valyala/fasthttp"
)

type server struct {
	cache *bigcache.BigCache
	token string
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error Loading .env File")
	}

	s := server{}

	err = s.setup()
	if err != nil {
		log.Panicln(err)
	}

	port := os.Getenv("PORT")
	if err := fasthttp.ListenAndServe(":"+port, s.handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}
