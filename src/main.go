package main

// This is a simple "curl" client implementation for basic docker images like scratch or alpine,
import (
	"flag"
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
	"net/url"
	"os"
	"time"
)

var healthReq http.Request

var timeout time.Duration

func main() {
	healthCheck := ""
	timeOut := ""
	flag.StringVar(&healthCheck, "h", "", "start health check client")
	flag.StringVar(&timeOut, "timeout", "10s", "defines timeout for health check request")
	flag.Parse()
	if healthCheck != "" {
		tout, err := time.ParseDuration(timeOut)
		if err != nil {
			log.Warn(fmt.Sprintf("could not parse timeout(expecting duration value like 1s or 10m) parameter with value %s. Using default 10s", timeOut))
			timeout = 10 * time.Second
		}
		timeout = tout
		u, err := url.Parse(healthCheck)
		if err != nil {
			log.Fatalf("could not parse url %s", healthCheck)
		}
		healthReq = http.Request{Method: http.MethodGet, URL: u}
		os.Exit(CheckHealth())
		return
	}
	log.Info("no url for health check configured!")
}

func CheckHealth() int {
	client := &http.Client{Timeout: timeout}
	resp, er := client.Do(&healthReq)

	if er != nil {
		log.Fatalf("failed to execute health check on endpoint %s with error %v", healthReq.URL, er)
		return 1
	} else {
		if resp.StatusCode != 200 {
			log.Fatalf("healthcheck response returned %d status code!", resp.StatusCode)
			return 1
		}
		log.Println("successful health check")
		return 0
	}
}

