package tree

import (
	"fmt"
	"text/template"

	"github.com/Ajnasz/i3tree/i3ipcnode"
	"github.com/Ajnasz/i3tree/tpl"
	i3ipc "github.com/ndemarinis/i3ipc-go"
)

func printTree(node i3ipc.I3Node, userTpl *template.Template, prefix string) error {
	str, err := tpl.String(userTpl, node)
	if err != nil {
		return err
	}
	fmt.Printf("%s %s", prefix, str)
	fmt.Println()

	children := i3ipcnode.GetAllChildren(node)
	if len(children) > 0 {
		for _, child := range children {
			printTree(child, userTpl, "|  "+prefix)
		}
	}

	return nil
}

func Print(node i3ipc.I3Node, tpl *template.Template) {
	printTree(node, tpl, "|")
}
