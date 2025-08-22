package main

import (
	"log"
	"net/http"
	"time"

	"github.com/PCoelho07/golb/internal/loadbalancer"
)

func main() {
	serverListUrl := []string{"http://localhost:3000", "http://localhost:3001"}
    rrStrategy := loadbalancer.NewRRStrategy(len(serverListUrl))
    lb := loadbalancer.NewLoadBalancer(serverListUrl, rrStrategy)

	mainServer := http.Server{
		Addr:    ":8080",
		Handler: lb,
	}

    go lb.HealthCheck(1 * time.Second)

	if err := mainServer.ListenAndServe(); err != nil {
		log.Fatalf("fail when attempt to listen: %s", err)
	}
}
