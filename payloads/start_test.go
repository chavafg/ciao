/*
// Copyright (c) 2016 Intel Corporation
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
*/

package payloads

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"testing"
)

func TestStartUnmarshal(t *testing.T) {
	startYaml := `start:
  instance_uuid: 3390740c-dce9-48d6-b83a-a717417072ce
  image_uuid: 59460b8a-5f53-4e3e-b5ce-b71fed8c7e64
  fw_type: efi
  persistence: host
  vm_type: qemu
  requested_resources:
    - type: vcpus
      value: 2
      mandatory: true
    - type: mem_mb
      value: 1014
      mandatory: true
    - type: disk_mb
      value: 10000
      mandatory: true
  estimated_resources:
    - type: vcpus
      value: 1
    - type: mem_mb
      value: 128
    - type: disk_mb
      value: 4096
`
	var cmd Start
	err := yaml.Unmarshal([]byte(startYaml), &cmd)
	if err != nil {
		t.Error(err)
	}

	fmt.Printf("Instance UUID [%s]\n", cmd.Start.InstanceUUID)
}

func TestStartMarshal(t *testing.T) {
	reqVcpus := RequestedResource{
		Type:      "vcpus",
		Value:     2,
		Mandatory: true,
	}
	reqMem := RequestedResource{
		Type:      "mem_mb",
		Value:     4096,
		Mandatory: true,
	}
	reqDisk := RequestedResource{
		Type:      "disk_mb",
		Value:     10000,
		Mandatory: true,
	}
	estVcpus := EstimatedResource{
		Type:  "vcpus",
		Value: 1,
	}
	estMem := EstimatedResource{
		Type:  "mem_mb",
		Value: 128,
	}
	estDisk := EstimatedResource{
		Type:  "disk_mb",
		Value: 4096,
	}
	var cmd Start
	cmd.Start.InstanceUUID = "c73322e8-d5fe-4d57-874c-dcee4fd368cd"
	cmd.Start.ImageUUID = "b265f62b-e957-47fd-a0a2-6dc261c7315c"
	cmd.Start.DockerImage = "ubuntu/latest"
	cmd.Start.RequestedResources = append(cmd.Start.RequestedResources, reqVcpus)
	cmd.Start.RequestedResources = append(cmd.Start.RequestedResources, reqMem)
	cmd.Start.RequestedResources = append(cmd.Start.RequestedResources, reqDisk)
	cmd.Start.EstimatedResources = append(cmd.Start.EstimatedResources, estVcpus)
	cmd.Start.EstimatedResources = append(cmd.Start.EstimatedResources, estMem)
	cmd.Start.EstimatedResources = append(cmd.Start.EstimatedResources, estDisk)
	cmd.Start.FWType = EFI
	cmd.Start.InstancePersistence = Host
	cmd.Start.VMType = QEMU

	y, err := yaml.Marshal(&cmd)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(string(y))
}

// make sure the yaml can be unmarshaled into the Start struct with
// optional data not present
func TestStartUnmarshalPartial(t *testing.T) {
	startYaml := `start:
  instance_uuid: 923d1f2b-aabe-4a9b-9982-8664b0e52f93
  image_uuid: 53cdd9ef-228f-4ce1-911d-706c2b41454a
  docker_image: ubuntu/latest
  fw_type: efi
  persistence: host
  vm_type: qemu
  requested_resources:
    - type: vcpus
      value: 2
      mandatory: true
`
	var cmd Start
	err := yaml.Unmarshal([]byte(startYaml), &cmd)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(cmd)

	var expectedCmd Start
	expectedCmd.Start.InstanceUUID = "923d1f2b-aabe-4a9b-9982-8664b0e52f93"
	expectedCmd.Start.ImageUUID = "53cdd9ef-228f-4ce1-911d-706c2b41454a"
	expectedCmd.Start.DockerImage = "ubuntu/latest"
	expectedCmd.Start.FWType = EFI
	expectedCmd.Start.InstancePersistence = Host
	expectedCmd.Start.VMType = QEMU
	vcpus := RequestedResource{
		Type:      "vcpus",
		Value:     2,
		Mandatory: true,
	}
	expectedCmd.Start.RequestedResources = append(expectedCmd.Start.RequestedResources, vcpus)

	if cmd.Start.InstanceUUID != expectedCmd.Start.InstanceUUID ||
		cmd.Start.ImageUUID != expectedCmd.Start.ImageUUID ||
		cmd.Start.DockerImage != expectedCmd.Start.DockerImage ||
		cmd.Start.FWType != expectedCmd.Start.FWType ||
		cmd.Start.InstancePersistence != expectedCmd.Start.InstancePersistence ||
		cmd.Start.VMType != expectedCmd.Start.VMType ||
		len(cmd.Start.RequestedResources) != 1 ||
		len(expectedCmd.Start.RequestedResources) != 1 ||
		cmd.Start.RequestedResources[0].Type != expectedCmd.Start.RequestedResources[0].Type ||
		cmd.Start.RequestedResources[0].Value != expectedCmd.Start.RequestedResources[0].Value ||
		cmd.Start.RequestedResources[0].Mandatory != expectedCmd.Start.RequestedResources[0].Mandatory {
		t.Error("Unexpected values in Start")
	}

	fmt.Println(cmd)
}
