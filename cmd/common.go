package cmd

import (
	"github.com/DennyLoko/cfdnsmanager/dns"
	"github.com/cloudflare/cloudflare-go"
	"github.com/fatih/color"
)

var (
	cfAPI *cloudflare.API

	printNotice  func(...interface{})
	printError   func(...interface{})
	printSuccess func(...interface{})
)

func init() {
	printNotice = color.New(color.FgYellow).PrintlnFunc()
	printError = color.New(color.FgRed).PrintlnFunc()
	printSuccess = color.New(color.FgHiGreen).PrintlnFunc()
}

func getIP() (string, error) {
	dnsAPI := dns.NewDNS()

	ip, err := dnsAPI.OwnAddress()
	if err != nil {
		return "", err
	}

	return ip[0].String(), nil
}

func updateRecord(zoneID, recordID, content string, ttl int) error {
	record, err := cfAPI.DNSRecord(zoneID, recordID)
	if err != nil {
		return err
	}

	record.Content = content

	if newTTL != 0 {
		record.TTL = newTTL
	}

	err = cfAPI.UpdateDNSRecord(zoneID, recordID, record)
	if err != nil {
		return err
	}

	return nil
}
