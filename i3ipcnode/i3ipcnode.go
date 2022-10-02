// Package i3ipcnode contains utility method to manage i3ipc.Nodes
package i3ipcnode

import i3ipc "github.com/ndemarinis/i3ipc-go"

// GetAllChildren merges an i3ipc.I3Node Nodes and Floating_Nodes into a single
// slice
func GetAllChildren(node i3ipc.I3Node) []i3ipc.I3Node {
	size := len(node.Nodes) + len(node.Floating_Nodes)

	output := make([]i3ipc.I3Node, size)
	if size == 0 {
		return output
	}

	copy(output, node.Nodes)
	if len(node.Floating_Nodes) > 0 {
		copy(output[len(node.Nodes):], node.Floating_Nodes)
	}
	return output
}
