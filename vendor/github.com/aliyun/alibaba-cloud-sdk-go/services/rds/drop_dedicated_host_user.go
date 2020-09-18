package rds

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

// DropDedicatedHostUser invokes the rds.DropDedicatedHostUser API synchronously
// api document: https://help.aliyun.com/api/rds/dropdedicatedhostuser.html
func (client *Client) DropDedicatedHostUser(request *DropDedicatedHostUserRequest) (response *DropDedicatedHostUserResponse, err error) {
	response = CreateDropDedicatedHostUserResponse()
	err = client.DoAction(request, response)
	return
}

// DropDedicatedHostUserWithChan invokes the rds.DropDedicatedHostUser API asynchronously
// api document: https://help.aliyun.com/api/rds/dropdedicatedhostuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DropDedicatedHostUserWithChan(request *DropDedicatedHostUserRequest) (<-chan *DropDedicatedHostUserResponse, <-chan error) {
	responseChan := make(chan *DropDedicatedHostUserResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DropDedicatedHostUser(request)
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

// DropDedicatedHostUserWithCallback invokes the rds.DropDedicatedHostUser API asynchronously
// api document: https://help.aliyun.com/api/rds/dropdedicatedhostuser.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DropDedicatedHostUserWithCallback(request *DropDedicatedHostUserRequest, callback func(response *DropDedicatedHostUserResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DropDedicatedHostUserResponse
		var err error
		defer close(result)
		response, err = client.DropDedicatedHostUser(request)
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

// DropDedicatedHostUserRequest is the request struct for api DropDedicatedHostUser
type DropDedicatedHostUserRequest struct {
	*requests.RpcRequest
	ResourceOwnerId      requests.Integer `position:"Query" name:"ResourceOwnerId"`
	DedicatedHostName    string           `position:"Query" name:"DedicatedHostName"`
	ResourceOwnerAccount string           `position:"Query" name:"ResourceOwnerAccount"`
	OwnerId              requests.Integer `position:"Query" name:"OwnerId"`
	UserName             string           `position:"Query" name:"UserName"`
}

// DropDedicatedHostUserResponse is the response struct for api DropDedicatedHostUser
type DropDedicatedHostUserResponse struct {
	*responses.BaseResponse
	RequestId string `json:"RequestId" xml:"RequestId"`
}

// CreateDropDedicatedHostUserRequest creates a request to invoke DropDedicatedHostUser API
func CreateDropDedicatedHostUserRequest() (request *DropDedicatedHostUserRequest) {
	request = &DropDedicatedHostUserRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("Rds", "2014-08-15", "DropDedicatedHostUser", "rds", "openAPI")
	request.Method = requests.POST
	return
}

// CreateDropDedicatedHostUserResponse creates a response to parse from DropDedicatedHostUser response
func CreateDropDedicatedHostUserResponse() (response *DropDedicatedHostUserResponse) {
	response = &DropDedicatedHostUserResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}