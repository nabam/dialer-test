package main

import (
	"context"
	"log"
	"net"
	"os"
	"strconv"
	"time"
)

func main() {
	interval, err := strconv.Atoi(os.Getenv("DIALER_INTERVAL_SEC"))
	target := os.Getenv("DIALER_TARGET")
	if err != nil {
		log.Fatalf(err.Error())
	}
	for {
		doDial(target)
		time.Sleep(time.Duration(interval) * time.Second)
	}
}

func doDial(target string) error {
	dialer := &net.Dialer{}

	c, err := dialer.DialContext(context.Background(), "tcp", target)
	if err != nil {
		log.Printf(err.Error())
		return err
	}
	defer c.Close()
	// log.Printf("connected to %s", target)
	return nil
}
