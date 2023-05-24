// Copyright Istio Authors
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

package proxy

import (
	"context"
	"encoding/json"
	"fmt"

	"istio.io/istio/pilot/pkg/model"
	"istio.io/istio/pilot/pkg/xds"
	"istio.io/istio/pkg/kube"
)

type sidecarSyncStatus struct {
	// nolint: structcheck, unused
	pilot string
	xds.SyncStatus
}

type NodeType string

// Info contains the version for a single data plane component
type Info struct {
	ID           string
	IstioVersion string
	Type         NodeType
}

func ToUserFacingNodeType(t model.NodeType) NodeType {
	switch t {
	case model.Router:
		return "gateway"
	case model.Ztunnel, model.Waypoint, model.SidecarProxy:
		return NodeType(t)
	}
	return ""
}

func (n NodeType) EqualsModelNodeType(t model.NodeType) bool {
	return n == ToUserFacingNodeType(t)
}

// GetProxyInfo retrieves infos of proxies that connect to the Istio control plane of specific revision.
func GetProxyInfo(kubeconfig, configContext, revision, istioNamespace string) (*[]Info, error) {
	kubeClient, err := kube.NewCLIClient(kube.BuildClientCmd(kubeconfig, configContext), revision)
	if err != nil {
		return nil, err
	}
	// Ask Pilot for the Envoy sidecar sync status, which includes the sidecar version info
	allSyncz, err := kubeClient.AllDiscoveryDo(context.TODO(), istioNamespace, "debug/syncz")
	if err != nil {
		return nil, err
	}

	pi := make([]Info, 0)
	for _, syncz := range allSyncz {
		var sss []*sidecarSyncStatus
		err = json.Unmarshal(syncz, &sss)
		if err != nil {
			return nil, err
		}

		for _, ss := range sss {
			pi = append(pi, Info{
				ID:           ss.ProxyID,
				IstioVersion: ss.SyncStatus.IstioVersion,
				Type:         ToUserFacingNodeType(ss.ProxyType),
			})
		}
	}

	return &pi, nil
}

// GetIDsFromProxyInfo is a helper function to retrieve list of IDs from Proxy.
func GetIDsFromProxyInfo(kubeconfig, configContext, revision, istioNamespace string) ([]string, error) {
	var IDs []string
	pi, err := GetProxyInfo(kubeconfig, configContext, revision, istioNamespace)
	if err != nil {
		return IDs, fmt.Errorf("failed to get proxy infos: %v", err)
	}
	for _, pi := range *pi {
		IDs = append(IDs, pi.ID)
	}
	return IDs, nil
}
