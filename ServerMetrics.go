package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func loadPage(rpath string) (body []byte, err error) {
	body, err = ioutil.ReadFile(rpath)
	if err != nil {
		return nil, err
	}
	return body, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	rn := r.URL.Path[len("/"):]
	b, _ := loadPage(rn)
	n, _ := w.Write(b)
	elapsed := time.Since(start)
	fmt.Println("elapsed ", elapsed.Seconds(), " secs for processing bytes", n)
}

func main() {

	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8000", nil)
}
