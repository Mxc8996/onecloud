// Copyright 2019 Yunion
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

package google

import (
	"yunion.io/x/jsonutils"

	api "yunion.io/x/onecloud/pkg/apis/compute"
)

type SGlobalNetwork struct {
	Id string
	//CreationTimestamp     time.Time
	Name                  string
	Description           string
	SelfLink              string
	AutoCreateSubnetworks bool
	Subnetworks           []string
	RoutingConfig         map[string]string
	Kind                  string
}

func (cli *SGoogleClient) GetGlobalNetwork(id string) (*SGlobalNetwork, error) {
	net := &SGlobalNetwork{}
	return net, cli.get(id, net)
}

func (cli *SGoogleClient) GetGlobalNetworks(maxResults int, pageToken string) ([]SGlobalNetwork, error) {
	networks := []SGlobalNetwork{}
	params := map[string]string{}
	resource := "global/networks"
	return networks, cli.list(resource, params, maxResults, pageToken, &networks)
}

func (net *SGlobalNetwork) GetId() string {
	return net.SelfLink
}

func (net *SGlobalNetwork) GetGlobalId() string {
	return getGlobalId(net.SelfLink)
}

func (net *SGlobalNetwork) GetName() string {
	return net.Name
}

func (net *SGlobalNetwork) GetMetadata() *jsonutils.JSONDict {
	return nil
}

func (net *SGlobalNetwork) GetStatus() string {
	return api.GLOBAL_NETWORK_STATUS_AVAILABLE
}

func (net *SGlobalNetwork) IsEmulated() bool {
	return false
}

func (net *SGlobalNetwork) Refresh() error {
	return nil
}