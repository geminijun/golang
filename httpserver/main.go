package main

import (
	"flag"
	"fmt"
	"os"
	"io"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/golang/glog"
)

func main() {
	flag.Set("v", "4")
	glog.V(2).Info("Starting http server...")
	// http.HandleFunc("/", rootHandler)
	// err := http.ListenAndServe(":80", nil)
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/healthz", healthz)
	// mux.HandleFunc("/debug/pprof/", pprof.Index)
	// mux.HandleFunc("/debug/pprof/profile", pprof.Profile)
	// mux.HandleFunc("/debug/pprof/symbol", pprof.Symbol)
	// mux.HandleFunc("/debug/pprof/trace", pprof.Trace)
	err := http.ListenAndServe(":80", mux)
	if err != nil {
		log.Fatal(err)
	}

}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering healthz handler")
	w.WriteHeader(http.StatusOK)
	io.WriteString(w, "ok\n")
}

func GetIP(r *http.Request) string {
	forwarded := r.Header.Get("X-FORWARDED-FOR")
	if forwarded != "" {
		return forwarded
	}
	return r.RemoteAddr
}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("entering root handler")
	ip := GetIP(r)
	if ip != "" {
		fmt.Println("IP: ", ip)
	} else {
		fmt.Println("Can't find client's ip!\n")
	}
	version := os.Getenv("VERSION")
	if version != "" {
		io.WriteString(w, fmt.Sprintf("VERSION: [%s]\n", version))
	} else {
		io.WriteString(w, "Can't find version\n")
	}
	io.WriteString(w, "===================Details of the http request header:============\n")
	for k, v := range r.Header {
		io.WriteString(w, fmt.Sprintf("%s=%s\n", k, v))
	}
}
