package gowdns

import (
	"errors"
	"strconv"
)

func AddRecord(d *DNSRecord) error {
	if d == nil {
		return errors.New("No DNS record provided.")
	}

	if !validateZoneName(d.Zone) {
		return errors.New("Invalid Zone format.")
	}

	if !validateNodeName(d.Node) {
		return errors.New("Invalid Node format.")
	}

	if !validateIPAddress(d.IP) {
		return errors.New("Invalid IP address format.")
	}

	return addRecord(d)
}

func addRecord(d *DNSRecord) error {
	ec := ExecutableConfig{
		Executable: "dnscmd",
		Flags: []string{
			"/recordadd",
			d.Zone,
			d.Node,
		},
	}

	if d.PTR {
		ec.Flags = append(ec.Flags, "/CreatePTR")
	}

	ttlToInt, err := strconv.Atoi(d.TTL)
	if err != nil {
		return err
	}

	if ttlToInt > 0 {
		ec.Flags = append(ec.Flags, d.TTL)
	}

	ec.Flags = append(ec.Flags, d.Type)
	ec.Flags = append(ec.Flags, d.IP)

	err = executeCommand(&ec)
	if err != nil {
		return err
	}

	return nil
}
