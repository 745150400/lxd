package cluster

//go:generate -command mapper lxd-generate db mapper -t nodes.mapper.go
//go:generate mapper reset -i -b "//go:build linux && cgo && !agent"
//
//go:generate mapper stmt -e node id
//
//go:generate mapper method -i -e node ID

// Node represents a LXD cluster node.
type Node struct {
	ID   int
	Name string
}

// NodeFilter specifies potential query parameter fields.
type NodeFilter struct {
	Name *string
}
