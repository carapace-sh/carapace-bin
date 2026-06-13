package cmd

import (
	"encoding/json"
	"fmt"
	"maps"
	"os"
	"slices"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/completers"
	"github.com/carapace-sh/carapace-bin/cmd/carapace/cmd/shim"
	carapacebin "github.com/carapace-sh/carapace-bin/pkg/actions/tools/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/env"
	"github.com/carapace-sh/carapace-bridge/pkg/choices"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "--list",
	Short: "list completers",
	Args:  cobra.MaximumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		if cmd.Flag("init").Changed {
			go func() {
				if err := shim.Update(); err != nil {
					carapace.LOG.Println(err.Error())
				}
			}()
		}

		if cmd.Flag("all").Changed {
			os.Setenv(env.CARAPACE_EXCLUDES, "")
			os.Setenv(env.CARAPACE_BUILTINS, "bash,fish,zsh") // TODO add more later
		}

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
	listCmd.Flags().Bool("all", false, "return unfiltered results")
	listCmd.Flags().Bool("names", false, "only list names")
	listCmd.Flags().Bool("init", false, "update shims")

	carapace.Gen(listCmd).PositionalCompletion(
		carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if listCmd.Flag("all").Changed {
				c.Setenv(env.CARAPACE_EXCLUDES, "")
				c.Setenv(env.CARAPACE_BUILTINS, "bash,fish,zsh") // TODO add more later
			}
			return carapacebin.ActionCompleters(false).NoSpace().Invoke(c).ToA()
		}),
	)
}
