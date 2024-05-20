package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
)

func main() {
	// Ask for the IP address to scan
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter IP address to scan: ")
	ip, _ := reader.ReadString('\n')

	// Remove the newline character
	ip = ip[:len(ip)-1]

	startPort := 1
	endPort := 1024

	var wg sync.WaitGroup

	// Iterate over the port range
	for port := startPort; port <= endPort; port++ {
		wg.Add(1)
		go func(port int) {
			defer wg.Done()
			address := ip + ":" + strconv.Itoa(port)

			// Try to connect to the port
			conn, err := net.DialTimeout("tcp", address, 10*time.Second)

			// If there's no error, the port is open
			if err == nil {
				fmt.Printf("Port %d is open\n", port)
				err := conn.Close()
				if err != nil {
					return
				}
			}
		}(port)
	}
	wg.Wait()
}
