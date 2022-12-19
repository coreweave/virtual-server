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
	"k8s.io/apimachinery/pkg/types"
	kvv1 "kubevirt.io/api/core/v1"
	cdiv1beta "kubevirt.io/containerized-data-importer-api/pkg/apis/core/v1beta1"
)

// VirtualServerSpec defines the desired state of VirtualServer
type VirtualServerSpec struct {
	Region    string                 `json:"region,omitempty"`
	Affinity  *corev1.Affinity       `json:"affinity,omitempty"`
	OS        VirtualServerOS        `json:"os"`
	Resources VirtualServerResources `json:"resources"`
	Storage   VirtualServerStorage   `json:"storage"`
	// +optional
	LivenessProbe *kvv1.Probe `json:"livenessProbe,omitempty"`
	// +optional
	ReadinessProbe *kvv1.Probe `json:"readinessProbe,omitempty"`
	// +optional
	Users []VirtualServerUser `json:"users,omitempty"`
	// +optional
	Network VirtualServerNetwork `json:"network"`
	// +optional
	InitializeRunning bool `json:"initializeRunning,omitempty"`
	// +optional
	CloudInit string `json:"cloudInit,omitempty"`
	// +kubebuilder:validation:Enum=Always;RerunOnFailure;Manual;Halted
	RunStrategy *kvv1.VirtualMachineRunStrategy `json:"runStrategy,omitempty"`
	// +optional
	Firmware Firmware `json:"firmware,omitempty"`
	// +optional
	UseVirtioTransitional         *bool  `json:"useVirtioTransitional,omitempty"`
	TerminationGracePeriodSeconds *int64 `json:"terminationGracePeriodSeconds,omitempty"`
}

type Firmware struct {
	// UUID reported by the vmi bios.
	// Defaults to a random generated uid.
	// +optional
	UUID types.UID `json:"uuid,omitempty"`
	// The system-serial-number in SMBIOS
	// +optional
	Serial string `json:"serial,omitempty"`
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
// +kubebuilder:printcolumn:JSONPath=".status.network.internalIP",name=Internal IP,type=string
// +kubebuilder:printcolumn:JSONPath=".status.network.externalIP",name=External IP,type=string

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
	// Configure the Virtual Server use a UEFI bootloader
	// +optional
	EnableUEFIBoot bool `json:"enableUEFIBoot,omitempty"`
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
	// +optional
	// +kubebuilder:validation:Minimum=1
	Count *uint32 `json:"count,omitempty"`
}

// VirtualServerStorage describes the Storage request for the VirtualServer
type VirtualServerStorage struct {
	// Root describes the root filesystem of the VirtualServer
	Root VirtualServerStorageRoot `json:"root"`
	// AdditionalDisks is an array of disks devices added to the VirtualServer
	AdditionalDisks []VirtualServerDisks `json:"additionalDisks,omitempty"`
	// Filesystems is an array of filesystem mounted to the VirtualServer
	FileSystems []VirtualServerFilesystem `json:"filesystems,omitempty"`
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
	Source *cdiv1beta.DataVolumeSource `json:"source"`
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
	// Ephemeral, if true, will disable disk persistence for the root filesystem.
	// A local image will be used to write changes, and will be discared when the Virtual Server is stopped or restarted.
	// Only a PVC source may be specified
	Ephemeral bool `json:"ephemeral,omitempty"`
	// Disk serial number
	// +optional
	Serial string `json:"serial,omitempty"`
}

// VirtualServerStorageVolume describes a named volume in the VirtualServer
type VirtualServerStorageVolume struct {
	Name string            `json:"name"`
	Spec kvv1.VolumeSource `json:"spec"`
}

// DiskAttributes describes disk sttributes
type DiskAttributes struct {
	// ReadOnly
	// +optional
	ReadOnly bool `json:"readOnly,omitempty"`
	// Disk serial number
	// +optional
	Serial string `json:"serial,omitempty"`
}

type VirtualServerDisks struct {
	VirtualServerStorageVolume `json:",inline"`
	DiskAttributes             `json:",inline"`
}

type VirtualServerFilesystem struct {
	VirtualServerStorageVolume `json:",inline"`
	// +optional
	Mountpoint *string `json:"mountPoint,omitempty"`
}

// VirtualServerUser defines user login information in the VirtualServer
// The user login information will be used to configure the VirtualServer via cloudinit if supported
type VirtualServerUser struct {
	Username     string `json:"username"`
	Password     string `json:"password,omitempty"`
	SSHPublicKey string `json:"sshpublickey,omitempty"`
}

// GetSize returns total size of fields in VirtualServerUser struct
func (vsu *VirtualServerUser) GetSize() int {
	var totalSize int = 0
	totalSize += len(vsu.Username)
	totalSize += len(vsu.Password)
	return totalSize
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
	// DNSConfig defines the DNS parameters of a VMI in addition to those generated from DNSPolicy.
	// +optional
	DNSConfig *corev1.PodDNSConfig `json:"dnsConfig,omitempty"`
	// Set DNS policy for the VMI. Defaults to "ClusterFirst".
	// Valid values are 'ClusterFirstWithHostNet', 'ClusterFirst', 'Default' or 'None'.
	// +optional
	// +kubebuilder:validation:Enum=ClusterFirstWithHostNet;ClusterFirst;Default;None
	DNSPolicy *corev1.DNSPolicy `json:"dnsPolicy,omitempty"`
	// Set MAC address for the VMI. It must be a local unicast type.
	// +optional
	// +kubebuilder:validation:Pattern="^[0-9a-f][26ae][:]([0-9a-f]{2}[:]){4}([0-9a-f]{2})|[0-9A-F][26AE][-]([0-9A-F]{2}[-]){4}([0-9A-F]{2})$"
	MACAddress string `json:"macAddress,omitempty"`
	// When DirectAttachLoadBalancerIP is false or no ports are specified, create a headless service. Defaults to false.
	// +optional
	// +kubebuilder:default=false
	Headless bool `json:"headless,omitempty"`
	// List of VPC networks
	// +optional
	VPCs []VirtualServerVPC `json:"vpcs,omitempty"`
	// Disable kubernetes pod network within the Virtual Server
	// Useful for isolating a Virtual Server in VPC networks
	DisableK8sNetworking bool `json:"disableK8sNetworking,omitempty"`
}

// VirtualServerVPC defines a VPC network for the Virtual Server to join
type VirtualServerVPC struct {
	Name string `json:"name"`
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
	ServiceName string `json:"serviceName"`
}

type VirtualServerNetworkStatus struct {
	InternalIP  *string           `json:"internalIP,omitempty"`
	ExternalIP  *string           `json:"externalIP,omitempty"`
	ServiceIP   *string           `json:"serviceIP,omitempty"`
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
	// VSConditionTypeServicesReady describes the ready state of the services dynamically created and/or those required by the VirtualServer
	VSConditionTypeSecretReady VirtualServerConditionType = "SecretReady"
)

type VirtualServerConditionReason string

const (
	// VSConditionReasonInitializing indicates that the VirtualServer is initializing for the first time
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
	// VSConditionReasonVMIShutdown indicates that the Virtual Machine Instance has been shut down
	VSConditionReasonVMIShutdown VirtualServerConditionReason = "VirtualMachineInstanceShutdown"
	// VSConditionReasonDefinitionDeprecated indicates that the definition used to configure the VirtualServer has been deprecated and a manual update is required
	VSConditionReasonDefinitionDeprecated VirtualServerConditionReason = "DefinitionDeprecated"
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
	// VSConditionReasonSecretCreated indicates that the VirtualServer secret have been successfully created
	VSConditionReasonSecretCreated VirtualServerConditionReason = "SecretCreated"
	// VSConditionReasonWaitingForSecret indicates that the VirtualServer is waiting for the required secret to be ready
	VSConditionReasonWaitingForSecrets VirtualServerConditionReason = "WaitingForSecret"
	// VSConditionReasonResizeInProgress indicates that the VirtualServer root disk is being resized
	VSConditionReasonResizeInProgress VirtualServerConditionReason = "RootDiskResizeinProgress"
)
