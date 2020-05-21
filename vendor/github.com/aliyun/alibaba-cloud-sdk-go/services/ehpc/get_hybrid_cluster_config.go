package ehpc

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

// GetHybridClusterConfig invokes the ehpc.GetHybridClusterConfig API synchronously
// api document: https://help.aliyun.com/api/ehpc/gethybridclusterconfig.html
func (client *Client) GetHybridClusterConfig(request *GetHybridClusterConfigRequest) (response *GetHybridClusterConfigResponse, err error) {
	response = CreateGetHybridClusterConfigResponse()
	err = client.DoAction(request, response)
	return
}

// GetHybridClusterConfigWithChan invokes the ehpc.GetHybridClusterConfig API asynchronously
// api document: https://help.aliyun.com/api/ehpc/gethybridclusterconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetHybridClusterConfigWithChan(request *GetHybridClusterConfigRequest) (<-chan *GetHybridClusterConfigResponse, <-chan error) {
	responseChan := make(chan *GetHybridClusterConfigResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.GetHybridClusterConfig(request)
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

// GetHybridClusterConfigWithCallback invokes the ehpc.GetHybridClusterConfig API asynchronously
// api document: https://help.aliyun.com/api/ehpc/gethybridclusterconfig.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) GetHybridClusterConfigWithCallback(request *GetHybridClusterConfigRequest, callback func(response *GetHybridClusterConfigResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *GetHybridClusterConfigResponse
		var err error
		defer close(result)
		response, err = client.GetHybridClusterConfig(request)
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

// GetHybridClusterConfigRequest is the request struct for api GetHybridClusterConfig
type GetHybridClusterConfigRequest struct {
	*requests.RpcRequest
	ClusterId string `position:"Query" name:"ClusterId"`
	Node      string `position:"Query" name:"Node"`
}

// GetHybridClusterConfigResponse is the response struct for api GetHybridClusterConfig
type GetHybridClusterConfigResponse struct {
	*responses.BaseResponse
	RequestId     string `json:"RequestId" xml:"RequestId"`
	ClusterConfig string `json:"ClusterConfig" xml:"ClusterConfig"`
}

// CreateGetHybridClusterConfigRequest creates a request to invoke GetHybridClusterConfig API
func CreateGetHybridClusterConfigRequest() (request *GetHybridClusterConfigRequest) {
	request = &GetHybridClusterConfigRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "GetHybridClusterConfig", "", "")
	return
}

// CreateGetHybridClusterConfigResponse creates a response to parse from GetHybridClusterConfig response
func CreateGetHybridClusterConfigResponse() (response *GetHybridClusterConfigResponse) {
	response = &GetHybridClusterConfigResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
