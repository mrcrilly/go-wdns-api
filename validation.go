package gowdns

import (
	"net"
	"regexp"
)

var validZoneName = regexp.MustCompile(`^[A-Za-z0-9\.-]+$`)
var validNodeName = validZoneName

func validateZoneName(z string) bool {
	return validZoneName.MatchString(z)
}

func validateNodeName(n string) bool {
	return validNodeName.MatchString(n)
}

func validateIPAddress(i string) bool {
	validIP := net.ParseIP(i + "/32")
	return validIP == nil
}
