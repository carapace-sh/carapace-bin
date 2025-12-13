package action

import (
	"encoding/json"
	"fmt"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/carapace-sh/carapace/pkg/style"
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
			// TODO align with new completer format
			// TODO overlays/bridge
			var completerMap completer.CompleterMap
			if err := json.Unmarshal(output, &completerMap); err != nil {
				return carapace.ActionMessage(err.Error())
			}

			batch := carapace.Batch() // TODO value map by tag should be better here
			for _, variants := range completerMap {
				for _, completer := range variants {
					var s, t string
					switch {
					case completer.Spec != "":
						if !opts.Spec {
							continue
						}
						s = style.Blue
						t = "user completers"
					// TODO re-enable bridge
					// case completer.Bridge != "":
					// 	if !opts.Bridge {
					// 		continue
					// 	}
					// 	s = style.Dim
					// 	t = "bridged completers"
					default:
						if !opts.Internal {
							continue
						}
						s = style.Default
						t = "internal completers"
					}

					// TODO re-enable overlay
					// if completer.Overlay != "" {
					// 	s = style.Of(s, style.Underlined)
					// }

					name := completer.Name
					if completer.Variant != "" {
						name = fmt.Sprintf("%s/%s", name, completer.Variant)
					}

					batch = append(batch, carapace.ActionStyledValuesDescribed(name, completer.Description, s).Tag(t))
				}
			}
			return batch.ToA()
		})
	})
}
