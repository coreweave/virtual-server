// +build !ignore_autogenerated

/*
Copyright 2020.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	apiv1 "kubevirt.io/client-go/api/v1"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServer) DeepCopyInto(out *VirtualServer) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServer.
func (in *VirtualServer) DeepCopy() *VirtualServer {
	if in == nil {
		return nil
	}
	out := new(VirtualServer)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServer) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerFilesystem) DeepCopyInto(out *VirtualServerFilesystem) {
	*out = *in
	in.VirtualServerStorageVolume.DeepCopyInto(&out.VirtualServerStorageVolume)
	if in.Mountpoint != nil {
		in, out := &in.Mountpoint, &out.Mountpoint
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerFilesystem.
func (in *VirtualServerFilesystem) DeepCopy() *VirtualServerFilesystem {
	if in == nil {
		return nil
	}
	out := new(VirtualServerFilesystem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerFloatingIP) DeepCopyInto(out *VirtualServerFloatingIP) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerFloatingIP.
func (in *VirtualServerFloatingIP) DeepCopy() *VirtualServerFloatingIP {
	if in == nil {
		return nil
	}
	out := new(VirtualServerFloatingIP)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerList) DeepCopyInto(out *VirtualServerList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]VirtualServer, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerList.
func (in *VirtualServerList) DeepCopy() *VirtualServerList {
	if in == nil {
		return nil
	}
	out := new(VirtualServerList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *VirtualServerList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerNetwork) DeepCopyInto(out *VirtualServerNetwork) {
	*out = *in
	if in.FloatingIPs != nil {
		in, out := &in.FloatingIPs, &out.FloatingIPs
		*out = make([]VirtualServerFloatingIP, len(*in))
		copy(*out, *in)
	}
	in.TCP.DeepCopyInto(&out.TCP)
	in.UDP.DeepCopyInto(&out.UDP)
	if in.DNSConfig != nil {
		in, out := &in.DNSConfig, &out.DNSConfig
		*out = new(v1.PodDNSConfig)
		(*in).DeepCopyInto(*out)
	}
	if in.DNSPolicy != nil {
		in, out := &in.DNSPolicy, &out.DNSPolicy
		*out = new(v1.DNSPolicy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerNetwork.
func (in *VirtualServerNetwork) DeepCopy() *VirtualServerNetwork {
	if in == nil {
		return nil
	}
	out := new(VirtualServerNetwork)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerNetworkStatus) DeepCopyInto(out *VirtualServerNetworkStatus) {
	*out = *in
	if in.InternalIP != nil {
		in, out := &in.InternalIP, &out.InternalIP
		*out = new(string)
		**out = **in
	}
	if in.ExternalIP != nil {
		in, out := &in.ExternalIP, &out.ExternalIP
		*out = new(string)
		**out = **in
	}
	if in.ServiceIP != nil {
		in, out := &in.ServiceIP, &out.ServiceIP
		*out = new(string)
		**out = **in
	}
	if in.FloatingIPs != nil {
		in, out := &in.FloatingIPs, &out.FloatingIPs
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerNetworkStatus.
func (in *VirtualServerNetworkStatus) DeepCopy() *VirtualServerNetworkStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualServerNetworkStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerOS) DeepCopyInto(out *VirtualServerOS) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerOS.
func (in *VirtualServerOS) DeepCopy() *VirtualServerOS {
	if in == nil {
		return nil
	}
	out := new(VirtualServerOS)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerResourceCPU) DeepCopyInto(out *VirtualServerResourceCPU) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerResourceCPU.
func (in *VirtualServerResourceCPU) DeepCopy() *VirtualServerResourceCPU {
	if in == nil {
		return nil
	}
	out := new(VirtualServerResourceCPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerResourceGPU) DeepCopyInto(out *VirtualServerResourceGPU) {
	*out = *in
	if in.Type != nil {
		in, out := &in.Type, &out.Type
		*out = new(string)
		**out = **in
	}
	if in.Count != nil {
		in, out := &in.Count, &out.Count
		*out = new(uint32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerResourceGPU.
func (in *VirtualServerResourceGPU) DeepCopy() *VirtualServerResourceGPU {
	if in == nil {
		return nil
	}
	out := new(VirtualServerResourceGPU)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerResources) DeepCopyInto(out *VirtualServerResources) {
	*out = *in
	in.GPU.DeepCopyInto(&out.GPU)
	in.CPU.DeepCopyInto(&out.CPU)
	out.Memory = in.Memory.DeepCopy()
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerResources.
func (in *VirtualServerResources) DeepCopy() *VirtualServerResources {
	if in == nil {
		return nil
	}
	out := new(VirtualServerResources)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerServiceTemplate) DeepCopyInto(out *VirtualServerServiceTemplate) {
	*out = *in
	if in.Ports != nil {
		in, out := &in.Ports, &out.Ports
		*out = make([]Port, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerServiceTemplate.
func (in *VirtualServerServiceTemplate) DeepCopy() *VirtualServerServiceTemplate {
	if in == nil {
		return nil
	}
	out := new(VirtualServerServiceTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerSpec) DeepCopyInto(out *VirtualServerSpec) {
	*out = *in
	if in.Affinity != nil {
		in, out := &in.Affinity, &out.Affinity
		*out = new(v1.Affinity)
		(*in).DeepCopyInto(*out)
	}
	out.OS = in.OS
	in.Resources.DeepCopyInto(&out.Resources)
	in.Storage.DeepCopyInto(&out.Storage)
	if in.LivenessProbe != nil {
		in, out := &in.LivenessProbe, &out.LivenessProbe
		*out = new(apiv1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.ReadinessProbe != nil {
		in, out := &in.ReadinessProbe, &out.ReadinessProbe
		*out = new(apiv1.Probe)
		(*in).DeepCopyInto(*out)
	}
	if in.Users != nil {
		in, out := &in.Users, &out.Users
		*out = make([]VirtualServerUser, len(*in))
		copy(*out, *in)
	}
	in.Network.DeepCopyInto(&out.Network)
	if in.RunStrategy != nil {
		in, out := &in.RunStrategy, &out.RunStrategy
		*out = new(apiv1.VirtualMachineRunStrategy)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerSpec.
func (in *VirtualServerSpec) DeepCopy() *VirtualServerSpec {
	if in == nil {
		return nil
	}
	out := new(VirtualServerSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerStatus) DeepCopyInto(out *VirtualServerStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]metav1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	in.Network.DeepCopyInto(&out.Network)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerStatus.
func (in *VirtualServerStatus) DeepCopy() *VirtualServerStatus {
	if in == nil {
		return nil
	}
	out := new(VirtualServerStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerStorage) DeepCopyInto(out *VirtualServerStorage) {
	*out = *in
	in.Root.DeepCopyInto(&out.Root)
	if in.AdditionalDisks != nil {
		in, out := &in.AdditionalDisks, &out.AdditionalDisks
		*out = make([]VirtualServerStorageVolume, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.FileSystems != nil {
		in, out := &in.FileSystems, &out.FileSystems
		*out = make([]VirtualServerFilesystem, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Swap != nil {
		in, out := &in.Swap, &out.Swap
		x := (*in).DeepCopy()
		*out = &x
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerStorage.
func (in *VirtualServerStorage) DeepCopy() *VirtualServerStorage {
	if in == nil {
		return nil
	}
	out := new(VirtualServerStorage)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerStorageRoot) DeepCopyInto(out *VirtualServerStorageRoot) {
	*out = *in
	out.Size = in.Size.DeepCopy()
	in.Source.DeepCopyInto(&out.Source)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerStorageRoot.
func (in *VirtualServerStorageRoot) DeepCopy() *VirtualServerStorageRoot {
	if in == nil {
		return nil
	}
	out := new(VirtualServerStorageRoot)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerStorageRootHTTPSource) DeepCopyInto(out *VirtualServerStorageRootHTTPSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerStorageRootHTTPSource.
func (in *VirtualServerStorageRootHTTPSource) DeepCopy() *VirtualServerStorageRootHTTPSource {
	if in == nil {
		return nil
	}
	out := new(VirtualServerStorageRootHTTPSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerStorageRootPVCSource) DeepCopyInto(out *VirtualServerStorageRootPVCSource) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerStorageRootPVCSource.
func (in *VirtualServerStorageRootPVCSource) DeepCopy() *VirtualServerStorageRootPVCSource {
	if in == nil {
		return nil
	}
	out := new(VirtualServerStorageRootPVCSource)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerStorageVolume) DeepCopyInto(out *VirtualServerStorageVolume) {
	*out = *in
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerStorageVolume.
func (in *VirtualServerStorageVolume) DeepCopy() *VirtualServerStorageVolume {
	if in == nil {
		return nil
	}
	out := new(VirtualServerStorageVolume)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *VirtualServerUser) DeepCopyInto(out *VirtualServerUser) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new VirtualServerUser.
func (in *VirtualServerUser) DeepCopy() *VirtualServerUser {
	if in == nil {
		return nil
	}
	out := new(VirtualServerUser)
	in.DeepCopyInto(out)
	return out
}
