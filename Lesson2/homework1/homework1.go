// 接收客户端 request，并将 request 中带的 header 写入 response header
// 读取当前系统的环境变量中的 VERSION 配置，并写入 response header
// Server 端记录访问日志包括客户端 IP，HTTP 返回码，输出到 server 端的标准输出
// 当访问 {url}/healthz 时，应返回200

package main

import (
	"fmt"
	"log"
	"net/http"
	"net/http/pprof"
	"os"
	"strings"
)

func index(w http.ResponseWriter, r *http.Request) {
	// 设置version
	os.Setenv("VERSION", "v0.0.1")
	version := os.Getenv("VERSION")
	w.Header().Set("VERSION", version)
	fmt.Printf("os.version: %s\n", version)
	// 将request中的header设置到response中
	for k, v := range r.Header {
		for _, vv := range v {
			fmt.Printf("Header key: %s, Header value: %s \n", k, v)
			w.Header().Set(k, vv)
		}
	}
	// 输出日志
	clientIP := GetCurrentIP(r)
	log.Printf("Success! Response code: %d", 200)
	log.Printf("Success! Client IP: %s", clientIP)
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "working")
}

func GetCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-Real-IP")
	if ip == "" {
		ip = strings.Split(r.RemoteAddr, ":")[0]
	}
	return ip
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/debug/pprof/", pprof.Index)
	mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatalf("Start http server failed, error: %s\n", err.Error())
	}
}
