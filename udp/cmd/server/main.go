package main

import (
	"flag"
	"fmt"
	"net"
	"os"
	"sync/atomic"
	"time"
)

var host = flag.String("host", "", "host")
var port = flag.String("port", "9090", "port")
var count int64

func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	fmt.Printf("Listening %v\n", addr)
	conn, err := net.ListenUDP("udp", addr)
	if err != nil {
		fmt.Println("Error listening:", err)
		os.Exit(1)
	}
	defer conn.Close()
	go func() {
		tick := time.NewTicker(5 * time.Second)
		for {
			select {
			case <-tick.C:
				fmt.Println("count: ", atomic.LoadInt64(&count))
			}
		}
	}()

	for {
		handleClient(conn)
	}
}
func handleClient(conn *net.UDPConn) {
	data := make([]byte, 1024)
	_, _, err := conn.ReadFromUDP(data)
	if err != nil {
		fmt.Println("failed to read UDP msg because of ", err.Error())
		return
	}
	atomic.AddInt64(&count, 1)

	// sec := binary.BigEndian.Uint64(data)
	// t := time.Unix(int64(sec), 0)
	// fmt.Println(n, remoteAddr, t.Format(time.RFC3339))

	// daytime := time.Now().Unix()
	// b := make([]byte, 4)
	// binary.BigEndian.PutUint32(b, uint32(daytime))
	// conn.WriteToUDP(b, remoteAddr)
}
