package server

import (
	"log"
	"net/http/httputil"
	"net/url"
	"sync"
)

type Server struct {
    Url string
    Proxy *httputil.ReverseProxy
    isAlive bool
    mu sync.Mutex
}

func NewServer(uri string) *Server { 
    u, err := url.Parse(uri)
    if err != nil {
        log.Fatalf("fail when attempt parse URL: %s", err)
    }

    return &Server{
        Url: uri,
        Proxy: httputil.NewSingleHostReverseProxy(u),
        isAlive: true,
    }
}

func (s *Server) IsNotAlive() {
    s.mu.Lock()
    s.isAlive = false
    s.mu.Unlock()
}

func (s *Server) IsAlive() {
    s.mu.Lock()
    s.isAlive = true
    s.mu.Unlock()
}

