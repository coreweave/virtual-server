package v1alpha1

import (
	"errors"
	"fmt"
	"regexp"

	"gopkg.in/yaml.v2"
	corev1 "k8s.io/api/core/v1"
	apimeta "k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/api/resource"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/types"
	kvv1 "kubevirt.io/client-go/api/v1"
	cdiv1alpha "kubevirt.io/containerized-data-importer/pkg/apis/core/v1alpha1"
)

const MacAddressRegEx = `^[0-9a-f][26ae][:]([0-9a-f]{2}[:]){4}([0-9a-f]{2})|[0-9A-F][26AE][-]([0-9A-F]{2}[-]){4}([0-9A-F]{2})$`
const FirmwareSerialRegEx = `^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`

// IsGpuServer returns true if the VirtualServer is GPU enabled
func (vs *VirtualServer) IsGpuServer() bool {
	return vs.Spec.Resources.GPU.Type != nil
}

// SystemClass returns the VirtualServer system class.
// Classes are either "cpu" or "cpu"
func (vs *VirtualServer) SystemClass() SystemPresetClass {
	if vs.IsGpuServer() {
		return SystemPresetClassGPU
	} else {
		return SystemPresetClassCPU
	}
}

// SystemType returns the VirtualServer system type.
// Types are defined in the VirtualServer resources request under either GPU or CPU
func (vs *VirtualServer) SystemType() string {
	if vs.IsGpuServer() {
		return *vs.Spec.Resources.GPU.Type
	} else {
		return *vs.Spec.Resources.CPU.Type
	}
}

// Set the status condition of the virtualServer.
// If message is nil, the condition message will be set to a string casted form of reason.
// If applyToTopLevelCondition is true the status and message will be applied to the top level, VSConditionTypeReady, condition as well
func (vs *VirtualServer) SetCondition(
	conditionType VirtualServerConditionType,
	status metav1.ConditionStatus,
	reason VirtualServerConditionReason,
	message *string,
	applyToTopLevelCondition bool,
) {
	msg := string(reason)
	if message != nil {
		msg = *message
	}
	condition := metav1.Condition{
		Type:    string(conditionType),
		Status:  status,
		Reason:  string(reason),
		Message: msg,
	}

	if applyToTopLevelCondition {
		topLevelCondition := apimeta.FindStatusCondition(vs.Status.Conditions, string(VSConditionTypeReady))
		topLevelCondition.Status = condition.Status
		topLevelCondition.Message = condition.Message
	}

	apimeta.SetStatusCondition(&vs.Status.Conditions, condition)
}

// UpdateVirtualMachineStartedCondition will set the VirtualServer conditions that indicate that the underlying VirtualMachine has started
func (vs *VirtualServer) UpdateVirtualMachineStartedCondition(running bool) {
	var status metav1.ConditionStatus
	var reason VirtualServerConditionReason
	if running {
		status = metav1.ConditionTrue
		reason = VSConditionReasonStarted
	} else if apimeta.FindStatusCondition(vs.Status.Conditions, string(VSConditionTypeReady)).Status != metav1.ConditionTrue {
		status = metav1.ConditionUnknown
		reason = VSConditionReasonPending
	} else {
		status = metav1.ConditionFalse
		reason = VSConditionReasonStopped
	}
	vs.SetCondition(VSConditionTypeStarted, status, reason, nil, false)
}

// InitializeStatus sets the default VirtualServer status and conditions
func (vs *VirtualServer) InitializeStatus() {
	vs.Status.Network.FloatingIPs = make(map[string]string)
	vs.SetCondition(VSConditionTypeReady, metav1.ConditionUnknown, VSConditionReasonInitializing, nil, false)
	vs.SetCondition(VSConditionTypeServicesReady, metav1.ConditionUnknown, VSConditionReasonInitializing, nil, false)
	vs.SetCondition(VSConditionTypeVMReady, metav1.ConditionUnknown, VSConditionReasonInitializing, nil, false)
	vs.SetCondition(VSConditionTypeStarted, metav1.ConditionUnknown, VSConditionReasonInitializing, nil, false)
}

// HasNoConditions returns true if the VirtualServer has no conditions defined
func (vs *VirtualServer) HasNoConditions() bool {
	return len(vs.Status.Conditions) == 0
}

// Returns a VirtualServer with the provided name and namespace
func NewVirtualServer(name string, namespace string) *VirtualServer {
	return &VirtualServer{
		ObjectMeta: metav1.ObjectMeta{
			Name:      name,
			Namespace: namespace,
		},
	}
}

// Set the VirtualServer region
func (vs *VirtualServer) SetRegion(region string) {
	vs.Spec.Region = region
}

// Set the VirtualServer OS
func (vs *VirtualServer) SetOS(os VirtualServerOSType) {
	vs.Spec.OS = VirtualServerOS{
		Type: os,
	}
}

func (vs *VirtualServer) EnableUEFIBoot(enable bool) {
	vs.Spec.OS.EnableUEFIBoot = enable
}

// Set the VirtualServer resource definition
func (vs *VirtualServer) SetResourceDefinition(definitionVersion string) {
	vs.Spec.Resources.Definition = definitionVersion
}

// Set the VirtualServer CPU type
func (vs *VirtualServer) SetCPUType(cpuType string) error {
	if vs.Spec.Resources.GPU.Type != nil {
		return fmt.Errorf("CPU type cannot be set if GPU type is set")
	}
	vs.Spec.Resources.CPU.Type = &cpuType
	return nil
}

// Set the VirtualServer CPU core count
func (vs *VirtualServer) SetCPUCount(cpuCount uint32) {
	vs.Spec.Resources.CPU.Count = cpuCount
}

// Set the VirtualServer GPU type
func (vs *VirtualServer) SetGPUType(gpuType string) error {
	if vs.Spec.Resources.CPU.Type != nil {
		return fmt.Errorf("GPU type cannot be set if CPU type is set")
	}
	vs.Spec.Resources.GPU.Type = &gpuType
	return nil
}

// Set the VirtualServer GPU count
func (vs *VirtualServer) SetGPUCount(gpuCount uint32) error {
	if vs.Spec.Resources.CPU.Type != nil {
		return fmt.Errorf("GPU count cannot be set if CPU type is set")
	}
	vs.Spec.Resources.GPU.Count = &gpuCount
	return nil
}

// Set the VirtualServer Memory (RAM) size
func (vs *VirtualServer) SetMemory(memory string) error {
	mem, err := resource.ParseQuantity(memory)
	if err != nil {
		return fmt.Errorf("Cound not parse memory string")
	}
	vs.Spec.Resources.Memory = mem
	return nil
}

// Add a user to the VirtualServer
// The user will be used to configure the VirtualServer via cloudinit if supported
func (vs *VirtualServer) AddUser(user VirtualServerUser) {
	for i, u := range vs.Spec.Users {
		if u.Username == user.Username {
			vs.Spec.Users[i].Password = user.Password
			vs.Spec.Users[i].SSHPublicKey = user.SSHPublicKey
			return
		}
	}
	vs.Spec.Users = append(vs.Spec.Users, user)
}

//Add custom CloudInit attribute to VirtualServer
func (vs *VirtualServer) AddCloudInit(cloudInit string) error {
	// Current size of the VS' script is about 250 characeters
	var totalUsersSize int = 0
	for i := range vs.Spec.Users {
		totalUsersSize += vs.Spec.Users[i].GetSize()
	}
	MaxCustomCloudInitLength := corev1.MaxSecretSize - totalUsersSize

	if len(cloudInit) <= MaxCustomCloudInitLength {
		vs.Spec.CloudInit = cloudInit
		return nil
	}
	msg := fmt.Sprintf("Cloud-init script must be %d characters length or less", MaxCustomCloudInitLength)
	return errors.New(msg)
}

func (vs *VirtualServer) IsValidCloudInit() error {
	var cloudInit map[string]interface{}
	cloudInitStr := vs.Spec.CloudInit
	return yaml.Unmarshal([]byte(cloudInitStr), &cloudInit)
}

func (vs *VirtualServer) AddDNSConfig(dnsConfig *corev1.PodDNSConfig) {
	vs.Spec.Network.DNSConfig = dnsConfig
}

func (vs *VirtualServer) AddDNSPolicy(dnsPolicy *corev1.DNSPolicy) {
	vs.Spec.Network.DNSPolicy = dnsPolicy
}

func (vs *VirtualServer) SetMacAddress(macAddress string) error {
	matched, err := regexp.MatchString(MacAddressRegEx, macAddress)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid format of MAC address, it must be ff:ff:ff:ff:ff:ff or FF-FF-FF-FF-FF-FF")
	}
	vs.Spec.Network.MACAddress = macAddress
	return nil
}

func (vs *VirtualServer) SetFirmwareSerial(serial string) error {
	matched, err := regexp.MatchString(FirmwareSerialRegEx, serial)
	if err != nil {
		return err
	}
	if !matched {
		return fmt.Errorf("invalid format of firmware serial, it must be ffffffff-ffff-ffff-ffff-ffff-ffffffffffff")
	}
	vs.Spec.Firmware.Serial = serial
	return nil
}

func (vs *VirtualServer) SetFirmwareUUID(uuid types.UID) error {
	vs.Spec.Firmware.UUID = uuid
	return nil
}

// Set whether the VirtualServer will automatically start upon creation
func (vs *VirtualServer) InitializeRunning(initRunning bool) {
	vs.Spec.InitializeRunning = initRunning
}

func (vs *VirtualServer) SetHeadless(headless bool) {
	vs.Spec.Network.Headless = headless
}

func (vs *VirtualServer) RunStrategy(runStrategy kvv1.VirtualMachineRunStrategy) {
	vs.Spec.RunStrategy = &runStrategy
}

func (vs *VirtualServer) UseVirtioTransitional(useVirtioTransitional bool) {
	vs.Spec.UseVirtioTransitional = &useVirtioTransitional
}

func (vs *VirtualServer) TerminationGracePeriodSeconds(seconds int64) {
	vs.Spec.TerminationGracePeriodSeconds = &seconds
}

// Expose a TCP port on the VirtualServer
func (vs *VirtualServer) ExposeTCPPort(port int32) error {
	return vs.exposePort(port, corev1.ProtocolTCP)
}

// Expose TCP ports on the VirtualServer
func (vs *VirtualServer) ExposeTCPPorts(ports []int32) error {
	for _, port := range ports {
		err := vs.exposePort(port, corev1.ProtocolTCP)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vs *VirtualServer) EnablePublicIP(enable bool) {
	vs.Spec.Network.Public = enable
}

// Expose a UDP port on the VirtualServer
func (vs *VirtualServer) ExposeUDPPort(port int32) error {
	return vs.exposePort(port, corev1.ProtocolUDP)
}

// Expose UDP ports on the VirtualServer
func (vs *VirtualServer) ExposeUDPPorts(ports []int32) error {
	for _, port := range ports {
		err := vs.exposePort(port, corev1.ProtocolUDP)
		if err != nil {
			return err
		}
	}
	return nil
}

func (vs *VirtualServer) exposePort(port int32, protocol corev1.Protocol) error {
	if vs.Spec.Network.DirectAttachLoadBalancerIP != false {
		return fmt.Errorf("Ports cannot be exposed if DirectAttachLoadBalancerIP is enabled")
	}

	var ports *[]Port
	if protocol == corev1.ProtocolTCP {
		ports = &vs.Spec.Network.TCP.Ports
	} else if protocol == corev1.ProtocolUDP {
		ports = &vs.Spec.Network.UDP.Ports
	}

	if len(*ports) >= 10 {
		return fmt.Errorf("A maximum of 10 exposed ports are permitted")
	}
	for _, p := range *ports {
		if int32(p) == port {
			return nil
		}
	}
	*ports = append(*ports, Port(port))
	return nil
}

// Enable/disable DirectAttachLoadBalancerIP
func (vs *VirtualServer) DirectAttachLoadBalancerIP(attach bool) {
	vs.Spec.Network.DirectAttachLoadBalancerIP = attach
}

// Add a floating IP to the VirtualServer
// The loadbalancer IP will be extracted from an existing loadbalancer service with the provided name
func (vs *VirtualServer) AddFloatingIP(loadBalancerServiceName string) {
	for _, flIP := range vs.Spec.Network.FloatingIPs {
		if flIP.ServiceName == loadBalancerServiceName {
			return
		}
	}
	vs.Spec.Network.FloatingIPs = append(vs.Spec.Network.FloatingIPs, VirtualServerFloatingIP{
		ServiceName: loadBalancerServiceName,
	})
}

type VirtualServerStorageRootPVCSource struct {
	Size             string
	PVCName          string
	PVCNamespace     string
	StorageClassName string
	VolumeMode       corev1.PersistentVolumeMode
	AccessMode       corev1.PersistentVolumeAccessMode
}

// Configure the root storage with a PVC as the source
func (vs *VirtualServer) ConfigureStorageRootWithPVCSource(source VirtualServerStorageRootPVCSource) error {
	sourcePVC := cdiv1alpha.DataVolumeSourcePVC{
		Name:      source.PVCName,
		Namespace: source.PVCNamespace,
	}

	sz, err := resource.ParseQuantity(source.Size)
	if err != nil {
		return fmt.Errorf("Cound not parse size string")
	}

	vs.Spec.Storage.Root = VirtualServerStorageRoot{
		Size: sz,
		Source: cdiv1alpha.DataVolumeSource{
			PVC: &sourcePVC,
		},
		StorageClassName: source.StorageClassName,
		VolumeMode:       source.VolumeMode,
		AccessMode:       source.AccessMode,
	}
	return nil
}

type VirtualServerStorageRootHTTPSource struct {
	Size             string
	ImageUrl         string
	StorageClassName string
	VolumeMode       corev1.PersistentVolumeMode
	AccessMode       corev1.PersistentVolumeAccessMode
}

// Configure the root storage with a url to an image as the source
func (vs *VirtualServer) ConfigureStorageRootWithHTTPSource(source VirtualServerStorageRootHTTPSource) error {
	sourceHTTP := cdiv1alpha.DataVolumeSourceHTTP{
		URL: source.ImageUrl,
	}

	sz, err := resource.ParseQuantity(source.Size)
	if err != nil {
		return fmt.Errorf("Cound not parse size string")
	}

	vs.Spec.Storage.Root = VirtualServerStorageRoot{
		Size: sz,
		Source: cdiv1alpha.DataVolumeSource{
			HTTP: &sourceHTTP,
		},
		StorageClassName: source.StorageClassName,
		VolumeMode:       source.VolumeMode,
		AccessMode:       source.AccessMode,
	}
	return nil
}

// Add a PVC as a disk to the VirtualServer
func (vs *VirtualServer) AddPVCDisk(name string, pvcName string, readOnly bool) {
	pvcSource := corev1.PersistentVolumeClaimVolumeSource{
		ClaimName: pvcName,
		ReadOnly:  readOnly,
	}
	disk := VirtualServerStorageVolume{
		Name: name,
		Spec: kvv1.VolumeSource{
			PersistentVolumeClaim: &pvcSource,
		},
	}

	for _, d := range vs.Spec.Storage.AdditionalDisks {
		if d.Name == disk.Name {
			d.Spec = disk.Spec
			return
		}
	}

	vs.Spec.Storage.AdditionalDisks = append(vs.Spec.Storage.AdditionalDisks, disk)
}

// Add an ephemeral EmptyDisk as a disk to the VirtualServer
func (vs *VirtualServer) AddEmptyDiskDisk(name string, size string) error {
	sz, err := resource.ParseQuantity(size)
	if err != nil {
		return fmt.Errorf("Cound not parse size string")
	}
	emptyDisk := kvv1.EmptyDiskSource{
		Capacity: sz,
	}
	disk := VirtualServerStorageVolume{
		Name: name,
		Spec: kvv1.VolumeSource{
			EmptyDisk: &emptyDisk,
		},
	}

	for _, d := range vs.Spec.Storage.AdditionalDisks {
		if d.Name == disk.Name {
			d.Spec = disk.Spec
			return nil
		}
	}
	vs.Spec.Storage.AdditionalDisks = append(vs.Spec.Storage.AdditionalDisks, disk)
	return nil
}

// Add a filesystem to be mounted into the VirtualServer from a PVC
func (vs *VirtualServer) AddPVCFileSystem(name string, pvcName string, readOnly bool) {
	pvcSource := corev1.PersistentVolumeClaimVolumeSource{
		ClaimName: pvcName,
		ReadOnly:  readOnly,
	}
	fs := VirtualServerFilesystem{
		VirtualServerStorageVolume: VirtualServerStorageVolume{
			Name: name,
			Spec: kvv1.VolumeSource{
				PersistentVolumeClaim: &pvcSource,
			},
		},
	}

	for _, d := range vs.Spec.Storage.FileSystems {
		if d.Name == fs.Name {
			d.Spec = fs.Spec
			return
		}
	}

	vs.Spec.Storage.FileSystems = append(vs.Spec.Storage.FileSystems, fs)
}

func (vs *VirtualServer) AddSwap(size string) error {
	sz, err := resource.ParseQuantity(size)
	if err != nil {
		return fmt.Errorf("Cound not parse size string")
	}
	vs.Spec.Storage.Swap = &sz
	return nil
}

func (vs *VirtualServer) GetReadyStatus() *metav1.Condition {
	condition := apimeta.FindStatusCondition(vs.Status.Conditions, string(VSConditionTypeReady))
	if condition == nil {
		return nil
	}
	return condition.DeepCopy()
}

func (vs *VirtualServer) GetVMReadyStatus() *metav1.Condition {
	condition := apimeta.FindStatusCondition(vs.Status.Conditions, string(VSConditionTypeVMReady))
	if condition == nil {
		return nil
	}
	return condition.DeepCopy()
}

func (s *VirtualServerStatus) InternalIP() string {
	if s.Network.InternalIP != nil {
		return *s.Network.InternalIP
	}
	return ""
}

func (s *VirtualServerStatus) ExternalIP() string {
	if s.Network.ExternalIP != nil {
		return *s.Network.ExternalIP
	}
	return ""
}

func (s *VirtualServerStatus) FloatingIPs() map[string]string {
	return s.Network.FloatingIPs
}
