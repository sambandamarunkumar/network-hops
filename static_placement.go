import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	nodes       = 5
	clients     = 50
	requests    = 2000
	hopDelayMs  = 2
)

type Server struct {
	id int
}

type Gateway struct {
	servers []Server
}

func NewGateway(n int) *Gateway {
	s := make([]Server, n)
	for i := 0; i < n; i++ {
		s[i] = Server{id: i}
	}
	return &Gateway{servers: s}
}

func hop() {
	time.Sleep(time.Duration(hopDelayMs) * time.Millisecond)
}

func (g *Gateway) route(req int) int {
	hops := 0

	hop()
	hops++

	hop()
	hops++

	target := rand.Intn(len(g.servers))

	hop()
	hops++

	_ = target

	return hops
}

func client(id int, g *Gateway, wg *sync.WaitGroup, results chan int) {
	defer wg.Done()

	for i := 0; i < requests; i++ {
		h := g.route(i)
		results <- h
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	gateway := NewGateway(nodes)

	var wg sync.WaitGroup
	results := make(chan int, clients*requests)

	start := time.Now()

	for i := 0; i < clients; i++ {
		wg.Add(1)
		go client(i, gateway, &wg, results)
	}

	wg.Wait()
	close(results)

	totalHops := 0
	count := 0

	for h := range results {
		totalHops += h
		count++
	}

	elapsed := time.Since(start)

	fmt.Println("Total Requests:", count)
	fmt.Println("Average Hops:", float64(totalHops)/float64(count))
	fmt.Println("Execution Time:", elapsed)
}
