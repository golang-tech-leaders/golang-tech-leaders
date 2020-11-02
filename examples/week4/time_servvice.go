package main

import (
	"fmt"
	"net/http"
	"time"
)

// func greet(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprint(w, time.Now().Format(time.RFC1123))
// }

func main() {
	http.HandleFunc("/", timeHandler(now))
	http.ListenAndServe(":8080", nil)
}

type currentTime func() time.Time

func now() time.Time {
	return time.Now()
}

func timeHandler(t currentTime) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, t().Format(time.RFC1123))
	}
}

type database interface {
	//......
}

type handler struct {
	now currentTime
	db  database
}

func (h handler) whatTimeIsIt(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, h.now().Format(time.RFC1123))
}
