package main

import (
	"net/http"
	"log"
	"fmt"
	"os"
	"io"
	"github.com/noocsharp/go-aquos"
)

var client *aquos.Client

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		f, err := os.Open("client.html")
		if err != nil {
			fmt.Fprintf(w, "couldn't load file\n")
		}

		_, err = io.Copy(w, f)
		if err != nil {
			fmt.Fprintf(w, "couldn't copy file\n")
		}

		f.Close()
	} else if r.Method == "POST" {
		if r.URL.Path == "/mute" {
			client.MuteToggle()
		} else if r.URL.Path == "/volup" {
			client.VolumeUp()
		} else if r.URL.Path == "/voldown" {
			client.VolumeDown()
		} else if r.URL.Path == "/chup" {
			client.ChannelUp()
		} else if r.URL.Path == "/chdown" {
			client.ChannelDown()
		} else if r.URL.Path == "/enter" {
			client.Enter()
		} else if r.URL.Path == "/up" {
			client.Up()
		} else if r.URL.Path == "/down" {
			client.Down()
		} else if r.URL.Path == "/left" {
			client.Left()
		} else if r.URL.Path == "/right" {
			client.Right()
		} else if r.URL.Path == "/return" {
			client.Return()
		} else if r.URL.Path == "/netflix" {
			client.Netflix()
		} else if r.URL.Path == "/input" {
			client.ToggleInput()
		} else if r.URL.Path == "/play" {
			client.Play()
		} else if r.URL.Path == "/pause" {
			client.Pause()
		} else if r.URL.Path == "/poweron" {
			client.Power(true)
		} else if r.URL.Path == "/poweroff" {
			client.Power(false)
		}
	}
}

func main() {
	http.HandleFunc("/", handler)

	client = &aquos.Client{}
	client.Address = "192.168.1.195:10002"

	log.Fatal(http.ListenAndServe(":8080", nil))

	client.Close()
}
