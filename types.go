package gowdns

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
}
