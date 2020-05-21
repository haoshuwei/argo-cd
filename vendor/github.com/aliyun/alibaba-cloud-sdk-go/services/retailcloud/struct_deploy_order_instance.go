package retailcloud

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

// DeployOrderInstance is a nested struct in retailcloud response
type DeployOrderInstance struct {
	AppInstanceType     string `json:"AppInstanceType" xml:"AppInstanceType"`
	CurrentPartitionNum int    `json:"CurrentPartitionNum" xml:"CurrentPartitionNum"`
	DeployOrderId       int64  `json:"DeployOrderId" xml:"DeployOrderId"`
	DeployPauseType     string `json:"DeployPauseType" xml:"DeployPauseType"`
	DeployPauseTypeName string `json:"DeployPauseTypeName" xml:"DeployPauseTypeName"`
	DeployType          string `json:"DeployType" xml:"DeployType"`
	DeployTypeName      string `json:"DeployTypeName" xml:"DeployTypeName"`
	Description         string `json:"Description" xml:"Description"`
	ElapsedTime         int    `json:"ElapsedTime" xml:"ElapsedTime"`
	EndTime             string `json:"EndTime" xml:"EndTime"`
	EnvId               int64  `json:"EnvId" xml:"EnvId"`
	EnvType             string `json:"EnvType" xml:"EnvType"`
	FailureRate         string `json:"FailureRate" xml:"FailureRate"`
	FinishAppInstanceCt int    `json:"FinishAppInstanceCt" xml:"FinishAppInstanceCt"`
	Name                string `json:"Name" xml:"Name"`
	PartitionType       string `json:"PartitionType" xml:"PartitionType"`
	PartitionTypeName   string `json:"PartitionTypeName" xml:"PartitionTypeName"`
	Result              int    `json:"Result" xml:"Result"`
	ResultName          string `json:"ResultName" xml:"ResultName"`
	SchemaId            int64  `json:"SchemaId" xml:"SchemaId"`
	StartTime           string `json:"StartTime" xml:"StartTime"`
	Status              int    `json:"Status" xml:"Status"`
	StatusName          string `json:"StatusName" xml:"StatusName"`
	TotalAppInstanceCt  int    `json:"TotalAppInstanceCt" xml:"TotalAppInstanceCt"`
	TotalPartitions     int    `json:"TotalPartitions" xml:"TotalPartitions"`
	UserId              string `json:"UserId" xml:"UserId"`
	UserNick            string `json:"UserNick" xml:"UserNick"`
}
