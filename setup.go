package main

import (
	"errors"
	"log"
	"os"
	"strconv"
	"time"

	"github.com/allegro/bigcache/v2"
)

func (s *server) setup() error {
	// Setup Bigcache
	expireTime, err := strconv.Atoi(os.Getenv("CACHE_EXPIRE_TIME"))
	if err != nil {
		log.Println("invalid number of CACHE_EXPIRE_TIME on env file")
		return err
	}

	s.cache, err = bigcache.NewBigCache(bigcache.DefaultConfig(time.Duration(expireTime) * time.Minute))
	if err != nil {
		log.Println("Error when setup cache")
		return err
	}

	s.token = os.Getenv("APP_KEY")
	if s.token == "" {
		return errors.New("APP_KEY not set")
	}
	return nil
}
