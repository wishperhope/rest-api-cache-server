package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/valyala/fasthttp"
)

// all handler not seperating route to optimaze response time
func (s *server) handler(ctx *fasthttp.RequestCtx) {
	method := string(ctx.Method())
	path := string(ctx.Path())

	// Root path cannot be key to cache
	if path == "/" {
		ctx.SetContentType("text/plain; charset=utf8")
		fmt.Fprint(ctx, "Hello, world!\n\n")
		return
	}

	// Static auth
	if string(ctx.Request.Header.Peek("Authorization")) != s.token {
		ctx.Error("Bad Request", fasthttp.StatusBadRequest)
		log.Printf("wrong token : %s", string(ctx.Request.Header.Peek("Authorization")))
		return
	}

	// get from example bigcache server
	if method == http.MethodGet && path == "/stats" {
		target, err := json.Marshal(s.cache.Stats())
		if err != nil {
			ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			log.Printf("cannot marshal cache stats. error: %s", err)
			return
		}
		// since we're sending a struct, make it easy for consumers to interface.
		ctx.SetContentType("application/json; charset=utf8")
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
		fmt.Fprint(ctx, string(target))
		return
	}

	if method == http.MethodGet {
		data, err := s.cache.Get(path)
		if err != nil {
			ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			log.Printf("cannot get cache of path error: %s", err)
			return
		}

		// delete data if have query delete=y
		if string(ctx.FormValue("delete")) == "y" {
			if err := s.cache.Delete(path); err != nil {
				ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
				log.Printf("cannot delete cache of path error: %s", err)
				return
			}
		}
		ctx.SetContentType("application/json; charset=utf8")
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
		fmt.Fprint(ctx, string(data))
		return
	}

	if method == http.MethodPost {
		err := s.cache.Set(path, ctx.FormValue("data"))
		if err != nil {
			ctx.Error("Internal Server Error", fasthttp.StatusInternalServerError)
			log.Printf("cannot set cache of path error: %s", err)
			return
		}
		ctx.SetContentType("application/json; charset=utf8")
		ctx.Response.SetStatusCode(fasthttp.StatusOK)
		fmt.Fprint(ctx, "{\"success\": true}")
		return
	}
}
