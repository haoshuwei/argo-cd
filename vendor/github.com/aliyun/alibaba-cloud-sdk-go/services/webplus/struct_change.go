package webplus

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

// Change is a nested struct in webplus response
type Change struct {
	ChangeMessage     string `json:"ChangeMessage" xml:"ChangeMessage"`
	FinishTime        int64  `json:"FinishTime" xml:"FinishTime"`
	UpdateTime        int64  `json:"UpdateTime" xml:"UpdateTime"`
	CreateUsername    string `json:"CreateUsername" xml:"CreateUsername"`
	ChangeSucceed     bool   `json:"ChangeSucceed" xml:"ChangeSucceed"`
	ChangePaused      bool   `json:"ChangePaused" xml:"ChangePaused"`
	CreateTime        int64  `json:"CreateTime" xml:"CreateTime"`
	ActionName        string `json:"ActionName" xml:"ActionName"`
	ChangeAborted     bool   `json:"ChangeAborted" xml:"ChangeAborted"`
	ChangeDescription string `json:"ChangeDescription" xml:"ChangeDescription"`
	ChangeTimedout    bool   `json:"ChangeTimedout" xml:"ChangeTimedout"`
	ChangeFinished    bool   `json:"ChangeFinished" xml:"ChangeFinished"`
	ChangeName        string `json:"ChangeName" xml:"ChangeName"`
	ChangeId          string `json:"ChangeId" xml:"ChangeId"`
	EnvId             string `json:"EnvId" xml:"EnvId"`
	ChangeSucceeded   bool   `json:"ChangeSucceeded" xml:"ChangeSucceeded"`
}
