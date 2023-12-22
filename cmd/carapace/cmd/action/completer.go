package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

func ActionCompleters() carapace.Action {
	return carapace.ActionCallback(func(c carapace.Context) carapace.Action {
		return carapace.ActionExecCommand("carapace", "--list", "--format", "json")(func(output []byte) carapace.Action {
			var completers []struct {
				Name        string
				Description string
				Spec        string
				Overlay     string
			}
			if err := json.Unmarshal(output, &completers); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			vals := make([]string, 0, len(completers))
			for _, completer := range completers {
				s := style.Default
				if completer.Spec != "" {
					s = style.Blue
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
