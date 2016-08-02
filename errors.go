package gowdns

import (
	"errors"
	"regexp"
)

type RegexpTable struct {
	Pattern       string
	HumanResponse string
}

var errorTable = []RegexpTable{
	{
		Pattern:       `DNS_ERROR_ZONE_DOES_NOT_EXIST[ ]+?9601`,
		HumanResponse: "(Error) The DNS zone you provided doesn't exist. Review your request.",
	},
	{
		Pattern:       `DNS_WARNING_PTR_CREATE_FAILED[ ]+?9715`,
		HumanResponse: "(Warning) Unable to create PTR. Does the record already exist? Does the PTR zone exist for the subnet of the IP?",
	},
	{
		Pattern:       `DNS_ERROR_RECORD_ALREADY_EXISTS[ ]+?9711`,
		HumanResponse: "(Error) The record already exists. Try again if you believe it's a cache issue.",
	},
	{
		Pattern:       `DNS_ERROR_RECORD_ALREADY_EXISTS[ ]+?9711`,
		HumanResponse: "(Error) The record already exists. Try again if you believe it's a cache issue.",
	},
}

func parseDNSCmdError(so, se string, err error) error {
	if so != "" {
		for _, e := range errorTable {
			matcher := regexp.MustCompile(e.Pattern)
			if matcher.MatchString(so) {
				return errors.New(e.HumanResponse)
			}
		}

		return errors.New("There was an error, but we don't have a match for it. Raw output: " + so)
	} else {
		return errors.New("Error reported, but SO is empty. Nothing to report, sorry.")
	}

	return nil
}
