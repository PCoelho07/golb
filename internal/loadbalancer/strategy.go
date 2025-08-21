package loadbalancer

import "sync"

type Strategy interface { 
    ChooseServer() int
}

type RRStrategy struct {
    currentReq int
    countServers int
    mu sync.Mutex
}

func NewRRStrategy(countServers int) *RRStrategy {
    return &RRStrategy{
        countServers: countServers,
        currentReq: 0,
    }
} 

func (rr *RRStrategy) ChooseServer() int {
    srvSelected := rr.currentReq % rr.countServers
    rr.incCurReq()

    return srvSelected
}

func (rr *RRStrategy) incCurReq() {
    rr.mu.Lock()
    rr.currentReq += 1
    rr.mu.Unlock()
}
