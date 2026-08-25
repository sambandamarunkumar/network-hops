package main

import (
	"fmt"
	"math"
	"math/rand"
	"sync"
	"sync/atomic"
	"time"
)

const (
	nodes      = 11
	clients    = 40
	requests   = 3000
	gridSize   = 10
)

type Node struct {
	id int
	x  float64
	y  float64
}

type Router struct {
	nodes []Node
}

var totalHops int64

func distance(aX, aY, bX, bY float64) float64 {
	return math.Sqrt((aX-bX)*(aX-bX) + (aY-bY)*(aY-bY))
}

func newRouter(n int) *Router {
	ns := make([]Node, n)
	for i := 0; i < n; i++ {
		ns[i] = Node{
			id: i,
			x:  rand.Float64() * gridSize,
			y:  rand.Float64() * gridSize,
		}
	}
	return &Router{nodes: ns}
}

func (r *Router) nearest(x, y float64) int {
	best := 0
	bestDist := math.MaxFloat64

	for i, n := range r.nodes {
		d := distance(x, y, n.x, n.y)
		if d < bestDist {
			bestDist = d
			best = i
		}
	}
	return best
}

func hopsFromDistance(d float64) int {
	if d < 2 {
		return 1
	}
	if d < 4 {
		return 2
	}
	if d < 6 {
		return 3
	}
	return 4
}

func client(id int, r *Router, wg *sync.WaitGroup) {
	defer wg.Done()

	x := rand.Float64() * gridSize
	y := rand.Float64() * gridSize

	for i := 0; i < requests; i++ {
		n := r.nearest(x, y)
		target := r.nodes[n]
		d := distance(x, y, target.x, target.y)
		h := hopsFromDistance(d)
		atomic.AddInt64(&totalHops, int64(h))
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	router := newRouter(nodes)

	start := time.Now()

	var wg sync.WaitGroup

	for i := 0; i < clients; i++ {
		wg.Add(1)
		go client(i, router, &wg)
	}

	wg.Wait()

	elapsed := time.Since(start)

	totalReq := clients * requests
	avgHops := float64(totalHops) / float64(totalReq)

	fmt.Println("Requests:", totalReq)
	fmt.Println("Average Hops:", avgHops)
	fmt.Println("Time:", elapsed)
}
