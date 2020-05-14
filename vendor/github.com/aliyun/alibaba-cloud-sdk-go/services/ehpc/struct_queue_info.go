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

// QueueInfo is a nested struct in ehpc response
type QueueInfo struct {
	QueueName           string        `json:"QueueName" xml:"QueueName"`
	ResourceGroupId     string        `json:"ResourceGroupId" xml:"ResourceGroupId"`
	SpotPriceLimit      float64       `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
	EnableAutoGrow      bool          `json:"EnableAutoGrow" xml:"EnableAutoGrow"`
	EnableAutoShrink    bool          `json:"EnableAutoShrink" xml:"EnableAutoShrink"`
	MinNodesInQueue     int           `json:"MinNodesInQueue" xml:"MinNodesInQueue"`
	InstanceType        string        `json:"InstanceType" xml:"InstanceType"`
	SpotStrategy        string        `json:"SpotStrategy" xml:"SpotStrategy"`
	ComputeInstanceType string        `json:"ComputeInstanceType" xml:"ComputeInstanceType"`
	MaxNodesInQueue     int           `json:"MaxNodesInQueue" xml:"MaxNodesInQueue"`
	Type                string        `json:"Type" xml:"Type"`
	InstanceTypes       InstanceTypes `json:"InstanceTypes" xml:"InstanceTypes"`
}
