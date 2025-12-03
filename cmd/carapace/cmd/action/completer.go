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
			name, variant, _ := strings.Cut(nameVariant, "/")
			return actionCompleterGroups(name, variant).Prefix(nameVariant + "@")

		case strings.Contains(c.Value, "/"):
			nameVariant, _, _ := strings.Cut(c.Value, "@")
			name, _, _ := strings.Cut(nameVariant, "/")
			return actionCompleterVariants(name).Prefix(name + "/")

		default:
			return actionCompleterNames()
		}
	})
}

func actionCompleters(f func(m completer.CompleterMap) carapace.Action) carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("carapace", "--list")(func(output []byte) carapace.Action {
			var m completer.CompleterMap
			if err := json.Unmarshal(output, &m); err != nil {
				return carapace.ActionMessage(err.Error())
			}
			return f(m)
		})
	})
}

func actionCompleterNames() carapace.Action {
	return actionCompleters(func(m completer.CompleterMap) carapace.Action {
		batch := carapace.Batch()
		for name, variants := range m {
			v := variants[0]
			batch = append(batch, carapace.ActionStyledValuesDescribed(name, v.Description, completerStyle(v)).Tag(completerTag(v)))
		}
		return batch.ToA()
	})
}

func actionCompleterVariants(filterName string) carapace.Action {
	return actionCompleters(func(m completer.CompleterMap) carapace.Action {
		// TODO slow
		batch := carapace.Batch()
		for name, variants := range m {
			if filterName != "" && filterName != name {
				continue
			}
			for _, v := range variants {
				if v.Variant != "" {
					switch filterName {
					case "":
						batch = append(batch, carapace.ActionValues(v.Variant).Tag(completerTag(v)))
					default:
						batch = append(batch, carapace.ActionStyledValuesDescribed(v.Variant, v.Description, completerStyle(v)).Tag(completerTag(v)))
					}
				}
			}
		}
		return batch.ToA().Unique()
	})
}

func actionCompleterGroups(filterName, filterVariant string) carapace.Action {
	return actionCompleters(func(m completer.CompleterMap) carapace.Action {
		// TODO slow
		batch := carapace.Batch()
		for name, variants := range m {
			if filterName != "" && filterName != name {
				continue
			}
			for _, v := range variants {
				if filterVariant == "" || filterVariant == v.Variant {
					switch {
					case filterName == "" || filterVariant == "":
						batch = append(batch, carapace.ActionValues(v.Group).Tag(completerTag(v)))
					default:
						batch = append(batch, carapace.ActionStyledValuesDescribed(v.Group, v.Description, completerStyle(v)).Tag(completerTag(v)))
					}
				}
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
