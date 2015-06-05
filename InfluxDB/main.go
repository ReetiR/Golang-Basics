package main

import (
	"fmt"
	"io/ioutil"
	"log"
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
	err := WriteMetricsDuration("responseTime", elapsed.Nanoseconds())
	if err != nil {
		log.Printf("%s", err)
	}
	err = WriteMetricsSize("bytesProcessed", int64(n))
	if err != nil {
		log.Printf("%s", err)
	}
}

func main() {

	http.HandleFunc("/", viewHandler)
	http.ListenAndServe(":8000", nil)
}
