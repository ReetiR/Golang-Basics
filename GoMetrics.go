package main

import (
	"log"
	"os"

	"github.com/rcrowley/go-metrics"
	"github.com/rcrowley/go-metrics/influxdb"
)

func main() {
	c := metrics.NewCounter()
	metrics.Register("foo", c)
	c.Inc(47)

	g := metrics.NewGauge()
	metrics.Register("bar", g)
	g.Update(47)

	s := metrics.NewExpDecaySample(1028, 0.015) // or metrics.NewUniformSample(1028)
	h := metrics.NewHistogram(s)
	metrics.Register("baz", h)
	h.Update(47)

	m := metrics.NewMeter()
	metrics.Register("quux", m)
	m.Mark(47)

	t := metrics.NewTimer()
	metrics.Register("bang", t)
	t.Time(func() {})
	t.Update(47)

	go metrics.Log(metrics.DefaultRegistry, 60e9, log.New(os.Stderr, "metrics: ", log.Lmicroseconds))

	influxdb.Influxdb(metrics.DefaultRegistry, 60e9, &influxdb.Config{
		Host:     "127.0.0.1:8086",
		Database: "blackbuck",
		Username: "root",
		Password: "root",
	})
}
