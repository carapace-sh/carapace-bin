package action

import (
	"encoding/json"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
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
			return ActionCompleterGroups(nameVariant).Prefix(nameVariant + "@")

		case strings.Contains(c.Value, "/"):
			nameVariant, _, _ := strings.Cut(c.Value, "@")
			name, _, _ := strings.Cut(nameVariant, "/")
			return ActionCompleterVariants(name).Prefix(name + "/").NoSpace()

		default:
			return ActionCompleterNames().NoSpace()
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

func ActionCompleterNames() carapace.Action {
	return actionCompleters("", func(m completer.CompleterMap) carapace.Action {
		batch := carapace.Batch()
		for name, variants := range m {
			v := variants[0]
			batch = append(batch, carapace.ActionStyledValuesDescribed(name, v.Description, completerStyle(v)).Tag(completerTag(v)))
		}
		return batch.ToA()
	})
}

func ActionCompleterVariants(name string) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		if name == "" {
			return carapace.ActionValues()
		}

		return actionCompleters(name, func(m completer.CompleterMap) carapace.Action {
			// TODO slow
			batch := carapace.Batch()
			for moep, variants := range m {
				if name != "" && name != moep {
					continue
				}
				for _, v := range variants {
					if v.Variant == "" {
						v.Variant = "default" // pseudo-variant TODO add/document as keyword
					}
					switch name {
					case "":
						batch = append(batch, carapace.ActionValues(v.Variant).Tag(completerTag(v)))
					default:
						batch = append(batch, carapace.ActionStyledValuesDescribed(v.Variant, v.Description, completerStyle(v)).Tag(completerTag(v)))
					}
				}
			}
			return batch.ToA().Unique()
		})
	})
}

func ActionCompleterGroups(nameVariant string) carapace.Action {
	return actionCompleters(nameVariant, func(m completer.CompleterMap) carapace.Action {
		// TODO slow
		batch := carapace.Batch()
		for _, variants := range m {
			for _, v := range variants {
				batch = append(batch, carapace.ActionStyledValuesDescribed(v.Group, v.Description, completerStyle(v)).Tag(completerTag(v)))
			}
		}
		return batch.ToA().Unique()
	})
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
