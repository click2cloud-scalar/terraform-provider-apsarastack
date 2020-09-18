package vpc

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

import (
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/requests"
	"github.com/aliyun/alibaba-cloud-sdk-go/sdk/responses"
)

// DescribeVpcs invokes the vpc.DescribeVpcs API synchronously
// api document: https://help.aliyun.com/api/vpc/describevpcs.html
func (client *Client) DescribeVpcs(request *DescribeVpcsRequest) (response *DescribeVpcsResponse, err error) {
	response = CreateDescribeVpcsResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeVpcsWithChan invokes the vpc.DescribeVpcs API asynchronously
// api document: https://help.aliyun.com/api/vpc/describevpcs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVpcsWithChan(request *DescribeVpcsRequest) (<-chan *DescribeVpcsResponse, <-chan error) {
	responseChan := make(chan *DescribeVpcsResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeVpcs(request)
		if err != nil {
			errChan <- err
		} else {
			responseChan <- response
		}
	})
	if err != nil {
		errChan <- err
		close(responseChan)
		close(errChan)
	}
	return responseChan, errChan
}

// DescribeVpcsWithCallback invokes the vpc.DescribeVpcs API asynchronously
// api document: https://help.aliyun.com/api/vpc/describevpcs.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeVpcsWithCallback(request *DescribeVpcsRequest, callback func(response *DescribeVpcsResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeVpcsResponse
		var err error
		defer close(result)
		response, err = client.DescribeVpcs(request)
		callback(response, err)
		result <- 1
	})
	if err != nil {
		defer close(result)
		callback(nil, err)
		result <- 0
	}
	return result
}

// DescribeVpcsRequest is the request struct for api DescribeVpcs
type DescribeVpcsRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer   `position:"Query" name:"ResourceOwnerId"`
	VpcOwnerId           requests.Integer   `position:"Query" name:"VpcOwnerId"`
	PageNumber           requests.Integer   `position:"Query" name:"PageNumber"`
	VpcName              string             `position:"Query" name:"VpcName"`
	ResourceGroupId      string             `position:"Query" name:"ResourceGroupId"`
	PageSize             requests.Integer   `position:"Query" name:"PageSize"`
	Tag                  *[]DescribeVpcsTag `position:"Query" name:"Tag"  type:"Repeated"`
	IsDefault            requests.Boolean   `position:"Query" name:"IsDefault"`
	DryRun               requests.Boolean   `position:"Query" name:"DryRun"`
	DhcpOptionsSetId     string             `position:"Query" name:"DhcpOptionsSetId"`
	ResourceOwnerAccount string             `position:"Query" name:"ResourceOwnerAccount"`
	OwnerAccount         string             `position:"Query" name:"OwnerAccount"`
	AdvancedFeature      requests.Boolean   `position:"Query" name:"AdvancedFeature"`
	OwnerId              requests.Integer   `position:"Query" name:"OwnerId"`
	VpcId                string             `position:"Query" name:"VpcId"`
}

// DescribeVpcsTag is a repeated param struct in DescribeVpcsRequest
type DescribeVpcsTag struct {
	Value string `name:"Value"`
	Key   string `name:"Key"`
}

// DescribeVpcsResponse is the response struct for api DescribeVpcs
type DescribeVpcsResponse struct {
	*responses.BaseResponse
	RequestId  string `json:"RequestId" xml:"RequestId"`
	TotalCount int    `json:"TotalCount" xml:"TotalCount"`
	PageNumber int    `json:"PageNumber" xml:"PageNumber"`
	PageSize   int    `json:"PageSize" xml:"PageSize"`
	Vpcs       Vpcs   `json:"Vpcs" xml:"Vpcs"`
}

// CreateDescribeVpcsRequest creates a request to invoke DescribeVpcs API
func CreateDescribeVpcsRequest() (request *DescribeVpcsRequest) {
	request = &DescribeVpcsRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Vpc", "2016-04-28", "DescribeVpcs", "vpc", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDescribeVpcsResponse creates a response to parse from DescribeVpcs response
func CreateDescribeVpcsResponse() (response *DescribeVpcsResponse) {
	response = &DescribeVpcsResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}