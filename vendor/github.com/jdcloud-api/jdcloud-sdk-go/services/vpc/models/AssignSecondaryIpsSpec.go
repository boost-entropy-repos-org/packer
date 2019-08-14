// Copyright 2018 JDCLOUD.COM
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//
// NOTE: This class is auto generated by the jdcloud code generator program.

package models


type AssignSecondaryIpsSpec struct {

    /* secondary ip被其他接口占用时，是否抢占。false：非抢占重分配，true：抢占重分配，默认抢占重分配。默认值：true (Optional) */
    Force bool `json:"force"`

    /* 指定分配的secondaryIp地址 (Optional) */
    SecondaryIps []string `json:"secondaryIps"`

    /* 指定自动分配的secondaryIp个数 (Optional) */
    SecondaryIpCount int `json:"secondaryIpCount"`
}