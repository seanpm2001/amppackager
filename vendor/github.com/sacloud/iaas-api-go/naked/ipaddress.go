// Copyright 2022 The sacloud/iaas-api-go Authors
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package naked

// IPAddress IPアドレス(IPv4)
type IPAddress struct {
	HostName  string     `yaml:"host_name"`
	IPAddress string     `json:",omitempty" yaml:"ip_address,omitempty" structs:",omitempty"`
	Interface *Interface `json:",omitempty" yaml:"interface,omitempty" structs:",omitempty"`
	Subnet    *Subnet    `json:",omitempty" yaml:"subnet,omitempty" structs:",omitempty"`
}