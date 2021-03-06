package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"sort"
)

func init() {
	log.SetOutput(os.Stdout)
}

func handler(w http.ResponseWriter, r *http.Request) {
	log.Println("Got", r.Method, "to:", r.URL)

	var keys []string
	for k := range r.Header {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	log.Println("Request Headers:")
	for _, k := range keys {
		log.Println("  ", k, ":", r.Header[k])
	}

	log.Println("Cookies:")
	log.Println(r.Cookies())

	log.Println("Body:")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		panic(-1)
	}
	log.Println(string(body))
	log.Println("---End Body---")

	fmt.Fprintln(w, "Logged request.")
}

func main() {
	http.HandleFunc("/", handler)

	port := "6000"
	if os.Getenv("PORT") != "" {
		port = os.Getenv("PORT")
	}
	log.Println("Listening on port", port, "...")
	http.ListenAndServe(":"+port, nil)
}
