package itaas

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

// RegisterHistoryInfo is a nested struct in itaas response
type RegisterHistoryInfo struct {
	CreateTimeL  int    `json:"CreateTimeL" xml:"CreateTimeL"`
	DrIp         string `json:"DrIp" xml:"DrIp"`
	DrMac        string `json:"DrMac" xml:"DrMac"`
	DrName       string `json:"DrName" xml:"DrName"`
	Eventinfo    string `json:"Eventinfo" xml:"Eventinfo"`
	Eventtype    int    `json:"Eventtype" xml:"Eventtype"`
	EventtypeTxt string `json:"EventtypeTxt" xml:"EventtypeTxt"`
	Memo         string `json:"Memo" xml:"Memo"`
	Screencode   string `json:"Screencode" xml:"Screencode"`
}
