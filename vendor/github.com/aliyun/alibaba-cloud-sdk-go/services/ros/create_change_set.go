package ros

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

// CreateChangeSet invokes the ros.CreateChangeSet API synchronously
// api document: https://help.aliyun.com/api/ros/createchangeset.html
func (client *Client) CreateChangeSet(request *CreateChangeSetRequest) (response *CreateChangeSetResponse, err error) {
	response = CreateCreateChangeSetResponse()
	err = client.DoAction(request, response)
	return
}

// CreateChangeSetWithChan invokes the ros.CreateChangeSet API asynchronously
// api document: https://help.aliyun.com/api/ros/createchangeset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChangeSetWithChan(request *CreateChangeSetRequest) (<-chan *CreateChangeSetResponse, <-chan error) {
	responseChan := make(chan *CreateChangeSetResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.CreateChangeSet(request)
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

// CreateChangeSetWithCallback invokes the ros.CreateChangeSet API asynchronously
// api document: https://help.aliyun.com/api/ros/createchangeset.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) CreateChangeSetWithCallback(request *CreateChangeSetRequest, callback func(response *CreateChangeSetResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *CreateChangeSetResponse
		var err error
		defer close(result)
		response, err = client.CreateChangeSet(request)
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

// CreateChangeSetRequest is the request struct for api CreateChangeSet
type CreateChangeSetRequest struct {
	*requests.RoaRequest
}

// CreateChangeSetResponse is the response struct for api CreateChangeSet
type CreateChangeSetResponse struct {
	*responses.BaseResponse
	Dummy string `json:"Dummy" xml:"Dummy"`
}

// CreateCreateChangeSetRequest creates a request to invoke CreateChangeSet API
func CreateCreateChangeSetRequest() (request *CreateChangeSetRequest) {
	request = &CreateChangeSetRequest{
		RoaRequest: &requests.RoaRequest{},
	}
	request.InitWithApiInfo("ROS", "2015-09-01", "CreateChangeSet", "/changeSets", "ROS", "openAPI")
	request.Method = requests.POST
	return
}

// CreateCreateChangeSetResponse creates a response to parse from CreateChangeSet response
func CreateCreateChangeSetResponse() (response *CreateChangeSetResponse) {
	response = &CreateChangeSetResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
