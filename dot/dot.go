package dot

import (
	"fmt"
	"strings"
	"text/template"

	"github.com/Ajnasz/i3tree/i3ipcnode"
	"github.com/Ajnasz/i3tree/tpl"
	i3ipc "github.com/ndemarinis/i3ipc-go"
)

func getName(node i3ipc.I3Node) string {
	return fmt.Sprintf("node%d", node.ID)

}

func getChildNodeNames(nodes []i3ipc.I3Node) string {
	var childNames []string
	for _, child := range nodes {
		childNames = append(childNames, getName(child))
	}

	return fmt.Sprintf("{%s}", strings.Join(childNames, ","))
}

func getNodeColor(node i3ipc.I3Node) string {
	switch node.Type {
	case "con":
		if len(node.Nodes) > 0 {
			return "orange"
		}
	case "floating_con":
		return "maroon"
	case "workspace":
		return "green"
	}

	return "black"
}

func getLabel(node i3ipc.I3Node, userTpl *template.Template) (string, error) {
	return tpl.String(userTpl, node)
}

func getDefName(node i3ipc.I3Node, userTpl *template.Template) (string, error) {
	label, err := getLabel(node, userTpl)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%s [label = %q, color = %q]", getName(node), label, getNodeColor(node)), nil
}

func printDot(node i3ipc.I3Node, userTpl *template.Template) error {
	defName, err := getDefName(node, userTpl)
	if err != nil {
		return err
	}
	fmt.Printf("%s;", defName)
	children := i3ipcnode.GetAllChildren(node)
	if len(children) > 0 {
		fmt.Printf("%s -> %s;", getName(node), getChildNodeNames(children))
		fmt.Println()

		for _, child := range children {
			printDot(child, userTpl)
		}
	}

	return nil
}

func Print(tree i3ipc.I3Node, userTpl *template.Template) {
	fmt.Println("digraph {")
	fmt.Println("node [shape=rect, style=rounded]")
	printDot(tree, userTpl)
	fmt.Println("}")
}
