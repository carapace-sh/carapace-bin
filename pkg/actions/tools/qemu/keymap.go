package qemu

import (
	"fmt"
	"strings"

	"github.com/carapace-sh/carapace"
)

// ActionKeymapLayouts completes XKB keyboard layouts
//
//	us (English (US))
//	de (German)
func ActionKeymapLayouts() carapace.Action {
	return carapace.ActionExecCommand("awk", "/^! layout/{found=1; next} /^!/{found=0} found && NF{sub(/^[ \\t]+/, \"\"); print}", "/usr/share/X11/xkb/rules/base.lst")(func(output []byte) carapace.Action {
		return parseXkbEntries(string(output)).Tag("layouts").Uid("qemu", "keymap-layouts")
	})
}

// ActionKeymapModels completes XKB keyboard models
//
//	pc105 (Generic 105-key PC)
//	pc104 (Generic 104-key PC)
func ActionKeymapModels() carapace.Action {
	return carapace.ActionExecCommand("awk", "/^! model/{found=1; next} /^!/{found=0} found && NF{sub(/^[ \\t]+/, \"\"); print}", "/usr/share/X11/xkb/rules/base.lst")(func(output []byte) carapace.Action {
		return parseXkbEntries(string(output)).Tag("models").Uid("qemu", "keymap-models")
	})
}

// ActionKeymapVariants completes XKB keyboard variants for a given layout
//
//	dvorak (Dvorak)
//	colemak (Colemak)
func ActionKeymapVariants() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		layout := c.Value
		if layout == "" {
			layout = "us"
		}
		return carapace.ActionExecCommand("awk", fmt.Sprintf("/^! variant/{found=1; next} /^!/{found=0} found && NF && $2 ~ /^%s:/{sub(/^[ \\t]+/, \"\"); gsub(/ %s:/, \" \", $0); print}", layout, layout), "/usr/share/X11/xkb/rules/base.lst")(func(output []byte) carapace.Action {
			return parseXkbEntries(string(output)).Tag("variants").Uid("qemu", "keymap-variants")
		})
	})
}

// ActionKeymapOptions completes XKB keyboard options
//
//	grp:alt_shift_toggle (Alt+Shift toggle)
//	caps:escape (Caps Lock acts as Escape)
func ActionKeymapOptions() carapace.Action {
	return carapace.ActionExecCommand("awk", "/^! option/{found=1; next} /^!/{found=0} found && NF{sub(/^[ \\t]+/, \"\"); print}", "/usr/share/X11/xkb/rules/base.lst")(func(output []byte) carapace.Action {
		return parseXkbEntries(string(output)).Tag("options").Uid("qemu", "keymap-options")
	})
}

func parseXkbEntries(output string) carapace.Action {
	lines := strings.Split(output, "\n")
	vals := make([]string, 0, len(lines)*2)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}
		parts := strings.SplitN(line, " ", 2)
		vals = append(vals, parts[0])
		if len(parts) > 1 {
			vals = append(vals, strings.TrimSpace(parts[1]))
		} else {
			vals = append(vals, parts[0])
		}
	}
	return carapace.ActionValuesDescribed(vals...)
}
