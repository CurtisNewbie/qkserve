package main

import (
	"flag"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"sync/atomic"
	"time"
)

var (
	File    = flag.String("file", "", "file")
	Port    = flag.Int("port", 8080, "port")
	OneTime = flag.Bool("one-time", true, "one time use")
)

var (
	GracefulShutdown = 5 * time.Second

	Stopped int32 = 0
)

func main() {
	flag.Parse()
	if strings.TrimSpace(*File) == "" {
		println("Please specify file path")
		return
	}

	http.HandleFunc("/", ServeFile)
	fmt.Printf("Download file at: http://%v:%v\n", GetLocalIPV4(), *Port)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *Port), nil)
}

func ServeFile(w http.ResponseWriter, req *http.Request) {
	if *OneTime && !atomic.CompareAndSwapInt32(&Stopped, 0, 1) {
		fmt.Println("Server stopped")
		w.WriteHeader(404)
		return
	}

	byt, err := os.ReadFile(*File)
	if err != nil {
		fmt.Printf("Failed to read file, %v\n", err)
		w.WriteHeader(404)
		return
	}
	name := filepath.Base(*File)
	w.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=\"%v\"", url.QueryEscape(name)))
	_, err = w.Write(byt)
	if err != nil {
		fmt.Printf("Failed to write response, %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if *OneTime {
		fmt.Printf("Stopping server, server exit in %v\n", GracefulShutdown)
		go func() {
			time.Sleep(GracefulShutdown)
			os.Exit(0)
		}()
	}
}

// Get local ipv4 address (excluding loopback address)
func GetLocalIPV4() string {
	// src: https://stackoverflow.com/questions/23558425/how-do-i-get-the-local-ip-address-in-go
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		return ""
	}
	for _, address := range addrs {
		if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
			if ipnet.IP.To4() != nil {
				return ipnet.IP.String()
			}
		}
	}
	return ""
}
