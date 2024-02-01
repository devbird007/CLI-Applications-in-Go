// Package scan provides types and functions to perform TCP port
// scans on a list of hosts
package scan

// PortState represents the state of a single TCP port
type PortState struct {
	Port int
	Open state
}
