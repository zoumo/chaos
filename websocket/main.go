package main

import (
	"flag"
	"net/http"
	"time"

	// "github.com/spf13/pflag"

	"github.com/golang/glog"
	"github.com/gorilla/websocket"
)

func main() {
	flag.Set("logtostderr", "true")
	flag.Parse()
	defer glog.Flush()

	header := http.Header{}
	header.Add("Authorization", "Basic YWRtaW46WkNuZVlJQjhua2dtWGUyeA==")
	header.Add("X-Forwarded-Proto", "http")
	// header.Add("Sec-Websocket-Protocol", "base64.channel.k8s.io")
	url := "wss://dev.caicloudprivatetest.com/api/v1/namespaces/clever/pods/terminal-24-3936308395-4g5h1/exec?container=fileserver&stdout=1&stdin=1&stderr=1&tty=1&command=bash&command=-i&uid=1467284823&access_token=6mbgYIf7JB1imM9BRRrSiUWM1eA"
	conn, _, err := websocket.DefaultDialer.Dial(url, header)
	if err != nil {
		glog.Error(err)
		return
	}
	defer conn.Close()

	done := make(chan struct{})
	go func() {
		defer conn.Close()
		defer close(done)
		for {
			conn.SetReadDeadline(time.Now().Add(5 * time.Second))
			_, message, err := conn.ReadMessage()
			if err != nil {
				glog.Errorf("read mesage err: %v", err)
				return
			}
			glog.Infof("revc: %v", string(message))
		}
	}()

	<-done

}
