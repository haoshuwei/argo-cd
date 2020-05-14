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

// DBInstanceAttributeInDescribeDBInstancesAsCsv is a nested struct in rds response
type DBInstanceAttributeInDescribeDBInstancesAsCsv struct {
	DBInstanceId                string `json:"DBInstanceId" xml:"DBInstanceId"`
	PayType                     string `json:"PayType" xml:"PayType"`
	DBInstanceClassType         string `json:"DBInstanceClassType" xml:"DBInstanceClassType"`
	DBInstanceType              string `json:"DBInstanceType" xml:"DBInstanceType"`
	RegionId                    string `json:"RegionId" xml:"RegionId"`
	ConnectionString            string `json:"ConnectionString" xml:"ConnectionString"`
	Port                        string `json:"Port" xml:"Port"`
	Engine                      string `json:"Engine" xml:"Engine"`
	EngineVersion               string `json:"EngineVersion" xml:"EngineVersion"`
	DBInstanceClass             string `json:"DBInstanceClass" xml:"DBInstanceClass"`
	DBInstanceMemory            int64  `json:"DBInstanceMemory" xml:"DBInstanceMemory"`
	DBInstanceStorage           int    `json:"DBInstanceStorage" xml:"DBInstanceStorage"`
	DBInstanceNetType           string `json:"DBInstanceNetType" xml:"DBInstanceNetType"`
	DBInstanceStatus            string `json:"DBInstanceStatus" xml:"DBInstanceStatus"`
	DBInstanceDescription       string `json:"DBInstanceDescription" xml:"DBInstanceDescription"`
	LockMode                    string `json:"LockMode" xml:"LockMode"`
	LockReason                  string `json:"LockReason" xml:"LockReason"`
	ReadDelayTime               string `json:"ReadDelayTime" xml:"ReadDelayTime"`
	DBMaxQuantity               int    `json:"DBMaxQuantity" xml:"DBMaxQuantity"`
	AccountMaxQuantity          int    `json:"AccountMaxQuantity" xml:"AccountMaxQuantity"`
	CreationTime                string `json:"CreationTime" xml:"CreationTime"`
	ExpireTime                  string `json:"ExpireTime" xml:"ExpireTime"`
	MaintainTime                string `json:"MaintainTime" xml:"MaintainTime"`
	AvailabilityValue           string `json:"AvailabilityValue" xml:"AvailabilityValue"`
	MaxIOPS                     int    `json:"MaxIOPS" xml:"MaxIOPS"`
	MaxConnections              int    `json:"MaxConnections" xml:"MaxConnections"`
	MasterInstanceId            string `json:"MasterInstanceId" xml:"MasterInstanceId"`
	DBInstanceCPU               string `json:"DBInstanceCPU" xml:"DBInstanceCPU"`
	IncrementSourceDBInstanceId string `json:"IncrementSourceDBInstanceId" xml:"IncrementSourceDBInstanceId"`
	GuardDBInstanceId           string `json:"GuardDBInstanceId" xml:"GuardDBInstanceId"`
	TempDBInstanceId            string `json:"TempDBInstanceId" xml:"TempDBInstanceId"`
	SecurityIPList              string `json:"SecurityIPList" xml:"SecurityIPList"`
	ZoneId                      string `json:"ZoneId" xml:"ZoneId"`
	InstanceNetworkType         string `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	Category                    string `json:"Category" xml:"Category"`
	AccountType                 string `json:"AccountType" xml:"AccountType"`
	SupportUpgradeAccountType   string `json:"SupportUpgradeAccountType" xml:"SupportUpgradeAccountType"`
	VpcId                       string `json:"VpcId" xml:"VpcId"`
	VSwitchId                   string `json:"VSwitchId" xml:"VSwitchId"`
	ConnectionMode              string `json:"ConnectionMode" xml:"ConnectionMode"`
	Tags                        string `json:"Tags" xml:"Tags"`
}
