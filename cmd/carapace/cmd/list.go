package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/action"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace-bridge/pkg/bridges"
	"github.com/rsteube/carapace-bridge/pkg/env"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "--list",
	Short: "list completers",
	Args:  cobra.ArbitraryArgs,
	Run: func(cmd *cobra.Command, args []string) {
		switch cmd.Flag("format").Value.String() {
		case "json":
			printCompletersJson(cmd.Flag("all").Changed, args...)
		default:
			printCompleters(cmd.Flag("all").Changed)
		}
	},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().BoolP("all", "a", false, "include bridged commands")
	listCmd.Flags().String("format", "plain", "output format")

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json").StyleF(func(s string, sc style.Context) string {
			return style.ForPathExt("."+s, sc)
		}),
	})

	carapace.Gen(listCmd).PositionalAnyCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if listCmd.Flag("format").Value.String() != "json" {
				return carapace.ActionValues()
			}
			return action.ActionCompleters(action.CompleterOpts{
				Internal: true,
				Spec:     true,
				Bridge:   listCmd.Flag("all").Changed,
			}).FilterArgs()
		}),
	)

}
func printCompleters(all bool) {
	_m := mapCompleters(all)

	maxlen := 0
	for name := range _m {
		if len := len(name); len > maxlen {
			maxlen = len
		}
	}

	for _, c := range _m {
		fmt.Printf("%-"+strconv.Itoa(maxlen)+"v %v\n", c.Name, c.Description)
	}
}

func printCompletersJson(all bool, names ...string) {
	completers := mapCompleters(all)
	if len(names) > 0 {
		filtered := make(map[string]_completer)
		for _, name := range names {
			filtered[name] = completers[name]
		}
		completers = filtered
	}

	if m, err := json.Marshal(completers); err == nil { // TODO handle error (log?)
		fmt.Println(string(m))
	}
}

type _completer struct {
	Name        string
	Description string
	Spec        string `json:",omitempty"`
	Overlay     string `json:",omitempty"`
	Bridge      string `json:",omitempty"`
}

func mapCompleters(all bool) map[string]_completer {
	// TODO move to completers package

	_completers := make(map[string]_completer, 0)
	for _, name := range completers.Names() {
		specPath, _ := completers.SpecPath(name)       // TODO handle error (log?)
		overlayPath, _ := completers.OverlayPath(name) // TODO handle error (log?)
		_completers[name] = _completer{
			Name:        name,
			Description: completers.Description(name),
			Spec:        specPath,
			Overlay:     overlayPath,
		}
	}

	if all {
		for name, bridge := range bridges.Config() {
			if _, ok := _completers[name]; ok {
				continue
			}
			specPath, _ := completers.SpecPath(name)       // TODO handle error (log?)
			overlayPath, _ := completers.OverlayPath(name) // TODO handle error (log?)
			_completers[name] = _completer{
				Name:        name,
				Description: completers.Description(name), // TODO
				Spec:        specPath,
				Overlay:     overlayPath,
				Bridge:      bridge,
			}
		}

		for _, b := range env.Bridges() {
			switch b {
			case "bash":
				for _, name := range bridges.Bash() {
					if _, ok := _completers[name]; ok {
						continue
					}

					specPath, _ := completers.SpecPath(name)       // TODO handle error (log?)
					overlayPath, _ := completers.OverlayPath(name) // TODO handle error (log?)
					_completers[name] = _completer{
						Name:        name,
						Description: completers.Description(name), // TODO
						Spec:        specPath,
						Overlay:     overlayPath,
						Bridge:      "bash",
					}
				}
			case "fish":
				for _, name := range bridges.Fish() {
					if _, ok := _completers[name]; ok {
						continue
					}

					specPath, _ := completers.SpecPath(name)       // TODO handle error (log?)
					overlayPath, _ := completers.OverlayPath(name) // TODO handle error (log?)
					_completers[name] = _completer{
						Name:        name,
						Description: completers.Description(name), // TODO
						Spec:        specPath,
						Overlay:     overlayPath,
						Bridge:      "fish",
					}
				}
			case "inshellisense":
				for _, name := range bridges.Inshellisense() {
					if _, ok := _completers[name]; ok {
						continue
					}

					specPath, _ := completers.SpecPath(name)       // TODO handle error (log?)
					overlayPath, _ := completers.OverlayPath(name) // TODO handle error (log?)
					_completers[name] = _completer{
						Name:        name,
						Description: completers.Description(name), // TODO
						Spec:        specPath,
						Overlay:     overlayPath,
						Bridge:      "inshellisense",
					}
				}
			case "zsh":
				for _, name := range bridges.Zsh() {
					if _, ok := _completers[name]; ok {
						continue
					}

					specPath, _ := completers.SpecPath(name)       // TODO handle error (log?)
					overlayPath, _ := completers.OverlayPath(name) // TODO handle error (log?)
					_completers[name] = _completer{
						Name:        name,
						Description: completers.Description(name), // TODO
						Spec:        specPath,
						Overlay:     overlayPath,
						Bridge:      "zsh",
					}
				}
			}
		}
	}

	return _completers
}
