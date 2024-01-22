package action

import (
	"encoding/json"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace/pkg/style"
)

type CompleterOpts struct {
	Internal bool
	Spec     bool
	Bridge   bool
}

func (o CompleterOpts) Default() CompleterOpts {
	return CompleterOpts{
		Internal: true,
		Spec:     true,
		Bridge:   true,
	}
}

func ActionCompleters(opts CompleterOpts) carapace.Action {
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

			batch := carapace.Batch() // TODO value map by tag should be better here
			for _, completer := range completers {
				var s, t string
				switch {
				case completer.Spec != "":
					if !opts.Spec {
						continue
					}
					s = style.Blue
					t = "user completers"
				case completer.Bridge != "":
					if !opts.Bridge {
						continue
					}
					s = style.Dim
					t = "bridged completers"
				default:
					if !opts.Internal {
						continue
					}
					s = style.Default
					t = "internal completers"
				}

				if completer.Overlay != "" {
					s = style.Of(s, style.Underlined)
				}

				batch = append(batch, carapace.ActionStyledValuesDescribed(completer.Name, completer.Description, s).Tag(t))
			}
			return batch.ToA()
		})
	})
}
