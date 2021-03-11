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

// Package v1alpha1 defines the VirtualServer types and associated functions.
// The VirtualServer facilitates the creation and management of Virtual Server instances on the Coreweave Cloud kubernetes platform.
//
// Examples on creation and management of VirtualServers are available in the Examples section.
package v1alpha1

import (
	corev1 "k8s.io/api/core/v1"
	resource "k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	kvv1 "kubevirt.io/client-go/api/v1"
	cdiv1alpha "kubevirt.io/containerized-data-importer/pkg/apis/core/v1alpha1"
)

// VirtualServerSpec defines the desired state of VirtualServer
type VirtualServerSpec struct {
	Region    string                 `json:"region,omitempty"`
	Affinity  *corev1.Affinity       `json:"affinity,omitempty"`
	OS        VirtualServerOS        `json:"os"`
	Resources VirtualServerResources `json:"resources"`
	Storage   VirtualServerStorage   `json:"storage"`
	Users     []VirtualServerUser    `json:"users"`
	// +optional
	Network VirtualServerNetwork `json:"network"`
	// +optional
	InitializeRunning bool `json:"initializeRunning,omitempty"`
}

// VirtualServerStatus defines the observed state of VirtualServer
type VirtualServerStatus struct {
	Conditions []metav1.Condition         `json:"conditions,omitempty"`
	Network    VirtualServerNetworkStatus `json:"network,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:shortName=vserver;vs
// +kubebuilder:subresource:status
// +kubebuilder:printcolumn:JSONPath=".status.conditions[0].reason",name=status,type=string
// +kubebuilder:printcolumn:JSONPath=".status.conditions[0].message",name=reason,type=string
// +kubebuilder:printcolumn:JSONPath=".status.conditions[3].status",name=started,type=string

// VirtualServer is the Schema for the virtualservers API.
// It allows for configuring a Virtual Server instance on the Coreweave Cloud kubernetes platform.
// The VirtualServer handles the creation and lifecycle of a VirtualMachine and Services.
type VirtualServer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   VirtualServerSpec   `json:"spec,omitempty"`
	Status VirtualServerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// VirtualServerList contains a list of VirtualServer
type VirtualServerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []VirtualServer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&VirtualServer{}, &VirtualServerList{})
}

// Virtual Server Types

// VirtualServerOS defines the Operating System of the VirtualServer
type VirtualServerOS struct {
	// The Operating System run in the Virtual Server
	// VirtualServerOSType may be "windows" or "linux"
	// +kubebuilder:validation:Enum=windows;linux
	Type VirtualServerOSType `json:"type"`
	// The operating system configuration definition for internal use
	// See https://docs.coreweave.com/virtual-desktop for details on which definition value best suits your configuration.
	// Defaults to "a"
	// +optional
	// +kubebuilder:default=a
	Definition string `json:"definition,omitempty"`
}

// VirtualServerResources defines the resources requested for the VirtualServer
type VirtualServerResources struct {
	// The resource configuration definition for internal use
	// See https://docs.coreweave.com/virtual-desktop for details on which definition value best suits your configuration.
	// Defaults to "a"
	// +optional
	// +kubebuilder:default=a
	Definition string `json:"definition,omitempty"`
	// GPU describes the GPU resource request
	// +optional
	GPU VirtualServerResourceGPU `json:"gpu,omitempty"`
	// CPU describes the CPU resource request
	// +optional
	// +kubebuilder:default={count: 2}
	CPU VirtualServerResourceCPU `json:"cpu,omitempty"`
	// Memory describes the memory resource request
	// +optional
	// +kubebuilder:default="8Gi"
	Memory resource.Quantity `json:"memory,omitempty"`
}

// VirtualServerResourceCPU describes the CPU request for the VirtualServer
type VirtualServerResourceCPU struct {
	// Type is the CPU type to request
	// See Coreweave Metadata API for available CPU types
	// +optional
	Type *string `json:"type,omitempty"`
	// The number of CPU cores to request
	// +optional
	// +kubebuilder:default=2
	// +kubebuilder:validation:Minimum=1
	Count uint32 `json:"count"`
}

// VirtualServerResourceGPU describes the GPU request for the VirtualServer
type VirtualServerResourceGPU struct {
	// Type is the GPU type to request
	// See Coreweave Metadata API for available GPU types
	Type *string `json:"type,omitempty"`
	// The number of GPUs to request.
	// Defaults to 1
	// +optional
	// +kubebuilder:default=1
	// +kubebuilder:validation:Minimum=1
	Count *uint32 `json:"count,omitempty"`
}

// VirtualServerStorage describes the Storage request for the VirtualServer
type VirtualServerStorage struct {
	// Root describes the root filesystem of the VirtualServer
	Root VirtualServerStorageRoot `json:"root"`
	// AdditionalDisks is an array of disks devices added to the VirtualServer
	AdditionalDisks []VirtualServerStorageVolume `json:"additionalDisks,omitempty"`
	// Filesystems is an array of filesystem mounted to the VirtualServer
	FileSystems []VirtualServerStorageVolume `json:"filesystems,omitempty"`
	// Swap describes a swap volume of the specified size added to the VirtualServer
	// An emptyDisk is created of the specified size to be used as the swap disk
	Swap *resource.Quantity `json:"swap,omitempty"`
}

// VirtualServerStorageRoot describes the Storage request for root filesystem of the VirtualServer
type VirtualServerStorageRoot struct {
	// Size specifies the root filesystem volume size
	Size resource.Quantity `json:"size"`
	// Source describes the DataVolumeSource for the root filesystem DataVolume
	// A DataVolume will be dynamically created alongside the VirtualServer, and the underlying PVC will be mounted as the root filesystem
	Source cdiv1alpha.DataVolumeSource `json:"source"`
	// StorageClassName specifies the StorageClassName of the root filesystem PVC
	StorageClassName string `json:"storageClassName"`
	// VolumeMode specifies the VolumeMode of the root filesystem PVC.
	// Defaults to Block
	// +kubebuilder:default=Block
	// +optional
	VolumeMode corev1.PersistentVolumeMode `json:"volumeMode,omitempty"`
	// AccessMode specifies the AccessMode of the root filesystem PVC.
	// Defaults to ReadWriteOnce
	// +kubebuilder:default=ReadWriteOnce
	// +optional
	AccessMode corev1.PersistentVolumeAccessMode `json:"accessMode,omitempty"`
}

// VirtualServerStorageVolume describes a named volume in the VirtualServer
type VirtualServerStorageVolume struct {
	Name string            `json:"name"`
	Spec kvv1.VolumeSource `json:"spec"`
}

// VirtualServerUser defines user login information in the VirtualServer
// The user login information will be used to configure the VirtualServer via cloudinit if supported
type VirtualServerUser struct {
	Username string `json:"username"`
	Password string `json:"password,omitempty"`
}

// VirtualServerNetwork defines the network configuration of the VirtualServer
type VirtualServerNetwork struct {
	// If enabled, a Service will be dynamically created, and its IP directly attached to the VirtualServer
	// DirectAttachLoadBalancerIP may not be set if UDP or TCP VirtualServerPorts are defined
	DirectAttachLoadBalancerIP bool `json:"directAttachLoadBalancerIP,omitempty"`
	// FloatingIPs is an array of LoadBalancer Services
	// The Services LoadBalancer IPs will be used for the floating IPs of the VirtualServer
	FloatingIPs []VirtualServerFloatingIP `json:"floatingIPs,omitempty"`
	// TCP describes a list of tcp ports that are exposed by the VirtualServer
	// A Service will be dynamically created and linked to the VirtualServer
	// A maximum of 10 ports may be defined
	TCP VirtualServerServiceTemplate `json:"tcp,omitempty"`
	// UDP describes a list of udp ports that are exposed by the VirtualServer
	// A Service will be dynamically created and linked to the VirtualServer
	// A maximum of 10 ports may be defined
	UDP VirtualServerServiceTemplate `json:"udp,omitempty"`
	// If Public is true a public IP will be assigned to the created Services
	// Defaults to true
	// +optional
	// +kubebuilder:default=true
	Public bool `json:"public"`
}

// VirtualServerServiceTemplate defines a service created by the VirtualServer
type VirtualServerServiceTemplate struct {
	// A list of ports.
	// The list is constrained to a maximum of 10 ports
	// +kubebuilder:validation:MaxItems=10
	Ports []Port `json:"ports,omitempty"`
}

// +kubebuilder:validation:Minimum=1
// +kubebuilder:validation:Maximum=65535
type Port int32

// VirtualServerFloatingIP represents a source that will be used for a VirtualServer floating IP
type VirtualServerFloatingIP struct {
	// The name of an existing LoadBalancer Service to use as the Floating IP source
	SericeName string `json:"serviceName"`
}

type VirtualServerNetworkStatus struct {
	InternalIP  *string           `json:"internalIP,omitempty"`
	TCP         *string           `json:"tcpIP,omitempty"`
	UDP         *string           `json:"udpIP,omitempty"`
	FloatingIPs map[string]string `json:"floatingIPs,omitempty"`
}

type VirtualServerOSType string

const (
	// VirtualServer Linux operating system
	VirtualServerOSTypeLinux VirtualServerOSType = "linux"
	// VirtualServer Windows operating system
	VirtualServerOSTypeWindows VirtualServerOSType = "windows"
)

type VirtualServerConditionType string

const (
	// VSConditionTypeReady describes the ready state of the Virtual Server
	VSConditionTypeReady VirtualServerConditionType = "Ready"
	// VSConditionTypeStarted describes whether the VirtualServer has been started
	VSConditionTypeStarted VirtualServerConditionType = "VirtualServerStarted"
	// VSConditionTypeServicesReady describes the ready state of the services dynamically created and/or those required by the VirtualServer
	VSConditionTypeServicesReady VirtualServerConditionType = "ServicesReady"
	// VSConditionTypeVMReady describes the ready state of the underlying VirtualMachine
	VSConditionTypeVMReady VirtualServerConditionType = "VirtualMachineReady"
)

type VirtualServerConditionReason string

const (
	// VSConditionReasonInitializing indicates that the VirtualServer is initilizing for the first time
	VSConditionReasonInitializing VirtualServerConditionReason = "Initializing"
	// VSConditionReasonPending indicates that the VirtualServer is pending an update to its spec
	VSConditionReasonPending VirtualServerConditionReason = "Pending"
	// VSConditionReasonTerminating indicates that a VirtualServer resource is terminating
	VSConditionReasonTerminating VirtualServerConditionReason = "Terminating"
	// VSConditionReasonFailed indicates that the VirtualServer was not able to create or start
	VSConditionReasonFailed VirtualServerConditionReason = "Failed"
	// VSConditionReasonReady indicates that the VirtualServer has successfully been created and is ready for use
	VSConditionReasonReady VirtualServerConditionReason = "VirtualServerReady"
	// VSConditionReasonStarted indicates that the VirtualServer has started
	VSConditionReasonStarted VirtualServerConditionReason = "VirtualServerStarted"
	// VSConditionReasonStopped indicates that the VirtualServer been stopped
	VSConditionReasonStopped VirtualServerConditionReason = "VirtualServerStopped"
	// VSConditionReasonDefinitionDepricated indicates that the definition used to configure the VirtualServer has been depricated and a manual update is required
	VSConditionReasonDefinitionDepricated VirtualServerConditionReason = "DefinitionDepricated"
	// VSConditionReasonServicesCreated indicates that the VirtualServer services have been successfully created
	VSConditionReasonServicesCreated VirtualServerConditionReason = "ServicesCreated"
	// VSConditionReasonWaitingForServices indicates that the VirtualServer is waiting for the required services to be ready
	VSConditionReasonWaitingForServices VirtualServerConditionReason = "WaitingForServices"
	// VSConditionReasonServicesReady indicates that the required services are ready
	VSConditionReasonServicesReady VirtualServerConditionReason = "ServicesReady"
	// VSConditionReasonVMNameTaken indicates that the name of the underlying VirtualMachine already exists or is owned by another VirtualServer and therefore could not be created
	VSConditionReasonVMNameTaken VirtualServerConditionReason = "VirtualMachineNameTaken"
	// VSConditionReasonVMReady indicates that the underlying VirtualMachine is ready
	VSConditionReasonVMReady VirtualServerConditionReason = "VirtualMachineReady"
)
