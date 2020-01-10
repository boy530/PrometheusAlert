package teslamaxcompute

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

// Instance is a nested struct in teslamaxcompute response
type Instance struct {
	User            string  `json:"User" xml:"User"`
	SkynetId        string  `json:"SkynetId" xml:"SkynetId"`
	CpuUsedRatioMin float64 `json:"CpuUsedRatioMin" xml:"CpuUsedRatioMin"`
	CpuRequest      int     `json:"CpuRequest" xml:"CpuRequest"`
	MemUsedRatioMin float64 `json:"MemUsedRatioMin" xml:"MemUsedRatioMin"`
	QuotaId         int     `json:"QuotaId" xml:"QuotaId"`
	RunTime         string  `json:"RunTime" xml:"RunTime"`
	Project         string  `json:"Project" xml:"Project"`
	MemUsedRatioMax float64 `json:"MemUsedRatioMax" xml:"MemUsedRatioMax"`
	InstanceId      string  `json:"InstanceId" xml:"InstanceId"`
	NickName        string  `json:"NickName" xml:"NickName"`
	CpuUsedRatioMax float64 `json:"CpuUsedRatioMax" xml:"CpuUsedRatioMax"`
	MemUsedTotal    int     `json:"MemUsedTotal" xml:"MemUsedTotal"`
	ProjectOwner    string  `json:"ProjectOwner" xml:"ProjectOwner"`
	QuotaName       string  `json:"QuotaName" xml:"QuotaName"`
	MemRequest      int     `json:"MemRequest" xml:"MemRequest"`
	CollectTime     string  `json:"CollectTime" xml:"CollectTime"`
	Cluster         string  `json:"Cluster" xml:"Cluster"`
	CpuUsed         int     `json:"CpuUsed" xml:"CpuUsed"`
	IsRealOwner     string  `json:"IsRealOwner" xml:"IsRealOwner"`
	Status          string  `json:"Status" xml:"Status"`
	TaskType        string  `json:"TaskType" xml:"TaskType"`
	UserAccount     string  `json:"UserAccount" xml:"UserAccount"`
	MemUsed         int     `json:"MemUsed" xml:"MemUsed"`
	CpuUsedTotal    int     `json:"CpuUsedTotal" xml:"CpuUsedTotal"`
}