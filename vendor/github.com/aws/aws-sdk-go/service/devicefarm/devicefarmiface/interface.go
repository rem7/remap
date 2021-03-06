// THIS FILE IS AUTOMATICALLY GENERATED. DO NOT EDIT.

// Package devicefarmiface provides an interface to enable mocking the AWS Device Farm service client
// for testing your code.
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters.
package devicefarmiface

import (
	"github.com/aws/aws-sdk-go/aws/request"
	"github.com/aws/aws-sdk-go/service/devicefarm"
)

// DeviceFarmAPI provides an interface to enable mocking the
// devicefarm.DeviceFarm service client's API operation,
// paginators, and waiters. This make unit testing your code that calls out
// to the SDK's service client's calls easier.
//
// The best way to use this interface is so the SDK's service client's calls
// can be stubbed out for unit testing your code with the SDK without needing
// to inject custom request handlers into the the SDK's request pipeline.
//
//    // myFunc uses an SDK service client to make a request to
//    // AWS Device Farm.
//    func myFunc(svc devicefarmiface.DeviceFarmAPI) bool {
//        // Make svc.CreateDevicePool request
//    }
//
//    func main() {
//        sess := session.New()
//        svc := devicefarm.New(sess)
//
//        myFunc(svc)
//    }
//
// In your _test.go file:
//
//    // Define a mock struct to be used in your unit tests of myFunc.
//    type mockDeviceFarmClient struct {
//        devicefarmiface.DeviceFarmAPI
//    }
//    func (m *mockDeviceFarmClient) CreateDevicePool(input *devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error) {
//        // mock response/functionality
//    }
//
//    func TestMyFunc(t *testing.T) {
//        // Setup Test
//        mockSvc := &mockDeviceFarmClient{}
//
//        myfunc(mockSvc)
//
//        // Verify myFunc's functionality
//    }
//
// It is important to note that this interface will have breaking changes
// when the service model is updated and adds new API operations, paginators,
// and waiters. Its suggested to use the pattern above for testing, or using
// tooling to generate mocks to satisfy the interfaces.
type DeviceFarmAPI interface {
	CreateDevicePoolRequest(*devicefarm.CreateDevicePoolInput) (*request.Request, *devicefarm.CreateDevicePoolOutput)

	CreateDevicePool(*devicefarm.CreateDevicePoolInput) (*devicefarm.CreateDevicePoolOutput, error)

	CreateNetworkProfileRequest(*devicefarm.CreateNetworkProfileInput) (*request.Request, *devicefarm.CreateNetworkProfileOutput)

	CreateNetworkProfile(*devicefarm.CreateNetworkProfileInput) (*devicefarm.CreateNetworkProfileOutput, error)

	CreateProjectRequest(*devicefarm.CreateProjectInput) (*request.Request, *devicefarm.CreateProjectOutput)

	CreateProject(*devicefarm.CreateProjectInput) (*devicefarm.CreateProjectOutput, error)

	CreateRemoteAccessSessionRequest(*devicefarm.CreateRemoteAccessSessionInput) (*request.Request, *devicefarm.CreateRemoteAccessSessionOutput)

	CreateRemoteAccessSession(*devicefarm.CreateRemoteAccessSessionInput) (*devicefarm.CreateRemoteAccessSessionOutput, error)

	CreateUploadRequest(*devicefarm.CreateUploadInput) (*request.Request, *devicefarm.CreateUploadOutput)

	CreateUpload(*devicefarm.CreateUploadInput) (*devicefarm.CreateUploadOutput, error)

	DeleteDevicePoolRequest(*devicefarm.DeleteDevicePoolInput) (*request.Request, *devicefarm.DeleteDevicePoolOutput)

	DeleteDevicePool(*devicefarm.DeleteDevicePoolInput) (*devicefarm.DeleteDevicePoolOutput, error)

	DeleteNetworkProfileRequest(*devicefarm.DeleteNetworkProfileInput) (*request.Request, *devicefarm.DeleteNetworkProfileOutput)

	DeleteNetworkProfile(*devicefarm.DeleteNetworkProfileInput) (*devicefarm.DeleteNetworkProfileOutput, error)

	DeleteProjectRequest(*devicefarm.DeleteProjectInput) (*request.Request, *devicefarm.DeleteProjectOutput)

	DeleteProject(*devicefarm.DeleteProjectInput) (*devicefarm.DeleteProjectOutput, error)

	DeleteRemoteAccessSessionRequest(*devicefarm.DeleteRemoteAccessSessionInput) (*request.Request, *devicefarm.DeleteRemoteAccessSessionOutput)

	DeleteRemoteAccessSession(*devicefarm.DeleteRemoteAccessSessionInput) (*devicefarm.DeleteRemoteAccessSessionOutput, error)

	DeleteRunRequest(*devicefarm.DeleteRunInput) (*request.Request, *devicefarm.DeleteRunOutput)

	DeleteRun(*devicefarm.DeleteRunInput) (*devicefarm.DeleteRunOutput, error)

	DeleteUploadRequest(*devicefarm.DeleteUploadInput) (*request.Request, *devicefarm.DeleteUploadOutput)

	DeleteUpload(*devicefarm.DeleteUploadInput) (*devicefarm.DeleteUploadOutput, error)

	GetAccountSettingsRequest(*devicefarm.GetAccountSettingsInput) (*request.Request, *devicefarm.GetAccountSettingsOutput)

	GetAccountSettings(*devicefarm.GetAccountSettingsInput) (*devicefarm.GetAccountSettingsOutput, error)

	GetDeviceRequest(*devicefarm.GetDeviceInput) (*request.Request, *devicefarm.GetDeviceOutput)

	GetDevice(*devicefarm.GetDeviceInput) (*devicefarm.GetDeviceOutput, error)

	GetDevicePoolRequest(*devicefarm.GetDevicePoolInput) (*request.Request, *devicefarm.GetDevicePoolOutput)

	GetDevicePool(*devicefarm.GetDevicePoolInput) (*devicefarm.GetDevicePoolOutput, error)

	GetDevicePoolCompatibilityRequest(*devicefarm.GetDevicePoolCompatibilityInput) (*request.Request, *devicefarm.GetDevicePoolCompatibilityOutput)

	GetDevicePoolCompatibility(*devicefarm.GetDevicePoolCompatibilityInput) (*devicefarm.GetDevicePoolCompatibilityOutput, error)

	GetJobRequest(*devicefarm.GetJobInput) (*request.Request, *devicefarm.GetJobOutput)

	GetJob(*devicefarm.GetJobInput) (*devicefarm.GetJobOutput, error)

	GetNetworkProfileRequest(*devicefarm.GetNetworkProfileInput) (*request.Request, *devicefarm.GetNetworkProfileOutput)

	GetNetworkProfile(*devicefarm.GetNetworkProfileInput) (*devicefarm.GetNetworkProfileOutput, error)

	GetOfferingStatusRequest(*devicefarm.GetOfferingStatusInput) (*request.Request, *devicefarm.GetOfferingStatusOutput)

	GetOfferingStatus(*devicefarm.GetOfferingStatusInput) (*devicefarm.GetOfferingStatusOutput, error)

	GetOfferingStatusPages(*devicefarm.GetOfferingStatusInput, func(*devicefarm.GetOfferingStatusOutput, bool) bool) error

	GetProjectRequest(*devicefarm.GetProjectInput) (*request.Request, *devicefarm.GetProjectOutput)

	GetProject(*devicefarm.GetProjectInput) (*devicefarm.GetProjectOutput, error)

	GetRemoteAccessSessionRequest(*devicefarm.GetRemoteAccessSessionInput) (*request.Request, *devicefarm.GetRemoteAccessSessionOutput)

	GetRemoteAccessSession(*devicefarm.GetRemoteAccessSessionInput) (*devicefarm.GetRemoteAccessSessionOutput, error)

	GetRunRequest(*devicefarm.GetRunInput) (*request.Request, *devicefarm.GetRunOutput)

	GetRun(*devicefarm.GetRunInput) (*devicefarm.GetRunOutput, error)

	GetSuiteRequest(*devicefarm.GetSuiteInput) (*request.Request, *devicefarm.GetSuiteOutput)

	GetSuite(*devicefarm.GetSuiteInput) (*devicefarm.GetSuiteOutput, error)

	GetTestRequest(*devicefarm.GetTestInput) (*request.Request, *devicefarm.GetTestOutput)

	GetTest(*devicefarm.GetTestInput) (*devicefarm.GetTestOutput, error)

	GetUploadRequest(*devicefarm.GetUploadInput) (*request.Request, *devicefarm.GetUploadOutput)

	GetUpload(*devicefarm.GetUploadInput) (*devicefarm.GetUploadOutput, error)

	InstallToRemoteAccessSessionRequest(*devicefarm.InstallToRemoteAccessSessionInput) (*request.Request, *devicefarm.InstallToRemoteAccessSessionOutput)

	InstallToRemoteAccessSession(*devicefarm.InstallToRemoteAccessSessionInput) (*devicefarm.InstallToRemoteAccessSessionOutput, error)

	ListArtifactsRequest(*devicefarm.ListArtifactsInput) (*request.Request, *devicefarm.ListArtifactsOutput)

	ListArtifacts(*devicefarm.ListArtifactsInput) (*devicefarm.ListArtifactsOutput, error)

	ListArtifactsPages(*devicefarm.ListArtifactsInput, func(*devicefarm.ListArtifactsOutput, bool) bool) error

	ListDevicePoolsRequest(*devicefarm.ListDevicePoolsInput) (*request.Request, *devicefarm.ListDevicePoolsOutput)

	ListDevicePools(*devicefarm.ListDevicePoolsInput) (*devicefarm.ListDevicePoolsOutput, error)

	ListDevicePoolsPages(*devicefarm.ListDevicePoolsInput, func(*devicefarm.ListDevicePoolsOutput, bool) bool) error

	ListDevicesRequest(*devicefarm.ListDevicesInput) (*request.Request, *devicefarm.ListDevicesOutput)

	ListDevices(*devicefarm.ListDevicesInput) (*devicefarm.ListDevicesOutput, error)

	ListDevicesPages(*devicefarm.ListDevicesInput, func(*devicefarm.ListDevicesOutput, bool) bool) error

	ListJobsRequest(*devicefarm.ListJobsInput) (*request.Request, *devicefarm.ListJobsOutput)

	ListJobs(*devicefarm.ListJobsInput) (*devicefarm.ListJobsOutput, error)

	ListJobsPages(*devicefarm.ListJobsInput, func(*devicefarm.ListJobsOutput, bool) bool) error

	ListNetworkProfilesRequest(*devicefarm.ListNetworkProfilesInput) (*request.Request, *devicefarm.ListNetworkProfilesOutput)

	ListNetworkProfiles(*devicefarm.ListNetworkProfilesInput) (*devicefarm.ListNetworkProfilesOutput, error)

	ListOfferingTransactionsRequest(*devicefarm.ListOfferingTransactionsInput) (*request.Request, *devicefarm.ListOfferingTransactionsOutput)

	ListOfferingTransactions(*devicefarm.ListOfferingTransactionsInput) (*devicefarm.ListOfferingTransactionsOutput, error)

	ListOfferingTransactionsPages(*devicefarm.ListOfferingTransactionsInput, func(*devicefarm.ListOfferingTransactionsOutput, bool) bool) error

	ListOfferingsRequest(*devicefarm.ListOfferingsInput) (*request.Request, *devicefarm.ListOfferingsOutput)

	ListOfferings(*devicefarm.ListOfferingsInput) (*devicefarm.ListOfferingsOutput, error)

	ListOfferingsPages(*devicefarm.ListOfferingsInput, func(*devicefarm.ListOfferingsOutput, bool) bool) error

	ListProjectsRequest(*devicefarm.ListProjectsInput) (*request.Request, *devicefarm.ListProjectsOutput)

	ListProjects(*devicefarm.ListProjectsInput) (*devicefarm.ListProjectsOutput, error)

	ListProjectsPages(*devicefarm.ListProjectsInput, func(*devicefarm.ListProjectsOutput, bool) bool) error

	ListRemoteAccessSessionsRequest(*devicefarm.ListRemoteAccessSessionsInput) (*request.Request, *devicefarm.ListRemoteAccessSessionsOutput)

	ListRemoteAccessSessions(*devicefarm.ListRemoteAccessSessionsInput) (*devicefarm.ListRemoteAccessSessionsOutput, error)

	ListRunsRequest(*devicefarm.ListRunsInput) (*request.Request, *devicefarm.ListRunsOutput)

	ListRuns(*devicefarm.ListRunsInput) (*devicefarm.ListRunsOutput, error)

	ListRunsPages(*devicefarm.ListRunsInput, func(*devicefarm.ListRunsOutput, bool) bool) error

	ListSamplesRequest(*devicefarm.ListSamplesInput) (*request.Request, *devicefarm.ListSamplesOutput)

	ListSamples(*devicefarm.ListSamplesInput) (*devicefarm.ListSamplesOutput, error)

	ListSamplesPages(*devicefarm.ListSamplesInput, func(*devicefarm.ListSamplesOutput, bool) bool) error

	ListSuitesRequest(*devicefarm.ListSuitesInput) (*request.Request, *devicefarm.ListSuitesOutput)

	ListSuites(*devicefarm.ListSuitesInput) (*devicefarm.ListSuitesOutput, error)

	ListSuitesPages(*devicefarm.ListSuitesInput, func(*devicefarm.ListSuitesOutput, bool) bool) error

	ListTestsRequest(*devicefarm.ListTestsInput) (*request.Request, *devicefarm.ListTestsOutput)

	ListTests(*devicefarm.ListTestsInput) (*devicefarm.ListTestsOutput, error)

	ListTestsPages(*devicefarm.ListTestsInput, func(*devicefarm.ListTestsOutput, bool) bool) error

	ListUniqueProblemsRequest(*devicefarm.ListUniqueProblemsInput) (*request.Request, *devicefarm.ListUniqueProblemsOutput)

	ListUniqueProblems(*devicefarm.ListUniqueProblemsInput) (*devicefarm.ListUniqueProblemsOutput, error)

	ListUniqueProblemsPages(*devicefarm.ListUniqueProblemsInput, func(*devicefarm.ListUniqueProblemsOutput, bool) bool) error

	ListUploadsRequest(*devicefarm.ListUploadsInput) (*request.Request, *devicefarm.ListUploadsOutput)

	ListUploads(*devicefarm.ListUploadsInput) (*devicefarm.ListUploadsOutput, error)

	ListUploadsPages(*devicefarm.ListUploadsInput, func(*devicefarm.ListUploadsOutput, bool) bool) error

	PurchaseOfferingRequest(*devicefarm.PurchaseOfferingInput) (*request.Request, *devicefarm.PurchaseOfferingOutput)

	PurchaseOffering(*devicefarm.PurchaseOfferingInput) (*devicefarm.PurchaseOfferingOutput, error)

	RenewOfferingRequest(*devicefarm.RenewOfferingInput) (*request.Request, *devicefarm.RenewOfferingOutput)

	RenewOffering(*devicefarm.RenewOfferingInput) (*devicefarm.RenewOfferingOutput, error)

	ScheduleRunRequest(*devicefarm.ScheduleRunInput) (*request.Request, *devicefarm.ScheduleRunOutput)

	ScheduleRun(*devicefarm.ScheduleRunInput) (*devicefarm.ScheduleRunOutput, error)

	StopRemoteAccessSessionRequest(*devicefarm.StopRemoteAccessSessionInput) (*request.Request, *devicefarm.StopRemoteAccessSessionOutput)

	StopRemoteAccessSession(*devicefarm.StopRemoteAccessSessionInput) (*devicefarm.StopRemoteAccessSessionOutput, error)

	StopRunRequest(*devicefarm.StopRunInput) (*request.Request, *devicefarm.StopRunOutput)

	StopRun(*devicefarm.StopRunInput) (*devicefarm.StopRunOutput, error)

	UpdateDevicePoolRequest(*devicefarm.UpdateDevicePoolInput) (*request.Request, *devicefarm.UpdateDevicePoolOutput)

	UpdateDevicePool(*devicefarm.UpdateDevicePoolInput) (*devicefarm.UpdateDevicePoolOutput, error)

	UpdateNetworkProfileRequest(*devicefarm.UpdateNetworkProfileInput) (*request.Request, *devicefarm.UpdateNetworkProfileOutput)

	UpdateNetworkProfile(*devicefarm.UpdateNetworkProfileInput) (*devicefarm.UpdateNetworkProfileOutput, error)

	UpdateProjectRequest(*devicefarm.UpdateProjectInput) (*request.Request, *devicefarm.UpdateProjectOutput)

	UpdateProject(*devicefarm.UpdateProjectInput) (*devicefarm.UpdateProjectOutput, error)
}

var _ DeviceFarmAPI = (*devicefarm.DeviceFarm)(nil)
