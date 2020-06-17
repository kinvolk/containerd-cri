/*
Copyright 2017-2020 The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package server

import (
	"golang.org/x/net/context"

	runtimespec "github.com/opencontainers/runtime-spec/specs-go"
	runtime "k8s.io/cri-api/pkg/apis/runtime/v1alpha2"
)

var UsernsMapping = &runtimespec.LinuxIDMapping{
	ContainerID: uint32(0),
	HostID:      uint32(100000),
	Size:        uint32(65535),
}

func (c *criService) GetRuntimeConfigInfo(_ context.Context, r *runtime.GetRuntimeConfigInfoRequest) (res *runtime.GetRuntimeConfigInfoResponse, err error) {
	// Mapping used when we just use the host user namespace
	// See /proc/self/uid_map on the host
	//hostMapping := &runtime.LinuxIDMapping{
	//	ContainerId: uint32(0),
	//	HostId:      uint32(0),
	//	Size_:       uint32(4294967295),
	//}

	// Example of mapping we can use in containers
	linuxConfig := &runtime.LinuxUserNamespaceConfig{
		UidMappings: []*runtime.LinuxIDMapping{
			&runtime.LinuxIDMapping{
				ContainerId: uint32(0),
				HostId:      uint32(100000),
				Size_:       uint32(65535),
			},
		},
		GidMappings: []*runtime.LinuxIDMapping{
			&runtime.LinuxIDMapping{
				ContainerId: uint32(0),
				HostId:      uint32(100000),
				Size_:       uint32(65535),
			},
		},
	}
	activeRuntimeConfig := &runtime.ActiveRuntimeConfig{UserNamespaceConfig: linuxConfig}
	return &runtime.GetRuntimeConfigInfoResponse{RuntimeConfig: activeRuntimeConfig}, nil
}
