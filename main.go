package main

import (
	"encoding/json"
	"net/http"
	"time"
)

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
	now := make(map[string]string)
	currentTime := time.Now()
	formattedTime := currentTime.Format(time.RFC3339)
	now["time"] = string(formattedTime)
	barr, err := json.Marshal(now)

	if err != nil {
		panic(err)
	}
	
	w.Write(barr)
}
