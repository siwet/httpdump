package main

import (
	"net"
	"net/http"
	"os"
	"testing"
)

func TestAsdf(t *testing.T) {
	os.Setenv("VERSION", "utest123")
	initHttp()
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		t.Fatal(err)
	}
	go func() {
		err := http.Serve(listener, nil)
		if err != nil {
			panic(err)
		}
	}()

	t.Run("test1_writeHeader", func(t *testing.T) {
		req, err := http.NewRequest("GET", "http://127.0.0.1:8080/xxx", nil)
		if err != nil {
			t.Fatal(err)
		}
		const (
			H = "UTestHeader1"
			V = "UTestValue1"
		)
		req.Header.Set(H, V)
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			t.Fatal(err)
		}
		if resp.Header.Get(H) != V {
			t.Fatal("test failed! ", resp.Header.Get(H))
		}
	})
	t.Run("test2_readEnvVersion", func(t *testing.T) {
		resp, err := http.Get("http://127.0.0.1:8080/xxx")
		if err != nil {
			t.Fatal(err)
		}
		if resp.Header.Get("VERSION") != "utest123" {
			t.Fatal("test failed! ", resp.Header.Get("VERSION"))
		}
	})
	t.Run("test3_log", func(t *testing.T) {
	})
	t.Run("test4_healthz", func(t *testing.T) {
		resp, err := http.Get("http://127.0.0.1:8080/healthz")
		if err != nil {
			t.Fatal(err)
		}
		if resp.StatusCode != 200 {
			t.Fatal()
		}
	})
}
