package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/route53"
	"log"
	"net"
	"time"
)

func DNSMode(s RemapSettings) {

	log.Printf("DNS Mode running")

	if s.RunOnce {
		_, err := DNSCheckAndUpdate(s)
		if err != nil {
			log.Fatal(err)
		}
	} else {

		for {

			SLEEP_TIME := s.Interval

			updated, err := DNSCheckAndUpdate(s)
			if err != nil {
				log.Print(err)
			}

			if updated {
				SLEEP_TIME += s.TTL
			}

			time.Sleep(time.Duration(SLEEP_TIME) * time.Second)
		}

	}

}

func DNSCheckAndUpdate(s RemapSettings) (bool, error) {

	updated := false

	myIP, err := getCurrentIP(s.UsePublicIP)
	if err != nil {
		return false, err
	}

	addrs, err := resolveName(s.DNSName)
	if err != nil {
		return false, err
	}

	log.Printf("%v %v", myIP, addrs)

	match := addressInAddresses(myIP, addrs)
	if !match {
		if err := updateDNS(s.HostedZoneID, s.DNSName, myIP, s.TTL); err != nil {
			log.Printf("Error updating DNS. IAM Role setup?\n%s", err)
			return false, err
		}
		updated = true
	}

	return updated, nil

}

func getCurrentIP(usePublicIP bool) (string, error) {

	myIP := ""
	var err error

	if usePublicIP {
		myIP, err = GetPublicIP()
	} else {
		myIP, err = GetPrivateIP()
	}

	return myIP, err

}

func updateDNS(hostedZoneID, dnsName, ip string, ttl int64) error {

	svc := route53.New(nil)

	params := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String(hostedZoneID),
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("UPSERT"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						TTL:  aws.Int64(ttl),
						Name: aws.String(dnsName),
						Type: aws.String("A"),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(ip),
							},
						},
					},
				},
			},
		},
	}

	resp, err := svc.ChangeResourceRecordSets(params)

	if err != nil {
		if awsErr, ok := err.(awserr.Error); ok {
			// Generic AWS error with Code, Message, and original error (if any)
			fmt.Println(awsErr.Code(), awsErr.Message(), awsErr.OrigErr())
			if reqErr, ok := err.(awserr.RequestFailure); ok {
				// A service error occurred
				fmt.Println(reqErr.Code(), reqErr.Message(), reqErr.StatusCode(), reqErr.RequestID())
			}
		} else {
			// This case should never be hit, the SDK should always return an
			// error which satisfies the awserr.Error interface.
			fmt.Println(err.Error())
		}
	}

	// Pretty-print the response data.
	fmt.Println(awsutil.Prettify(resp))
	return err
}

func resolveName(name string) ([]string, error) {

	var addresses []string
	addrs, err := net.LookupIP(name)
	if err != nil {
		return nil, err
	}

	for _, addr := range addrs {
		addresses = append(addresses, addr.String())
	}

	return addresses, nil

}

func addressInAddresses(addr string, addresses []string) bool {

	for _, a := range addresses {
		if addr == a {
			return true
		}
	}
	return false
}
