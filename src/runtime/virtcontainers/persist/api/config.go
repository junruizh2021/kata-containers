// Copyright (c) 2016 Intel Corporation
// Copyright (c) 2019 Huawei Corporation
//
// SPDX-License-Identifier: Apache-2.0
//

package persistapi

import (
	"github.com/kata-containers/kata-containers/src/runtime/pkg/device/config"
	"github.com/opencontainers/runc/libcontainer/configs"
	specs "github.com/opencontainers/runtime-spec/specs-go"
)

// HypervisorConfig saves configurations of sandbox hypervisor
type HypervisorConfig struct {
	// KernelPath is the guest kernel host path.
	KernelPath string

	// ImagePath is the guest image host path.
	ImagePath string

	// InitrdPath is the guest initrd image host path.
	// ImagePath and InitrdPath cannot be set at the same time.
	InitrdPath string

	// FirmwarePath is the bios host path
	FirmwarePath string

	// MachineAccelerators are machine specific accelerators
	MachineAccelerators string

	// CPUFeatures are cpu specific features
	CPUFeatures string

	// HypervisorPath is the hypervisor executable host path.
	HypervisorPath string

	// HypervisorCtlPath is the hypervisor ctl executable host path.
	HypervisorCtlPath string

	// HypervisorCtlPath is the hypervisor ctl executable host path.
	// JailerPath is the jailer executable host path.
	JailerPath string

	// BlockDeviceDriver specifies the driver to be used for block device
	// either VirtioSCSI or VirtioBlock with the default driver being defaultBlockDriver
	BlockDeviceDriver string

	// HypervisorMachineType specifies the type of machine being
	// emulated.
	HypervisorMachineType string

	// MemoryPath is the memory file path of VM memory. Used when either BootToBeTemplate or
	// BootFromTemplate is true.
	MemoryPath string

	// DevicesStatePath is the VM device state file path. Used when either BootToBeTemplate or
	// BootFromTemplate is true.
	DevicesStatePath string

	// EntropySource is the path to a host source of
	// entropy (/dev/random, /dev/urandom or real hardware RNG device)
	EntropySource string

	// Shared file system type:
	//   - virtio-9p (default)
	//   - virtio-fs
	SharedFS string

	// VirtioFSDaemon is the virtio-fs vhost-user daemon path
	VirtioFSDaemon string

	// VirtioFSCache cache mode for fs version cache
	VirtioFSCache string

	// File based memory backend root directory
	FileBackedMemRootDir string

	// VhostUserStorePath is the directory path where vhost-user devices
	// related folders, sockets and device nodes should be.
	VhostUserStorePath string

	// SeccompSandbox is the qemu function which enables the seccomp feature
	SeccompSandbox string

	// GuestHookPath is the path within the VM that will be used for 'drop-in' hooks
	GuestHookPath string

	// VMid is the id of the VM that create the hypervisor if the VM is created by the factory.
	// VMid is "" if the hypervisor is not created by the factory.
	VMid string

	// HypervisorPathList is the list of hypervisor paths names allowed in annotations
	HypervisorPathList []string

	// HypervisorCtlPathList is the list of hypervisor control paths names allowed in annotations
	HypervisorCtlPathList []string

	// JailerPathList is the list of jailer paths names allowed in annotations
	JailerPathList []string

	// EntropySourceList is the list of valid entropy sources
	EntropySourceList []string

	// VirtioFSDaemonList is the list of valid virtiofs names for annotations
	VirtioFSDaemonList []string

	// VirtioFSExtraArgs passes options to virtiofsd daemon
	VirtioFSExtraArgs []string

	// FileBackedMemRootList is the list of valid root directories values for annotations
	FileBackedMemRootList []string

	// VhostUserStorePathList is the list of valid values for vhost-user paths
	VhostUserStorePathList []string

	// Enable annotations by name
	EnableAnnotations []string

	// MemOffset specifies memory space for nvdimm device
	MemOffset uint64

	// RxRateLimiterMaxRate is used to control network I/O inbound bandwidth on VM level.
	RxRateLimiterMaxRate uint64

	// TxRateLimiterMaxRate is used to control network I/O outbound bandwidth on VM level.
	TxRateLimiterMaxRate uint64

	// SGXEPCSize specifies the size in bytes for the EPC Section.
	// Enable SGX. Hardware-based isolation and memory encryption.
	SGXEPCSize int64

	// NumVCPUs specifies default number of vCPUs for the VM.
	NumVCPUs uint32

	//DefaultMaxVCPUs specifies the maximum number of vCPUs for the VM.
	DefaultMaxVCPUs uint32

	// DefaultMem specifies default memory size in MiB for the VM.
	MemorySize uint32

	// DefaultBridges specifies default number of bridges for the VM.
	// Bridges can be used to hot plug devices
	DefaultBridges uint32

	// Msize9p is used as the msize for 9p shares
	Msize9p uint32

	// MemSlots specifies default memory slots the VM.
	MemSlots uint32

	// VirtioFSCacheSize is the DAX cache size in MiB
	VirtioFSCacheSize uint32

	// BlockDeviceCacheSet specifies cache-related options will be set to block devices or not.
	BlockDeviceCacheSet bool

	// BlockDeviceCacheDirect specifies cache-related options for block devices.
	// Denotes whether use of O_DIRECT (bypass the host page cache) is enabled.
	BlockDeviceCacheDirect bool

	// BlockDeviceCacheNoflush specifies cache-related options for block devices.
	// Denotes whether flush requests for the device are ignored.
	BlockDeviceCacheNoflush bool

	// DisableBlockDeviceUse disallows a block device from being used.
	DisableBlockDeviceUse bool

	// EnableIOThreads enables IO to be processed in a separate thread.
	// Supported currently for virtio-scsi driver.
	EnableIOThreads bool

	// Debug changes the default hypervisor and kernel parameters to
	// enable debug output where available.
	Debug bool

	// MemPrealloc specifies if the memory should be pre-allocated
	MemPrealloc bool

	// HugePages specifies if the memory should be pre-allocated from huge pages
	HugePages bool

	// VirtioMem is used to enable/disable virtio-mem
	VirtioMem bool

	// DisableNestingChecks is used to override customizations performed
	// when running on top of another VMM.
	DisableNestingChecks bool

	// DisableImageNvdimm disables nvdimm for guest rootfs image
	DisableImageNvdimm bool

	// HotPlugVFIO is used to indicate if devices need to be hotplugged on the
	// root, switch, bridge or no-port
	HotPlugVFIO config.PCIePort

	// ColdPlugVFIO is used to indicate if devices need to be coldplugged on the
	// root, bridge, switch or no-port
	ColdPlugVFIO config.PCIePort

	// BootToBeTemplate used to indicate if the VM is created to be a template VM
	BootToBeTemplate bool

	// BootFromTemplate used to indicate if the VM should be created from a template VM
	BootFromTemplate bool

	// DisableVhostNet is used to indicate if host supports vhost_net
	DisableVhostNet bool

	// EnableVhostUserStore is used to indicate if host supports vhost-user-blk/scsi
	EnableVhostUserStore bool
}

// KataAgentConfig is a structure storing information needed
// to reach the Kata Containers agent.
type KataAgentConfig struct {
	LongLiveConn bool
}

// ShimConfig is the structure providing specific configuration
// for shim implementation.
type ShimConfig struct {
	Path  string
	Debug bool
}

// NetworkConfig is the network configuration related to a network.
type NetworkConfig struct {
	NetworkID         string
	NetworkCreated    bool
	DisableNewNetwork bool
	InterworkingModel int
}

type ContainerConfig struct {
	Annotations map[string]string
	// Resources for recoding update
	Resources specs.LinuxResources
	ID        string
	RootFs    string
}

// SandboxConfig is a sandbox configuration.
// Refs: virtcontainers/sandbox.go:SandboxConfig
type SandboxConfig struct {
	// Cgroups specifies specific cgroup settings for the various subsystems that the container is
	// placed into to limit the resources the container has available
	Cgroups *configs.Cgroup `json:"cgroups"`

	// only one agent config can be non-nil according to agent type
	KataAgentConfig *KataAgentConfig `json:",omitempty"`

	KataShimConfig *ShimConfig

	// Custom SELinux security policy to the container process inside the VM
	GuestSeLinuxLabel string

	HypervisorType string

	// SandboxBindMounts - list of paths to mount into guest
	SandboxBindMounts []string

	// Experimental enables experimental features
	Experimental []string

	// Information for fields not saved:
	// * Annotation: this is kind of casual data, we don't need casual data in persist file,
	// if you know this data needs to persist, please gives it a specific field
	ContainerConfigs []ContainerConfig

	NetworkConfig NetworkConfig

	HypervisorConfig HypervisorConfig

	ShmSize uint64

	// SharePidNs sets all containers to share the same sandbox level pid namespace.
	SharePidNs bool

	// Stateful keeps sandbox resources in memory across APIs. Users will be responsible
	// for calling Release() to release the memory resources.
	Stateful bool

	// SystemdCgroup enables systemd cgroup support
	SystemdCgroup bool

	// SandboxCgroupOnly enables cgroup only at podlevel in the host
	SandboxCgroupOnly bool

	DisableGuestSeccomp bool

	// EnableVCPUsPinning controls whether each vCPU thread should be scheduled to a fixed CPU
	EnableVCPUsPinning bool
}
