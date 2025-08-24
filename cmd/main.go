package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/PCoelho07/golb/internal/interpreter"
	"github.com/PCoelho07/golb/internal/loadbalancer"
)

func main() {
	lbConfig := parseConfig()
    serverListUrl := lbConfig.BackendUrls()

    if len(serverListUrl) <= 0 {
        log.Fatalln("backend list is empty")
    }

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

func parseConfig() *loadbalancer.LoadBalancerConfig {
    now := time.Now()
    lx := interpreter.NewLexer("./golb.conf")
    if err := lx.Tokenize(); err != nil {
        log.Fatal(err)
    }

    parser := interpreter.NewParser(lx.TokenList)
    lbConfig, err := parser.Parse()
    if err != nil {
        log.Fatal(err)
    }

    fmt.Printf("parsing took %s\n", time.Since(now))

    return lbConfig
}
