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

`[{{.Type}} {{.Layout}}] {{.Name}}`

Override the default template:

```
./i3tree -template '{{.Name}}'
```

## Example output

### text

```
| [root splith] root
|  | [output output] __i3
|  |  | [con splith] content
|  |  |  | [workspace splith] __i3_scratch
|  | [output output] HDMI-1
|  |  | [dockarea dockarea] topdock
|  |  | [con splith] content
|  |  |  | [workspace splith] 01
|  |  |  |  | [con splitv]
|  |  |  |  |  | [con splith]
|  |  |  |  |  |  | [con splith]
|  |  |  |  |  |  |  | [con tabbed]
|  |  |  |  |  |  |  |  | [con splith] work:12.0 i3tree zsh
|  |  |  |  |  |  |  |  | [con splith] Editing i3tree/README.md at master · Ajnasz/i3tree — Firefox Developer Edition
|  |  |  | [workspace splith] 02 >_ mutt/irssi
|  |  |  |  | [con splith] comm:0.0 mutt
|  |  |  | [workspace splith] 03   Telegram
|  |  |  |  | [con splith] Telegram (9)
|  |  |  | [workspace splith] 04   Spotify
|  |  |  |  | [con splith] Spotify
|  |  |  | [workspace splith] 10 KeePassXC
|  |  |  |  | [con splith] KeePassXC
|  |  | [dockarea dockarea] bottomdock
|  |  |  | [con splith] i3bar for output HDMI-1

```

### SVG

![a](https://user-images.githubusercontent.com/38329/193446544-feb7082d-50a5-479f-b428-c1264f86637b.svg)
