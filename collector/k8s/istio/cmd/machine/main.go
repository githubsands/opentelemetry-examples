package main

import (
	"errors"
	"flag"
	"fmt"
	"net"
	"os"
	"sync"
	"time"
)

var defaultPeerRedialFrequency time.Duration = 2

var peerRedialFrequency = flag.Int("peerRedialFrequency", 30, "the time in second to retry a peer")
var peerURLOpt = flag.String("peerURL", "", "write peer url")
var listenURLOpt = flag.String("listenURL", "", "url to listen on")

func main() {
	flag.Parse()
	peerRedialFrequencyDuration := defaultPeerRedialFrequency * time.Second
	if *peerRedialFrequency != 0 {
		peerRedialFrequencyDuration = time.Duration(*peerRedialFrequency) * time.Second
	}
	fmt.Println("Starting machine")
	go func() {
		loopCount := 0
		for {
			loopCount++
			fmt.Println("Loopcount:", loopCount)
			time.Sleep(10000000000)
		}
	}()
	peerURL, ok := os.LookupEnv("PEER_URL")
	listenURL, ok := os.LookupEnv("LISTEN_URL")
	if *peerURLOpt != "" {
		peerURL = *peerURLOpt
	}
	if *listenURLOpt != "" {
		listenURL = *listenURLOpt
	}
	if !ok {
		fmt.Println("ERROR: you must set environmental variables: PEER_URL && LISTEN_URL")
		os.Exit(1)
	}
	errs := make(chan error, 1)
	ln, err := net.Listen("tcp", listenURL)
	if err != nil {
		errs <- err
	}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		for {
			conn, err := ln.Accept()
			if err != nil {
				errs <- err
			}
			fmt.Println("Accepted")
			var bytes []byte
			size, err := conn.Read(bytes)
			if err != nil {
				errs <- err
			}
			fmt.Println("Received message: ", string(bytes))
			if size > 0 {
				fmt.Println("Received:", string(bytes))
			}
		}
	}()
	wg.Add(1)
	go func() {
		peerRecount := 0
		connected := false
		var conn net.Conn
		for {
			if peerRecount <= 5 && connected == false {
				conn, err = net.Dial("tcp", peerURL)
				if err != nil {
					fmt.Println("Unable to dial peer redialing, ", peerRecount)
					time.Sleep(peerRedialFrequencyDuration)
					peerRecount++
					continue
				}
				connected = true
				fmt.Println("Successfully dialed: ", conn.RemoteAddr())
			}
			if peerRecount >= 5 {
				errs <- errors.New("Failed to connect to peer")
				return
			}
			writeFrequency := 2 * time.Second
			time.Sleep(writeFrequency)
			fmt.Println("Writing msg")
			conn.Write([]byte("connected"))
			defer conn.Close()
		}
	}()
	for {
		select {
		case err := <-errs:
			fmt.Println(err.Error(), "\nexiting program")
			os.Exit(1)
		}
	}
}
