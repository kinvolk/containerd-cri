package specs

// State holds information about the runtime state of the container.
type State struct {
	// Version is the version of the specification that is supported.
	Version string `json:"ociVersion"`
	// ID is the container ID
	ID string `json:"id"`
	// Status is the runtime status of the container.
	Status string `json:"status"`
	// Pid is the process ID for the container process.
	Pid int `json:"pid,omitempty"`
	// Bundle is the path to the container's bundle directory.
	Bundle string `json:"bundle"`
	// Annotations are key values associated with the container.
	Annotations map[string]string `json:"annotations,omitempty"`
}

// FdIndexKey is the key used in the FdIndexes map of the ContainerProcessState struct.
type FdIndexKey string

const (
	// SeccompFdIndexKey is the index of the seccomp notify file descriptor.
	SeccompFdIndexKey FdIndexKey = "seccompFd"
	// PidFdIndexKey is the index of the target process file descriptor.
	PidFdIndexKey FdIndexKey = "pidFd"
)

// ContainerProcessState holds information about the state of a container process.
type ContainerProcessState struct {
	// Version is the version of the specification that is supported.
	Version string `json:"ociVersion"`
	// FdIndexes is a map containing the indexes of the file descriptors in the `SCM_RIGHTS` array.
	FdIndexes map[FdIndexKey]int `json:"fdIndexes"`
	// Pid is the process ID as seen by the runtime.
	Pid int `json:"pid"`
	// Opaque metadata.
	Metadata string `json:"metadata,omitempty"`
	// State of the container.
	State State `json:"state"`
}
