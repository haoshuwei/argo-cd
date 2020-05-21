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

// DescribeGWSClusters invokes the ehpc.DescribeGWSClusters API synchronously
// api document: https://help.aliyun.com/api/ehpc/describegwsclusters.html
func (client *Client) DescribeGWSClusters(request *DescribeGWSClustersRequest) (response *DescribeGWSClustersResponse, err error) {
	response = CreateDescribeGWSClustersResponse()
	err = client.DoAction(request, response)
	return
}

// DescribeGWSClustersWithChan invokes the ehpc.DescribeGWSClusters API asynchronously
// api document: https://help.aliyun.com/api/ehpc/describegwsclusters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGWSClustersWithChan(request *DescribeGWSClustersRequest) (<-chan *DescribeGWSClustersResponse, <-chan error) {
	responseChan := make(chan *DescribeGWSClustersResponse, 1)
	errChan := make(chan error, 1)
	err := client.AddAsyncTask(func() {
		defer close(responseChan)
		defer close(errChan)
		response, err := client.DescribeGWSClusters(request)
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

// DescribeGWSClustersWithCallback invokes the ehpc.DescribeGWSClusters API asynchronously
// api document: https://help.aliyun.com/api/ehpc/describegwsclusters.html
// asynchronous document: https://help.aliyun.com/document_detail/66220.html
func (client *Client) DescribeGWSClustersWithCallback(request *DescribeGWSClustersRequest, callback func(response *DescribeGWSClustersResponse, err error)) <-chan int {
	result := make(chan int, 1)
	err := client.AddAsyncTask(func() {
		var response *DescribeGWSClustersResponse
		var err error
		defer close(result)
		response, err = client.DescribeGWSClusters(request)
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

// DescribeGWSClustersRequest is the request struct for api DescribeGWSClusters
type DescribeGWSClustersRequest struct {
	*requests.RpcRequest
	ClusterId  string           `position:"Query" name:"ClusterId"`
	PageNumber requests.Integer `position:"Query" name:"PageNumber"`
	PageSize   requests.Integer `position:"Query" name:"PageSize"`
}

// DescribeGWSClustersResponse is the response struct for api DescribeGWSClusters
type DescribeGWSClustersResponse struct {
	*responses.BaseResponse
	RequestId  string                        `json:"RequestId" xml:"RequestId"`
	TotalCount int                           `json:"TotalCount" xml:"TotalCount"`
	PageNumber int                           `json:"PageNumber" xml:"PageNumber"`
	PageSize   int                           `json:"PageSize" xml:"PageSize"`
	CallerType string                        `json:"CallerType" xml:"CallerType"`
	Clusters   ClustersInDescribeGWSClusters `json:"Clusters" xml:"Clusters"`
}

// CreateDescribeGWSClustersRequest creates a request to invoke DescribeGWSClusters API
func CreateDescribeGWSClustersRequest() (request *DescribeGWSClustersRequest) {
	request = &DescribeGWSClustersRequest{
		RpcRequest: &requests.RpcRequest{},
	}
	request.InitWithApiInfo("EHPC", "2018-04-12", "DescribeGWSClusters", "", "")
	return
}

// CreateDescribeGWSClustersResponse creates a response to parse from DescribeGWSClusters response
func CreateDescribeGWSClustersResponse() (response *DescribeGWSClustersResponse) {
	response = &DescribeGWSClustersResponse{
		BaseResponse: &responses.BaseResponse{},
	}
	return
}
