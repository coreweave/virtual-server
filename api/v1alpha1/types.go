package v1alpha1

// +kubebuilder:validation:Enum=cpu;gpu
type SystemPresetClass string

const (
	SystemPresetClassCPU SystemPresetClass = "cpu"
	SystemPresetClassGPU SystemPresetClass = "gpu"
)
