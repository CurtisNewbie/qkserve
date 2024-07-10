package main

import (
	"flag"
	"fmt"
	"math/rand"
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
	Port    = flag.Int("port", 80, "port")
	OneTime = flag.Bool("one-time", true, "file served for one time only")
)

var (
	GracefulShutdown = 2 * time.Second

	Stopped int32 = 0
	digits        = "0123456789"
)

func main() {
	flag.Parse()
	if strings.TrimSpace(*File) == "" {
		println("Please specify file path")
		return
	}

	rtk := RandNum(15)
	http.HandleFunc("/"+rtk, ServeFile)
	fmt.Printf("\nDownload file at: http://%v:%v/%v\n", GetLocalIPV4(), *Port, rtk)
	http.ListenAndServe(fmt.Sprintf("0.0.0.0:%d", *Port), nil)
}

func ServeFile(w http.ResponseWriter, req *http.Request) {
	if *OneTime && atomic.LoadInt32(&Stopped) != 0 {
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

	atomic.StoreInt32(&Stopped, 1)
	if *OneTime {
		if fl, ok := w.(http.Flusher); ok {
			fl.Flush()
		}
		go func() {
			fmt.Printf("Server exit in %v\n", GracefulShutdown)
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

func doRand(n int, set []rune) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = set[rand.Intn(len(set))]
	}
	return string(b)
}

// Generate random numeric string with specified length
//
// the generated string will contains [0-9]
func RandNum(n int) string {
	return doRand(n, []rune(digits))
}
