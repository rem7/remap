package main

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/ec2"
	"log"
	"time"
)

const (
	DELAY                     = 30
	REQ_TIMEOUT time.Duration = 5
)

func EIPMode(s RemapSettings) {

	region, err := getRegion()
	if err != nil {
		log.Fatal(err)
	}

	instanceId, err := getInstanceId()
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Running on EIP mode")
	log.Printf("Region: %s", region)
	log.Printf("Instance: %s", instanceId)

	if s.RunOnce {
		_, err := updateIP(s, region, instanceId)
		if err != nil {
			log.Printf("Error updating IP.")
			log.Fatal(err)
		}
	} else {

		for {

			SLEEP_TIME := s.Interval

			updated, err := updateIP(s, region, instanceId)
			if err != nil {
				log.Printf("Error updating IP: %v", err)
				log.Printf("trying again on next loop")
			}

			if updated {
				SLEEP_TIME = s.Interval + DELAY
			}

			time.Sleep(time.Duration(SLEEP_TIME) * time.Second)
		}
	}

}

func updateIP(s RemapSettings, region, instanceId string) (bool, error) {

	updated := false
	eip := s.Eip
	eipAllocationId := s.EipAllocationId

	ip, err := GetPublicIP()
	if err != nil {
		log.Printf("error getting public IP: %+v", err)
		return false, err
	}

	if !eipMatches(ip, eip) {
		log.Printf("My IP: %s", ip)
		log.Printf("Public IP and EIP don't match. Stealing EIP. Assigning %s to %s", eip, instanceId)
		err := stealIp(eipAllocationId, instanceId, region)
		if err != nil {
			return false, err
		}
		updated = true
		log.Printf("Updated EIP allocation succesfully")
	}

	return updated, nil

}

func stealIp(eipAllocationId, instanceId, region string) error {

	config := &aws.Config{Region: aws.String(region)}
	svc := ec2.New(config)

	params := &ec2.AssociateAddressInput{
		AllocationId:       aws.String(eipAllocationId),
		AllowReassociation: aws.Bool(true),
		DryRun:             aws.Bool(false),
		InstanceId:         aws.String(instanceId),
	}

	log.Printf("Associating address...")
	resp, err := svc.AssociateAddress(params)

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
		return err
	}

	fmt.Println(awsutil.Prettify(resp))
	return nil
}

func eipMatches(ip1, ip2 string) bool {
	return ip1 == ip2
}
