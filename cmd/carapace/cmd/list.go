package cmd

import (
	"encoding/json"
	"fmt"
	"strconv"

	"github.com/rsteube/carapace"
	"github.com/rsteube/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/rsteube/carapace/pkg/style"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "--list",
	Short: "",
	Run: func(cmd *cobra.Command, args []string) {
		switch cmd.Flag("format").Value.String() {
		case "json":
			printCompletersJson()
		default:
			printCompleters()
		}
	},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().String("format", "plain", "output format")

	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"format": carapace.ActionValues("plain", "json").StyleF(func(s string, sc style.Context) string {
			return style.ForPathExt("."+s, sc)
		}),
	})

}
func printCompleters() {
	_m := mapCompleters()

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

func printCompletersJson() {
	if m, err := json.Marshal(mapCompleters()); err == nil { // TODO handle error (log?)
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

func mapCompleters() map[string]_completer {
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

	for _, name := range completers.FishCompleters() {
		if _, ok := _completers[name]; ok {
			continue
		}

		specPath, _ := completers.SpecPath(name)       // TODO handle error (log?)
		overlayPath, _ := completers.OverlayPath(name) // TODO handle error (log?)
		_completers[name] = _completer{
			Name:        name,
			Description: completers.Description(name),
			Spec:        specPath,
			Overlay:     overlayPath,
			Bridge:      "fish",
		}
	}

	return _completers
}
