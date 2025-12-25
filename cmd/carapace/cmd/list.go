package cmd

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	carapacebin "github.com/carapace-sh/carapace-bin/pkg/actions/tools/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
	"github.com/carapace-sh/carapace-bridge/pkg/choice"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "--list",
	Short: "list completers",
	Args:  cobra.MaximumNArgs(1),
	FParseErrWhitelist: cobra.FParseErrWhitelist{
		UnknownFlags: true, // TODO remove - just to keep compability with tabdance until things are merged
	},
	RunE: func(cmd *cobra.Command, args []string) error {
		filter := choice.Choice{}
		if len(args) > 0 {
			filter = choice.Parse(args[0])
		}
		c, err := completers.Completers(filter, !cmd.Flag("names").Changed)
		if err != nil {
			return err
		}

		// TODO add in completers.Completers?
		c["carapace"] = completer.Completers{
			{
				Name:  "carapace",
				Group: "common",
			},
		}
		c = c.Filter(filter) // TODO this is bullshit, the pseudo carapace completer needs to be added before filtering

		if cmd.Flag("names").Changed {
			for _, name := range slices.Sorted(maps.Keys(c)) {
				fmt.Println(name)
			}
			return nil
		}

		c.SortVariants() // TODO implicitly do this in marshal?
		m, err := json.Marshal(c)
		if err != nil {
			return err
		}
		fmt.Println(string(m))
		return nil
	},
}

func init() {
	carapace.Gen(listCmd).Standalone()
	listCmd.Flags().Bool("names", false, "only list names")

	carapace.Gen(listCmd).PositionalCompletion(
		carapacebin.ActionCompleters().NoSpace(),
	)
}
