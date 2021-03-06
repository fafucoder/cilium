// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package ec2

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/internal/awsutil"
)

type MonitorInstancesInput struct {
	_ struct{} `type:"structure"`

	// Checks whether you have the required permissions for the action, without
	// actually making the request, and provides an error response. If you have
	// the required permissions, the error response is DryRunOperation. Otherwise,
	// it is UnauthorizedOperation.
	DryRun *bool `locationName:"dryRun" type:"boolean"`

	// The IDs of the instances.
	//
	// InstanceIds is a required field
	InstanceIds []string `locationName:"InstanceId" locationNameList:"InstanceId" type:"list" required:"true"`
}

// String returns the string representation
func (s MonitorInstancesInput) String() string {
	return awsutil.Prettify(s)
}

// Validate inspects the fields of the type to determine if they are valid.
func (s *MonitorInstancesInput) Validate() error {
	invalidParams := aws.ErrInvalidParams{Context: "MonitorInstancesInput"}

	if s.InstanceIds == nil {
		invalidParams.Add(aws.NewErrParamRequired("InstanceIds"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type MonitorInstancesOutput struct {
	_ struct{} `type:"structure"`

	// The monitoring information.
	InstanceMonitorings []InstanceMonitoring `locationName:"instancesSet" locationNameList:"item" type:"list"`
}

// String returns the string representation
func (s MonitorInstancesOutput) String() string {
	return awsutil.Prettify(s)
}

const opMonitorInstances = "MonitorInstances"

// MonitorInstancesRequest returns a request value for making API operation for
// Amazon Elastic Compute Cloud.
//
// Enables detailed monitoring for a running instance. Otherwise, basic monitoring
// is enabled. For more information, see Monitoring your instances and volumes
// (https://docs.aws.amazon.com/AWSEC2/latest/UserGuide/using-cloudwatch.html)
// in the Amazon Elastic Compute Cloud User Guide.
//
// To disable detailed monitoring, see .
//
//    // Example sending a request using MonitorInstancesRequest.
//    req := client.MonitorInstancesRequest(params)
//    resp, err := req.Send(context.TODO())
//    if err == nil {
//        fmt.Println(resp)
//    }
//
// Please also see https://docs.aws.amazon.com/goto/WebAPI/ec2-2016-11-15/MonitorInstances
func (c *Client) MonitorInstancesRequest(input *MonitorInstancesInput) MonitorInstancesRequest {
	op := &aws.Operation{
		Name:       opMonitorInstances,
		HTTPMethod: "POST",
		HTTPPath:   "/",
	}

	if input == nil {
		input = &MonitorInstancesInput{}
	}

	req := c.newRequest(op, input, &MonitorInstancesOutput{})

	return MonitorInstancesRequest{Request: req, Input: input, Copy: c.MonitorInstancesRequest}
}

// MonitorInstancesRequest is the request type for the
// MonitorInstances API operation.
type MonitorInstancesRequest struct {
	*aws.Request
	Input *MonitorInstancesInput
	Copy  func(*MonitorInstancesInput) MonitorInstancesRequest
}

// Send marshals and sends the MonitorInstances API request.
func (r MonitorInstancesRequest) Send(ctx context.Context) (*MonitorInstancesResponse, error) {
	r.Request.SetContext(ctx)
	err := r.Request.Send()
	if err != nil {
		return nil, err
	}

	resp := &MonitorInstancesResponse{
		MonitorInstancesOutput: r.Request.Data.(*MonitorInstancesOutput),
		response:               &aws.Response{Request: r.Request},
	}

	return resp, nil
}

// MonitorInstancesResponse is the response type for the
// MonitorInstances API operation.
type MonitorInstancesResponse struct {
	*MonitorInstancesOutput

	response *aws.Response
}

// SDKResponseMetdata returns the response metadata for the
// MonitorInstances request.
func (r *MonitorInstancesResponse) SDKResponseMetdata() *aws.Response {
	return r.response
}
