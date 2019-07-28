package main

import (
	"encoding/binary"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var host = flag.String("host", "127.0.0.1", "host")
var port = flag.String("port", "9090", "port")
var number = flag.Int("number", 3600, "times")

//go run timeclient.go -host time.nist.gov
func main() {
	flag.Parse()
	addr, err := net.ResolveUDPAddr("udp", *host+":"+*port)
	if err != nil {
		fmt.Println("Can't resolve address: ", err)
		os.Exit(1)
	}
	waitGroup := sync.WaitGroup{}
	for i := 0; i < *number; i++ {
		// waitGroup.Add(i)
		// fmt.Printf("Send gourp %v\n", i)
		// go func(i int) {
		// 	for j := 0; j < 1000; j++ {
		// 		send(i, j, addr)
		// 	}
		// 	waitGroup.Done()
		// }(i)
		fmt.Printf("Sending %v\n", i)
		send(i, 0, addr)
		time.Sleep(time.Second)
	}
	waitGroup.Wait()
}

func send(i, j int, addr *net.UDPAddr) {
	conn, err := net.DialUDP("udp", nil, addr)
	if err != nil {
		fmt.Println("Can't dial: ", err)
		os.Exit(1)
	}
	defer conn.Close()
	daytime := time.Now().Unix()
	d := make([]byte, 4)
	binary.BigEndian.PutUint32(d, uint32(daytime))
	_, err = conn.Write(d)
	if err != nil {
		fmt.Println("failed:", err)
		os.Exit(1)
	}
	// data := make([]byte, 4)
	// _, err = conn.Read(data)
	// if err != nil {
	// 	fmt.Println("failed to read UDP msg because of ", err)
	// 	os.Exit(1)
	// }
	// t := binary.BigEndian.Uint32(data)
	// fmt.Println(i, j, time.Unix(int64(t), 0).Format(time.RFC3339))
}
