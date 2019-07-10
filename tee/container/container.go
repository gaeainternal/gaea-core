package container

// Container to do trusted execution
type Container interface {
	// Create a container
	Create() error

	// Upload algorithm and data to container
	Upload(algorithm []byte, dataList [][]byte) error

	// Verify algorithm and data integrity
	Verify(algorithmHash string, dataHash []string) error

	// Execute the container
	Execute() ([]byte, error)

	// Destroy the container
	Destroy() error
}

// Type for how to use the trusted execution environment
type Type int

// Docker is a folder for using container inside a chaincode container.
//
// Sibling is a sibling docker container for using container inside a chaincode container.
// It needs to update core.yaml and dockercontroller.go to bind docker.sock to mounts.
//
// SGX is a Hardware chip CPU.
const (
	Docker Type = iota
	Sibling
	SGX
)
