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

func dnsLoop(settings RemapSettings) {

	dnsName := settings.DNSName
	usePublicIP := settings.UsePublicIP
	ttl := settings.TTL
	hostedZoneID := settings.HostedZoneID
	interval := settings.Interval

	for {

		SLEEP_TIME := 0
		myIP := ""

		if usePublicIP {
			log.Print("using-public-ip")
			publicIP, err := GetPublicIP()
			if err != nil {
				log.Printf("we dont have internet. keep looping until we do")
				time.Sleep(5 * time.Second)
				continue
			}
			myIP = publicIP
		} else {
			log.Print("using-local-ip")
			privateIP, err := GetPrivateIP()
			if err != nil {
				log.Printf("can't find eth0?")
				log.Fatal(err)
			}
			myIP = privateIP
		}

		log.Printf("IP: %s", myIP)

		addrs, err := resolveName(dnsName)
		if err != nil {
			log.Fatal(err)
		}

		log.Printf("DNS %s resolves to %s", dnsName, addrs)

		match := addressInAddresses(myIP, addrs)
		if match {
			log.Print("matched")
			SLEEP_TIME += interval
		} else {
			log.Print("not matched. updating DNS")
			SLEEP_TIME += interval + int(ttl)
			if err := updateDNS(hostedZoneID, dnsName, myIP, ttl); err != nil {
				log.Fatalf("Error updating DNS. IAM Role setup?\n%s", err)
			}
		}

		log.Printf("Sleeping for %v", SLEEP_TIME)
		time.Sleep(time.Duration(SLEEP_TIME) * time.Second)
	}

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
