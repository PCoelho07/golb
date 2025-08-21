package loadbalancer

import (
	"log"
	"net/http"
	"time"

	"github.com/PCoelho07/golb/internal/server"
)

type LoadBalancer struct {
    lbStrategy Strategy
	proxies []*server.Server
    currentProxy int
}

func NewLoadBalancer(srvrList []string, lbStrategy Strategy) *LoadBalancer {
    return &LoadBalancer{
        proxies: buildProxies(srvrList) ,
        currentProxy: 0,
        lbStrategy: lbStrategy,
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

    srvIndex := l.lbStrategy.ChooseServer()
    log.Printf("request to %d server", srvIndex)
    l.proxies[srvIndex].Proxy.ServeHTTP(rw, req)
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
