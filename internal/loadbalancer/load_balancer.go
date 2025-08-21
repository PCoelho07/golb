package loadbalancer

import (
	"log"
	"net/http"
	"time"

	"github.com/PCoelho07/golb/internal/server"
)

type LoadBalancer struct {
	proxies []*server.Server
    currentProxy int
}

func NewLoadBalancer(srvrList []string) *LoadBalancer {
    return &LoadBalancer{
        proxies: buildProxies(srvrList) ,
        currentProxy: 0,
    }
}

func buildProxies(srvrList []string) []*server.Server {
    var serverList []*server.Server
    for _, serverUrl := range srvrList {
        srv := server.NewServer(serverUrl)
        serverList = append(serverList, srv)
    }

    return serverList
}

func (l *LoadBalancer) ServeHTTP(rw http.ResponseWriter, req *http.Request) { 
    log.Printf("request %s from %s", req.Body, req.RemoteAddr)
    l.proxies[0].Proxy.ServeHTTP(rw, req)
}

func (l *LoadBalancer) HealthCheck(duration time.Duration) {
    ticker := time.NewTicker(duration)

    for range ticker.C { 
        for _, p := range l.proxies {
            _, err := http.Get(p.Url)
            if err != nil {
                log.Printf("Server %s is down.", p.Url)
                p.IsNotAlive()
            } else {
                log.Printf("Server %s is up.", p.Url)
                p.IsAlive()
            }
        }
    }
}
