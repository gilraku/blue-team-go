package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

func main() {
	host := flag.String("host", "", "Target host")
	start := flag.Int("start", 1, "Start port")
	end := flag.Int("end", 1024, "End port")
	timeout := flag.Duration("timeout", 500*time.Millisecond, "Connection timeout per port")
	concurrency := flag.Int("c", 100, "Concurrent scan goroutines")
	flag.Parse()

	if *host == "" {
		fmt.Fprintln(os.Stderr, "Usage: port-scanner -host <target> [-start 1] [-end 1024] [-timeout 500ms] [-c 100]")
		os.Exit(1)
	}

	fmt.Printf("Scanning %s ports %d-%d...\n\n", *host, *start, *end)

	ports := make(chan int, *concurrency)
	var wg sync.WaitGroup
	var mu sync.Mutex
	var open []int

	for i := 0; i < *concurrency; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for port := range ports {
				addr := fmt.Sprintf("%s:%d", *host, port)
				conn, err := net.DialTimeout("tcp", addr, *timeout)
				if err == nil {
					conn.Close()
					mu.Lock()
					open = append(open, port)
					mu.Unlock()
					fmt.Printf("  OPEN  %d\n", port)
				}
			}
		}()
	}

	for p := *start; p <= *end; p++ {
		ports <- p
	}
	close(ports)
	wg.Wait()

	fmt.Printf("\n--- %d open port(s) found ---\n", len(open))
}
