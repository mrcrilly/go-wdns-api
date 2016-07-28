package gowdns

import (
	"bytes"
	"os"
)

type DNSRecord struct {
	Type string
	Zone string
	Node string
	IP   string
	TTL  string
	PTR  bool
}

type ExecutableConfig struct {
	Executable string
	Flags      []string
	Stdout     *os.File
	Stdin      *bytes.Buffer
	Stderr     *bytes.Buffer
}
