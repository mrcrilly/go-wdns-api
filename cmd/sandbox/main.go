package main

import (
	"flag"
	"fmt"
	"os"
)

import (
	dns "github.com/mrcrilly/go-wdns-api"
)

var typeFromCli = flag.String("type", "A", "The record type to manage. Defaults to 'A'.")
var zoneFromCli = flag.String("zone", "", "The DNS zone to manage with this request.")
var nodeFromCli = flag.String("node", "", "The node name to manage with this request.")
var ipFromCli = flag.String("ip", "", "The IP to add or update the record with.")
var ttlFromCli = flag.String("ttl", "800", "The TTL value for the new record. Defaults to 800 seconds.")
var ptrFromCli = flag.Bool("ptr", true, "Whether or not to setup a PTR record. Defaults to true.")
var addRecordFromCli = flag.Bool("add", true, "Add the record to the zone. Must delete the record first!!")
var deleteRecordFromCli = flag.Bool("delete", false, "Delete the record from the zone.")

func checkError(e error) {
	if e != nil {
		fmt.Fprint(os.Stderr, e)
		os.Exit(1)
	}
}

func main() {
	flag.Parse()

	record := dns.DNSRecord{
		Type: *typeFromCli,
		Zone: *zoneFromCli,
		Node: *nodeFromCli,
		IP:   *ipFromCli,
		TTL:  *ttlFromCli,
		PTR:  *ptrFromCli,
	}

	if *deleteRecordFromCli {
		result := dns.DeleteRecord(&record)
		checkError(result)
		os.Exit(0)
	}

	result := dns.AddRecord(&record)
	checkError(result)
	os.Exit(0)
}
