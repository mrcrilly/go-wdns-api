package main

import (
	"fmt"
	"net/http"
	"os"
)

import (
	"github.com/gorilla/mux"
	dns "github.com/mrcrilly/go-wdns-api"
)

func addDnsRecordHttpHandler(response http.ResponseWriter, request *http.Request) {
	variables := mux.Vars(request)

	record := dns.DNSRecord{
		Type: variables["type"],
		Zone: variables["zone"],
		Node: variables["node"],
		IP:   variables["value"],
		TTL:  "300",
		PTR:  true,
	}

	fmt.Fprintf(os.Stdout, "Request for: %+v\n", record)

	err := dns.AddRecord(&record)
	if err != nil {
		panic(err)
	}
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/v1/dns/add/{zone}/{node}/{type}/{value}", addDnsRecordHttpHandler)
	http.Handle("/", r)

	fmt.Println("Starting web server on 8080...")
	err := http.ListenAndServe("0.0.0.0:8080", nil)
	if err != nil {
		panic(err)
	}

}
