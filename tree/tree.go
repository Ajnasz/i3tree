package tree

import (
	"fmt"

	"github.com/Ajnasz/i3tree/i3ipcnode"
	i3ipc "github.com/ndemarinis/i3ipc-go"
)

func printTree(node i3ipc.I3Node, prefix string) {
	fmt.Printf("%s [%s %s] %s #%s#", prefix, node.Type, node.Layout, node.Name, node.Floating)
	fmt.Println()

	children := i3ipcnode.GetAllChildren(node)
	if len(children) > 0 {
		for _, child := range children {
			printTree(child, "|  "+prefix)
		}
	}
}

func Print(node i3ipc.I3Node) {
	printTree(node, "|")
}
