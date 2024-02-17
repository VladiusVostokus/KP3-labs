package main

import (
	"net/http"
)

func main(){
	http.HandleFunc("/", api)

	server := &http.Server{
		Addr: ":8795",
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}

func api(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("AAAAA"))
}