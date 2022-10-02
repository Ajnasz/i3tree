package main

import (
	"flag"
	"fmt"
	"os"
	"text/template"

	"github.com/Ajnasz/i3tree/dot"
	"github.com/Ajnasz/i3tree/tree"
	i3ipc "github.com/ndemarinis/i3ipc-go"
)

func main() {
	var printDot bool
	var tplStr string
	flag.BoolVar(&printDot, "dot", false, "")
	flag.StringVar(&tplStr, "template", "[{{.Type}} {{.Layout}}] {{.Name}}", "template to render the node name")

	flag.Parse()
	ipcsocket, err := i3ipc.GetIPCSocket()

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	tpl, err := template.New("node").Parse(tplStr)

	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	i3tree, err := ipcsocket.GetTree()
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		os.Exit(1)
	}

	if printDot {
		dot.Print(i3tree, tpl)
	} else {
		tree.Print(i3tree, tpl)
	}
}
