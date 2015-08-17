// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package configserviceiface provides an interface for the AWS Config.
package configserviceiface

import (
	"github.com/aws/aws-sdk-go/aws/service"
	"github.com/aws/aws-sdk-go/service/configservice"
)

// ConfigServiceAPI is the interface type for configservice.ConfigService.
type ConfigServiceAPI interface {
	DeleteDeliveryChannelRequest(*configservice.DeleteDeliveryChannelInput) (*service.Request, *configservice.DeleteDeliveryChannelOutput)

	DeleteDeliveryChannel(*configservice.DeleteDeliveryChannelInput) (*configservice.DeleteDeliveryChannelOutput, error)

	DeliverConfigSnapshotRequest(*configservice.DeliverConfigSnapshotInput) (*service.Request, *configservice.DeliverConfigSnapshotOutput)

	DeliverConfigSnapshot(*configservice.DeliverConfigSnapshotInput) (*configservice.DeliverConfigSnapshotOutput, error)

	DescribeConfigurationRecorderStatusRequest(*configservice.DescribeConfigurationRecorderStatusInput) (*service.Request, *configservice.DescribeConfigurationRecorderStatusOutput)

	DescribeConfigurationRecorderStatus(*configservice.DescribeConfigurationRecorderStatusInput) (*configservice.DescribeConfigurationRecorderStatusOutput, error)

	DescribeConfigurationRecordersRequest(*configservice.DescribeConfigurationRecordersInput) (*service.Request, *configservice.DescribeConfigurationRecordersOutput)

	DescribeConfigurationRecorders(*configservice.DescribeConfigurationRecordersInput) (*configservice.DescribeConfigurationRecordersOutput, error)

	DescribeDeliveryChannelStatusRequest(*configservice.DescribeDeliveryChannelStatusInput) (*service.Request, *configservice.DescribeDeliveryChannelStatusOutput)

	DescribeDeliveryChannelStatus(*configservice.DescribeDeliveryChannelStatusInput) (*configservice.DescribeDeliveryChannelStatusOutput, error)

	DescribeDeliveryChannelsRequest(*configservice.DescribeDeliveryChannelsInput) (*service.Request, *configservice.DescribeDeliveryChannelsOutput)

	DescribeDeliveryChannels(*configservice.DescribeDeliveryChannelsInput) (*configservice.DescribeDeliveryChannelsOutput, error)

	GetResourceConfigHistoryRequest(*configservice.GetResourceConfigHistoryInput) (*service.Request, *configservice.GetResourceConfigHistoryOutput)

	GetResourceConfigHistory(*configservice.GetResourceConfigHistoryInput) (*configservice.GetResourceConfigHistoryOutput, error)

	GetResourceConfigHistoryPages(*configservice.GetResourceConfigHistoryInput, func(*configservice.GetResourceConfigHistoryOutput, bool) bool) error

	PutConfigurationRecorderRequest(*configservice.PutConfigurationRecorderInput) (*service.Request, *configservice.PutConfigurationRecorderOutput)

	PutConfigurationRecorder(*configservice.PutConfigurationRecorderInput) (*configservice.PutConfigurationRecorderOutput, error)

	PutDeliveryChannelRequest(*configservice.PutDeliveryChannelInput) (*service.Request, *configservice.PutDeliveryChannelOutput)

	PutDeliveryChannel(*configservice.PutDeliveryChannelInput) (*configservice.PutDeliveryChannelOutput, error)

	StartConfigurationRecorderRequest(*configservice.StartConfigurationRecorderInput) (*service.Request, *configservice.StartConfigurationRecorderOutput)

	StartConfigurationRecorder(*configservice.StartConfigurationRecorderInput) (*configservice.StartConfigurationRecorderOutput, error)

	StopConfigurationRecorderRequest(*configservice.StopConfigurationRecorderInput) (*service.Request, *configservice.StopConfigurationRecorderOutput)

	StopConfigurationRecorder(*configservice.StopConfigurationRecorderInput) (*configservice.StopConfigurationRecorderOutput, error)
}