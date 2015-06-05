package influxdb

import (
	"fmt"
	influxClient "github.com/GeertJohan/influxdb-client-old"
	"github.com/rcrowley/go-metrics"
	"github.com/influxdb/influxdb/influxql"
	"github.com/rcrowley/go-metrics"
	"log"
	"time"
	"net/url"
)

type Config struct {
	Host     string
	Database string
	Username string
	Password string
}

func Influxdb(r metrics.Registry, d time.Duration, config *Config) {
	u, _ := url.Parse(config.Host)
	client, err := influxClient.NewClient(influxClient.Config{
		URL:     *u,
		Username: config.Username,
		Password: config.Password,
	})
	if err != nil {
		log.Println(err)
		return
	}

	for _ = range time.Tick(d) {
		if err := send(r, client); err != nil {
			log.Println(err)
		}
	}
}

func send(r metrics.Registry, client *influxClient.Client) error {
	//series := []*influxClient.Series{}
	series := make(influxql.Rows, 0)
	r.Each(func(name string, i interface{}) {
		now := getCurrentTime()
		switch metric := i.(type) {
		case metrics.Counter:
			oneRow := &influxql.Row{
				Name:    fmt.Sprintf("%s.count", name),
				Columns: []string{"time", "count"},
				Values: [][]interface{}{
					{now, metric.Count()},
				},
		}
			series = append(series, oneRow)
		case metrics.Gauge:
			series = append(series, &influxql.Row{
				Name:    fmt.Sprintf("%s.value", name),
				Columns: []string{"time", "value"},
				Values: [][]interface{}{
					{now, metric.Value()},
				},
			})
		case metrics.GaugeFloat64:
			series = append(series, &influxql.Row{
				Name:    fmt.Sprintf("%s.value", name),
				Columns: []string{"time", "value"},
				Values: [][]interface{}{
					{now, metric.Value()},
				},
			})
		case metrics.Histogram:
			h := metric.Snapshot()
			ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			series = append(series, &influxql.Row{
				Name: fmt.Sprintf("%s.histogram", name),
				Columns: []string{"time", "count", "min", "max", "mean", "std-dev",
					"50-percentile", "75-percentile", "95-percentile",
					"99-percentile", "999-percentile"},
				Values: [][]interface{}{
					{now, h.Count(), h.Min(), h.Max(), h.Mean(), h.StdDev(),
						ps[0], ps[1], ps[2], ps[3], ps[4]},
				},
			})
		case metrics.Meter:
			m := metric.Snapshot()
			series = append(series, &influxql.Row{
				Name: fmt.Sprintf("%s.meter", name),
				Columns: []string{"count", "one-minute",
					"five-minute", "fifteen-minute", "mean"},
				Values: [][]interface{}{
					{m.Count(), m.Rate1(), m.Rate5(), m.Rate15(), m.RateMean()},
				},
			})
		case metrics.Timer:
			h := metric.Snapshot()
			ps := h.Percentiles([]float64{0.5, 0.75, 0.95, 0.99, 0.999})
			series = append(series, &influxql.Row{
				Name: fmt.Sprintf("%s.timer", name),
				Columns: []string{"count", "min", "max", "mean", "std-dev",
					"50-percentile", "75-percentile", "95-percentile",
					"99-percentile", "999-percentile", "one-minute", "five-minute", "fifteen-minute", "mean-rate"},
				Values: [][]interface{}{
					{h.Count(), h.Min(), h.Max(), h.Mean(), h.StdDev(),
						ps[0], ps[1], ps[2], ps[3], ps[4],
						h.Rate1(), h.Rate5(), h.Rate15(), h.RateMean()},
				},
			})
		}
	})
	if err := client.WriteSeries(series); err != nil {
		log.Println(err)
	}
	return nil
}

func getCurrentTime() int64 {
	return time.Now().UnixNano() / 1000000
}
