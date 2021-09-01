package main

import (
	"net/http"
	"log"
	"fmt"
	"os"
	"io"
	"context"
	"net"
	"github.com/noocsharp/go-aquos"
)

var client *aquos.Client

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		f, err := os.Open("client.html")
		if err != nil {
			fmt.Fprintf(w, "couldn't load file\n")
		}

		_, err := io.Copy(w, f)
		if err != nil {
			fmt.Fprintf(w, "couldn't copy file\n")
		}

		f.Close()
	} else if r.Method == "POST" {
		if r.URL.Path == "/mute" {
			client.MuteToggle()
		}
	}
}

func main() {
	http.HandleFunc("/", handler)

	client = &aquos.Client{}
	err := client.Connect(context.Background(), net.JoinHostPort("192.168.1.195", "10002"))
	if err != nil {
		log.Fatal("failed to connect to tv\n")
	}


	log.Fatal(http.ListenAndServe(":8080", nil))

	client.Close()
}
