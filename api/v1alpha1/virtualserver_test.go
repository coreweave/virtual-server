package v1alpha1_test

import (
	"context"
	"fmt"
	"os"

	vsv1alpha "github.com/coreweave/virtual-server/api/v1alpha1"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/client/config"
)

// The example describes the creation of a VirtualServer on the Coreweave Cloud kubernetes platform
func Example_create() {
	// Create the a new kube client
	// Uses the value of the KUBECONFIG environment variable as a filepath to a kube config file
	c, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		fmt.Println("Failed to create client")
		os.Exit(1)
	}
	vsv1alpha.AddToScheme(c.Scheme())

	// Create a new VirtualServer with the name "my-virtual-server" to be deployed in the "default" namespace
	virtualServer := vsv1alpha.NewVirtualServer("my-virtual-server", "default")
	// Set the region the VirtualServer will be deployed to
	virtualServer.SetRegion("ord1")
	// Specify the VirtualServer operating system
	virtualServer.SetOS(vsv1alpha.VirtualServerOSTypeLinux)
	// Set a GPU type request for the VirtualServer
	virtualServer.SetGPUType("Quadro_RTX_4000")
	// Set the number of GPUs to request for the VirtualServer
	virtualServer.SetGPUCount(1)
	// Set the cpu core count for the VirtualServer
	virtualServer.SetCPUCount(2)
	// Set the memory request for the VirtualServer
	virtualServer.SetMemory("16Gi")
	// Add a user to be added by cloudinit
	virtualServer.AddUser(vsv1alpha.VirtualServerUser{
		Username: "myuser",
		Password: "password",
	})
	// Configure the root filesystem of the VirtualServer to clone a preexisting PVC named ubuntu1804-docker-master-20210210-ewr1
	// sourced in the vd-images namespace
	err = virtualServer.ConfigureStorageRootWithPVCSource(vsv1alpha.VirtualServerStorageRootPVCSource{
		Size:             "40Gi",
		PVCName:          "ubuntu1804-docker-master-20210210-ewr1",
		PVCNamespace:     "vd-images",
		StorageClassName: "block-nvme-ewr1",
		VolumeMode:       corev1.PersistentVolumeBlock,
		AccessMode:       corev1.ReadWriteOnce,
	})
	if err != nil {
		fmt.Println("Cound not configure root filesystem")
	}
	// Expose tcp ports 22 and 443 on the VirtualServer
	virtualServer.ExposeTCPPorts([]int32{22, 443})
	// Expose a single udp port 4172 on the VirtualServer
	virtualServer.ExposeUDPPort(4172)
	// Set the VirtualServer to start as soon as it is created
	virtualServer.InitializeRunning(true)

	// Create the VirtualServer via the kube client
	err = c.Create(context.Background(), virtualServer)
	if err != nil {
		fmt.Printf("Failed to create VirtualServer\nReason: %s", err.Error())
	}
}

// The example describes how to get a VirtualServer running on the Coreweave Cloud kubernetes platform
func Example_get() {
	// Create the a new kube client
	// Uses the value of the KUBECONFIG environment variable as a filepath to a kube config file
	c, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		fmt.Println("Failed to create client")
		os.Exit(1)
	}
	vsv1alpha.AddToScheme(c.Scheme())

	virtualServer := &vsv1alpha.VirtualServer{}
	// Get the VirtualServer named "my-virtual-server" in the "default" namespace
	err = c.Get(context.Background(), client.ObjectKey{
		Namespace: "default",
		Name:      "my-virtual-server",
	}, virtualServer)

	if err != nil {
		fmt.Println("Failed to get VirtualServer")
	}

	fmt.Printf("Name: %s\nNamespace: %s\n", virtualServer.Name, virtualServer.Namespace)
	// output:
	// Name: my-virtual-server
	// Namespace: default
}

// The example describes how to get the ready status of a VirtualServer running on the Coreweave Cloud kubernetes platform
func Example_getStatus() {
	// Create the a new kube client
	// Uses the value of the KUBECONFIG environment variable as a filepath to a kube config file
	c, err := client.New(config.GetConfigOrDie(), client.Options{})
	if err != nil {
		fmt.Println("Failed to create client")
		os.Exit(1)
	}
	vsv1alpha.AddToScheme(c.Scheme())

	virtualServer := &vsv1alpha.VirtualServer{}
	// Get the VirtualServer named "my-virtual-server" in the "default" namespace
	err = c.Get(context.Background(), client.ObjectKey{
		Namespace: "default",
		Name:      "my-virtual-server",
	}, virtualServer)

	if err != nil {
		fmt.Println("Failed to get VirtualServer")
	}

	// Get the ready status of the VirtualServer
	status := virtualServer.GetReadyStatus()
	if status != nil {
		fmt.Printf("Status: %s", status.Status)
	}
	// output:
	// Status: "True" or "False"
}
