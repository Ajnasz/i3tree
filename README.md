# I3Tree

Generates the i3 window manager node tree in a text or graphwiz format.

## Parameters

### -dot

Generate graphwiz format:

```
./i3tree -dot | dot -Tsvg -o a.svg && xdg-open a.svg
```

### -template

Default template to render the node name:

`[{{.Type}} {{.Layout}}] {{.Name}} {{.Floating}}`

Override the default template:

```
./i3tree -template '{{.Name}}'
```
