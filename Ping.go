package main

import (
	"fmt"
	influxClient "github.com/influxdb/influxdb/client"
	"log"
	"time"
)


func main()
{
	conf := influxClient.ClientConfig{
		Host:     config.Host,
		Database: config.Database,
		Username: config.Username,
		Password: config.Password,
	}

	con, err := client.NewClient(conf)
    if err != nil {
        log.Fatal(err)
    }

    

    if err != nil {
        log.Fatal(err)
    }

    dur, ver, err := con.Ping()
    if err != nil {
        log.Fatal(err)
    }
    log.Printf("Ping works! %v, %s", dur, ver)
    
}