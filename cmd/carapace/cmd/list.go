package cmd

import (
	"encoding/json"
	"fmt"
	"maps"
	"slices"
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/action"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/carapace-sh/carapace-bin/pkg/completer"
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
		filter := ""
		if len(args) > 0 {
			filter = args[0]
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

		if len(args) > 0 {
			// TODO should matchall be a `*` since the default completers have an empty string as variant?
			// TODO `skhd@darwin` should be possible so its better to implicitly add a `default` variant to all (or use that term to filter)
			nameVariant, group, _ := strings.Cut(args[0], "@")
			name, variant, _ := strings.Cut(nameVariant, "/")

			filtered := make(completer.CompleterMap)
			for cName, cVariants := range c {
				if name != "" && name != cName {
					continue
				}

				for _, cVariant := range cVariants {
					if variant != "" && variant != cVariant.Variant {
						continue
					}
					if group != "" && group != cVariant.Group {
						continue
					}
					filtered[cName] = append(filtered[cName], cVariant)
				}
			}
			c = filtered
		}

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
		action.ActionCompleters().NoSpace(),
	)
}
