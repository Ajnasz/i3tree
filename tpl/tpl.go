// Package tpl implements utility function to render a go template into a string
package tpl

import (
	"bytes"
	"text/template"

	i3ipc "github.com/ndemarinis/i3ipc-go"
)

// String renders go template to a string
func String(tpl *template.Template, node i3ipc.I3Node) (string, error) {
	var buf bytes.Buffer
	if err := tpl.Execute(&buf, node); err != nil {
		return "", err
	}

	return buf.String(), nil
}
