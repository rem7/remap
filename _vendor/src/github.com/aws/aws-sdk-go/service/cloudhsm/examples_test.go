// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

package cloudhsm_test

import (
	"bytes"
	"fmt"
	"time"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/aws/awsutil"
	"github.com/aws/aws-sdk-go/service/cloudhsm"
)

var _ time.Duration
var _ bytes.Buffer

func ExampleCloudHSM_CreateHapg() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.CreateHapgInput{
		Label: aws.String("Label"), // Required
	}
	resp, err := svc.CreateHapg(params)

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
}

func ExampleCloudHSM_CreateHsm() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.CreateHsmInput{
		IamRoleArn:       aws.String("IamRoleArn"),       // Required
		SshKey:           aws.String("SshKey"),           // Required
		SubnetId:         aws.String("SubnetId"),         // Required
		SubscriptionType: aws.String("SubscriptionType"), // Required
		ClientToken:      aws.String("ClientToken"),
		EniIp:            aws.String("IpAddress"),
		ExternalId:       aws.String("ExternalId"),
		SyslogIp:         aws.String("IpAddress"),
	}
	resp, err := svc.CreateHsm(params)

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
}

func ExampleCloudHSM_CreateLunaClient() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.CreateLunaClientInput{
		Certificate: aws.String("Certificate"), // Required
		Label:       aws.String("ClientLabel"),
	}
	resp, err := svc.CreateLunaClient(params)

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
}

func ExampleCloudHSM_DeleteHapg() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DeleteHapgInput{
		HapgArn: aws.String("HapgArn"), // Required
	}
	resp, err := svc.DeleteHapg(params)

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
}

func ExampleCloudHSM_DeleteHsm() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DeleteHsmInput{
		HsmArn: aws.String("HsmArn"), // Required
	}
	resp, err := svc.DeleteHsm(params)

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
}

func ExampleCloudHSM_DeleteLunaClient() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DeleteLunaClientInput{
		ClientArn: aws.String("ClientArn"), // Required
	}
	resp, err := svc.DeleteLunaClient(params)

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
}

func ExampleCloudHSM_DescribeHapg() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DescribeHapgInput{
		HapgArn: aws.String("HapgArn"), // Required
	}
	resp, err := svc.DescribeHapg(params)

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
}

func ExampleCloudHSM_DescribeHsm() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DescribeHsmInput{
		HsmArn:          aws.String("HsmArn"),
		HsmSerialNumber: aws.String("HsmSerialNumber"),
	}
	resp, err := svc.DescribeHsm(params)

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
}

func ExampleCloudHSM_DescribeLunaClient() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.DescribeLunaClientInput{
		CertificateFingerprint: aws.String("CertificateFingerprint"),
		ClientArn:              aws.String("ClientArn"),
	}
	resp, err := svc.DescribeLunaClient(params)

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
}

func ExampleCloudHSM_GetConfig() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.GetConfigInput{
		ClientArn:     aws.String("ClientArn"),     // Required
		ClientVersion: aws.String("ClientVersion"), // Required
		HapgList: []*string{ // Required
			aws.String("HapgArn"), // Required
			// More values...
		},
	}
	resp, err := svc.GetConfig(params)

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
}

func ExampleCloudHSM_ListAvailableZones() {
	svc := cloudhsm.New(nil)

	var params *cloudhsm.ListAvailableZonesInput
	resp, err := svc.ListAvailableZones(params)

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
}

func ExampleCloudHSM_ListHapgs() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ListHapgsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListHapgs(params)

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
}

func ExampleCloudHSM_ListHsms() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ListHsmsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListHsms(params)

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
}

func ExampleCloudHSM_ListLunaClients() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ListLunaClientsInput{
		NextToken: aws.String("PaginationToken"),
	}
	resp, err := svc.ListLunaClients(params)

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
}

func ExampleCloudHSM_ModifyHapg() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ModifyHapgInput{
		HapgArn: aws.String("HapgArn"), // Required
		Label:   aws.String("Label"),
		PartitionSerialList: []*string{
			aws.String("PartitionSerial"), // Required
			// More values...
		},
	}
	resp, err := svc.ModifyHapg(params)

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
}

func ExampleCloudHSM_ModifyHsm() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ModifyHsmInput{
		HsmArn:     aws.String("HsmArn"), // Required
		EniIp:      aws.String("IpAddress"),
		ExternalId: aws.String("ExternalId"),
		IamRoleArn: aws.String("IamRoleArn"),
		SubnetId:   aws.String("SubnetId"),
		SyslogIp:   aws.String("IpAddress"),
	}
	resp, err := svc.ModifyHsm(params)

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
}

func ExampleCloudHSM_ModifyLunaClient() {
	svc := cloudhsm.New(nil)

	params := &cloudhsm.ModifyLunaClientInput{
		Certificate: aws.String("Certificate"), // Required
		ClientArn:   aws.String("ClientArn"),   // Required
	}
	resp, err := svc.ModifyLunaClient(params)

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
}
