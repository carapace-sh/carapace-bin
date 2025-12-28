package action

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/carapace-sh/carapace-bridge/pkg/actions/bridge"
	"github.com/carapace-sh/carapace/pkg/style"
)

func ActionCompleters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if strings.Count(c.Value, "@") > 1 || strings.Count(c.Value, "/") > 1 {
			return carapace.ActionValues()
		}

		switch {
		case strings.Contains(c.Value, "@"):
			nameVariant, _, _ := strings.Cut(c.Value, "@")
			return ActionGroups(nameVariant).Prefix(nameVariant + "@")

		case strings.Contains(c.Value, "/"):
			nameVariant, _, _ := strings.Cut(c.Value, "@")
			name, _, _ := strings.Cut(nameVariant, "/")
			return ActionVariants(name).Prefix(name + "/").NoSpace()

		default:
			return ActionNames().NoSpace()
		}
	})
}

func actionCompleters(filter string, f func(m completer.CompleterMap) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("carapace", "--list", filter)(func(output []byte) carapace.Action {
			var m completer.CompleterMap
			if err := json.Unmarshal(output, &m); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return f(m)
		})
	})
}

func ActionNames() carapace.Action {
	return actionCompleters("", func(m completer.CompleterMap) carapace.Action {
		batch := carapace.Batch()
		for name, variants := range m {
			v := variants[0]
			batch = append(batch, carapace.ActionStyledValuesDescribed(name, v.Description, completerStyle(v)).Tag(completerTag(v)))
		}
		return batch.ToA()
	})
}

func ActionVariants(name string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if name == "" {
			return carapace.ActionValues()
		}

		return actionCompleters(name, func(m completer.CompleterMap) carapace.Action {
			// TODO slow
			batch := carapace.Batch()
			for _, variants := range m {
				for _, v := range variants {
					if v.Variant == "" {
						continue
					}
					switch name {
					case "":
						batch = append(batch, carapace.ActionValues(v.Variant).Tag(completerTag(v)))
					default:
						batch = append(batch, carapace.ActionStyledValuesDescribed(v.Variant, v.Description, completerStyle(v)).Tag(completerTag(v)))
					}
				}
			}
			return batch.ToA()
		})
	})
}

func ActionGroups(nameVariant string) carapace.Action {
	return actionCompleters(nameVariant, func(m completer.CompleterMap) carapace.Action {
		// TODO slow
		batch := carapace.Batch()

		_, variant, _ := strings.Cut(nameVariant, "/")
		if _, ok := bridge.Get(variant); ok {
			batch = append(batch, carapace.ActionValues("bridge")) // potentially unknown bridge (e.g. `cobra`)
		}

		descriptions := map[string]string{
			"android": "termux completers",
			"bridge":  "bridged completions",
			"bsd":     "bsd-like completers",
			"common":  "common completers",
			"darwin":  "macos completers",
			"linux":   "linux completers",
			"system":  "system specs",
			"unix":    "unix-like completers",
			"user":    "user specs",
			"windows": "windows completers",
		}

		for _, variants := range m {
			for _, v := range variants {
				// TODO skip when group already exists
				batch = append(batch, carapace.ActionValuesDescribed(v.Group, descriptions[v.Group]))
			}
		}
		return batch.ToA().Unique()
	}).Tag("groups")
}

func completerStyle(c completer.Completer) string {
	switch {
	case c.Spec != "":
		return style.Blue
	case c.Group == "bridge": // TODO detect bridges correctly (config/env)
		return style.Dim
	default:
		return style.Default
	}
}

func completerTag(c completer.Completer) string {
	switch {
	case c.Spec != "":
		return "user completers"
	case c.Group == "bridge": // TODO detect bridges correctly (config/env)
		return "bridged completers"
	default:
		return "internal completers"
	}
}
