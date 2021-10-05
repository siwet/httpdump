package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

type WrapperRespWriter struct {
	http.ResponseWriter
	Status int
}

func (r *WrapperRespWriter) WriteHeader(status int) {
	r.Status = status
	r.ResponseWriter.WriteHeader(status)
}

func logMiddleware(h http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		ww := &WrapperRespWriter{
			ResponseWriter: w,
			Status:         200,
		}
		h.ServeHTTP(ww, req)
		log.Printf("resp code: %d; remote client: %s\n", ww.Status, req.RemoteAddr)
	}
}

func initHttp() {
	http.HandleFunc("/", logMiddleware(func(w http.ResponseWriter, req *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		for k, v := range req.Header {
			for _, hs := range v {
				w.Header().Add(k, hs)
			}
		}
		w.Header().Set("VERSION", os.Getenv("VERSION"))
	}))
	http.HandleFunc("/healthz", logMiddleware(func(w http.ResponseWriter, request *http.Request) {
		w.Write([]byte("OK"))
	}))
}

func main() {
	initHttp()
	var port = 8080
	listenAddr := fmt.Sprintf("[::]:%d", port)
	fmt.Println("start http server, listen on: " + listenAddr)
	err := http.ListenAndServe(listenAddr, nil)
	if err != nil {
		panic(err)
	}
}
