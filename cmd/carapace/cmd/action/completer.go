package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func ActionCompleters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("carapace", "--list", "--all", "--format", "json")(func(output []byte) carapace.Action {
			var completers map[string]struct {
				Name        string
				Description string
				Spec        string
				Overlay     string
				Bridge      string
			}
			if err := json.Unmarshal(output, &completers); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0, len(completers))
			for _, completer := range completers {
				s := style.Default
				if completer.Spec != "" {
					s = style.Blue
				} else if completer.Bridge != "" {
					s = style.Dim
				}
				if completer.Overlay != "" {
					s = style.Of(s, style.Underlined)
				}

				vals = append(vals, completer.Name, completer.Description, s)
			}
			return carapace.ActionStyledValuesDescribed(vals...)
		})
	})
}
