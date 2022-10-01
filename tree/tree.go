package tree

import (
	"fmt"

	i3ipc "github.com/ndemarinis/i3ipc-go"
)

func printTree(node i3ipc.I3Node, prefix string) {
	fmt.Printf("%s [%s %s] %s", prefix, node.Type, node.Layout, node.Name)
	fmt.Println()

	if len(node.Nodes) > 0 {
		for _, child := range node.Nodes {
			printTree(child, "|  "+prefix)
		}
	}
}

func Print(node i3ipc.I3Node) {
	printTree(node, "|")
}
