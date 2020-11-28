package main

import (
	"net/http"
	"flag"
	"log"
	"errors"
	"LRUCache/lru"
)

type keyvalue struct {
	key string
	val string
}

type Server struct {
	cache lru.LRU
}


func main() {
	addr := flag.String("addr", ":8080", "endpoint address")
	s := &Server{lru.NewLRU(3)} // Initialize server with cache
	mux := http.NewServeMux()
	mux.HandleFunc("/cache", s.handleCache)
	log.Println("Starting webserver on", *addr)
	http.ListenAndServe(*addr, mux)
	log.Println("Stopping...")
}



// Handler
func (s *Server) handleCache(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case "GET":
		s.handleCacheGet(w,r)
		return

		case "POST":
		s.handleCachePost(w,r)
		return
	}
	respondHTTPErr(w, r, http.StatusNotFound)
}

// Post handler
func (s *Server) handleCachePost(w http.ResponseWriter, r *http.Request) {
	var kv keyvalue
	if err := decodeBody(r, &kv); err != nil { // Decode request
		respondErr(w, r, http.StatusBadRequest, "failed to read key value from request", err)
		return
	}
	if err := s.cache.Put(kv.key, kv.val); err != nil { // Insert to cache
		respondErr(w, r, http.StatusInternalServerError, "failed to insert", err)
		return
	}
	respond(w, r, http.StatusCreated, nil)

}



func (s *Server) handleCacheGet(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("GET not implemented"))
}
/*
// Default GET handler
func (s *Server) handleCachePost(w http.ResponseWriter, r *http.Request) {
	respondErr(w, r, http.StatusInternalServerError, errors.New("POST not implemented"))
}
*/

