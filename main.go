package main

import (
	"context"
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"github.com/gorilla/websocket"
)

var (

	// any wss address!
	sourceURL = "wss://echo.websocket.org/"

	proxyservice = "brd.superproxy.io:22225"
	userName     = "brd-customer-<CUSTOMER>-zone-<ZONE>"
	password     = "<PASSWORD>"

	headers = http.Header{
		"User-Agent": {
			"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/117.0.0.0 Safari/537.36",
		},
		"Accept": {
			"*/*",
		},
	}
)

func main() {

	count := 3

	for count > 0 {

		<-time.After(time.Second * 1)

		count--

		u := fmt.Sprintf("http://%s:%s@%s", userName, password, proxyservice)

		p, err := url.Parse(u)
		if err != nil {
			log.Println(err)
			return
		}

		dialer := websocket.Dialer{
			Proxy: http.ProxyURL(p),
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		}

		con, _, err := dialer.DialContext(context.TODO(), sourceURL, headers)
		if err != nil {
			log.Println(err)
			return
		}

		log.Printf("remote addr: %v\n", con.RemoteAddr().String())
	}
}
