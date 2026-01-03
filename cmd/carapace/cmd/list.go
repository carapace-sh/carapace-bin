package cmd

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	carapacebin "github.com/carapace-sh/carapace-bin/pkg/actions/tools/carapace"
	"github.com/carapace-sh/carapace-bridge/pkg/choices"
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
		filter := choices.Choice{}
		if len(args) > 0 {
			filter = choices.Parse(args[0])
		}
		c, err := completers.Completers(filter, !cmd.Flag("names").Changed)
		if err != nil {
			return err
		}

		if cmd.Flag("names").Changed {
			for _, name := range slices.Sorted(maps.Keys(c)) {
				fmt.Println(name)
			}
			return nil
		}

		c.SortVariants()
		m, err := json.MarshalIndent(c, "", "  ")
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
