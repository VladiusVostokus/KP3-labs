package main

import (
	"encoding/json"
	"net/http"
	"time"
)

var now map[string]string

func main() {
	http.HandleFunc("/time", api)

	server := &http.Server{
		Addr: ":8795",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	now = make(map[string]string)
	currentTime := time.Now()
	formattedTime := currentTime.Format("2006-01-02 15:04:05")
	now["time"] = string(formattedTime)
	barr, _ := json.Marshal(now)
	w.Write(barr)
}
