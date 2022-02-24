package main

import (
	"fmt"
	"log"
	"net/http"
)

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "200")
}

//func rootHandler(w http.ResponseWriter, r *http.Request) {
//	fmt.Println()
//}

func main() {
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":80", nil)
	http.ListenAndServe(":80", nil)
	if err != nil {
		log.Fatal(err)
	}
	//flag.Set("v", "4")
	//glog.V(2).Info("Starting http server...")
	//mux := http.NewServeMux()
	//mux.HandleFunc("/", rootHandler)
	//err := http.ListenAndServe(":80", mux)
	//if err != nil {
	//	log.Fatal(err)
	//}
}
