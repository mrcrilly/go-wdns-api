package main

import (
	"fmt"
	"os"
)

import (
	dns "github.com/mrcrilly/go-wdns-api"
)

func main() {
	fmt.Fprint(os.Stdout, "Building DNS record...\n")
	record := dns.DNSRecord{
		Type: "A",
		Zone: "something.com",
		Node: "example-a-record",
		IP:   "10.1.1.1",
		TTL:  "300",
		PTR:  true,
	}

	fmt.Fprintf(os.Stdout, "Executing command: dnscmd /recordadd %s %s /CreatePTR %s %s %s\n", record.Zone, record.Node, record.TTL, record.Type, record.IP)
	result := dns.AddRecord(&record)
	if result != nil {
		fmt.Fprintf(os.Stderr, "error: %s\n", result)
		os.Exit(1)
	}

	fmt.Fprint(os.Stdout, "Finished building DNS record...\n")
	fmt.Fprintf(os.Stdout, "%s", result)
}
