package jj

import (
	"strconv"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace/pkg/style"
	"github.com/pelletier/go-toml"
)

func ActionConfigs(includeDefaults bool) carapace.Action {
	if !includeDefaults {
		return actionConfigs(false)
	}
	return carapace.Batch(
		actionConfigs(true),
		actionConfigs(false).Style(style.Blue),
	).ToA()
}

func actionConfigs(includeDefaults bool) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		args := []string{"config", "list"}
		if includeDefaults {
			args = append(args, "--include-defaults")
		}
		return carapace.ActionExecCommand("jj", args...)(func(output []byte) carapace.Action {
			var config map[string]any
			if err := toml.Unmarshal(output, &config); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0)
			for key, value := range flatten(config) {
				switch {
				case strings.Contains(value, "\n"): // multi-line templates
					vals = append(vals, key, "")
				default:
					vals = append(vals, key, value)
				}
			}
			return carapace.ActionValuesDescribed(vals...)
		})
	})
}

func flatten(m map[string]any) map[string]string {
	flattened := make(map[string]string)
	for key, value := range m {
		switch value := value.(type) {
		case bool:
			flattened[key] = strconv.FormatBool(value)
		case string:
			flattened[key] = value
		case int:
			flattened[key] = strconv.Itoa(value)
		case map[string]any:
			for k, v := range flatten(value) {
				flattened[key+"."+k] = v
			}
		default:
			flattened[key] = "" // TODO support more types
		}
	}
	return flattened
}
