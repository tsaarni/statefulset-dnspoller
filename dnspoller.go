package main

import (
	"fmt"
	"log/slog"
	"net"
	"os"
	"time"
)

func lookup(host string) ([]string, error) {
	ips, err := net.LookupIP(host)
	if err != nil {
		return nil, err
	}
	var addresses []string
	for _, ip := range ips {
		addresses = append(addresses, ip.String())
	}
	return addresses, nil
}

func contains(addresses []string, address string) bool {
	for _, a := range addresses {
		if a == address {
			return true
		}
	}
	return false
}

func printDifferences(prevAddresses, currentAddresses []string) {
	for _, address := range prevAddresses {
		if !contains(currentAddresses, address) {
			slog.Info("Removed", "address", address)
		}
	}
	for _, address := range currentAddresses {
		if !contains(prevAddresses, address) {
			slog.Info("Added", "address", address)
		}
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: dnspoller <hostname>")
		return
	}
	host := os.Args[1]

	// Replace the default logger to get msec timestamps.
	slog.SetDefault(slog.New(slog.NewTextHandler(os.Stdout, nil)))

	// Log own IP address.
	addrs, err := net.InterfaceAddrs()
	if err != nil {
		slog.Error("Error getting interface addresses", "error", err)
	} else {
		var ownAddresses []string
		for _, addr := range addrs {
			ownAddresses = append(ownAddresses, addr.String())
		}
		slog.Info("Own addresses", "addresses", ownAddresses)
	}

	// Write pid file for implementing readiness probe.
	pid := os.Getpid()
	pidFile, err := os.Create("dnspoller.pid")
	if err != nil {
		slog.Error("Error creating pid file", "error", err)
		return
	}
	defer pidFile.Close()
	fmt.Fprintf(pidFile, "%d\n", pid)

	slog.Info("Lookup", "hostname", host)

	var prevAddresses []string
	for {
		currentAddresses, err := lookup(host)
		if err != nil {
			slog.Error("Error looking up", "hostname", host, "error", err)
			time.Sleep(1 * time.Second)
			continue
		}

		printDifferences(prevAddresses, currentAddresses)
		prevAddresses = currentAddresses

		time.Sleep(1 * time.Second)
	}
}
