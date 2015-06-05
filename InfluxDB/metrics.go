package main

import (
	"fmt"
	"log"
	"net/url"
	"time"

	"github.com/influxdb/influxdb/client"
)

func WriteMetricsDuration(label string, durationNs int64) error {
	l
	u, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		log.Fatal(err)
		return err
	}
	conf := client.Config{
		URL:      *u,
		Username: "root",
		Password: "root",
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var pts = make([]client.Point, 1)

	pts[0] = client.Point{
		Name: label,
		Tags: map[string]string{
			"region": "Singapore",
			"server": "01",
		},
		Fields: map[string]interface{}{
			"value": durationNs,
		},
		Time:      time.Now(),
		Precision: "n",
	}

	bps := client.BatchPoints{
		Points:   pts,
		Database: "techops",
		Time:     time.Now(),
	}
	_, err = con.Write(bps)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}

func WriteMetricsSize(label string, size int64) error {

	u, err := url.Parse(fmt.Sprintf("http://%s:%d", "localhost", 8086))
	if err != nil {
		log.Fatal(err)
		return err
	}
	conf := client.Config{
		URL:      *u,
		Username: "root",
		Password: "root",
	}

	con, err := client.NewClient(conf)
	if err != nil {
		log.Fatal(err)
		return err
	}
	var pts = make([]client.Point, 1)

	pts[0] = client.Point{
		Name: label,
		Tags: map[string]string{
			"region": "Singapore",
			"server": "01",
		},
		Fields: map[string]interface{}{
			"value": size,
		},
		Time:      time.Now(),
		Precision: "n",
	}

	bps := client.BatchPoints{
		Points:   pts,
		Database: "techops",
		Time:     time.Now(),
	}
	_, err = con.Write(bps)
	if err != nil {
		log.Fatal(err)
		return err
	}
	return err
}
