package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/devbox"
	"github.com/carapace-sh/carapace/pkg/traverse"
	"github.com/spf13/cobra"
)

var rmCmd = &cobra.Command{
	Use:   "rm <pkg>...",
	Short: "Remove a package from your devbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(rmCmd).Standalone()

	rmCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	rmCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	rootCmd.AddCommand(rmCmd)

	// TODO environment
	carapace.Gen(rmCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	carapace.Gen(rmCmd).PositionalAnyCompletion(
		devbox.ActionInstalledPackages().
			ChdirF(traverse.Flag(rmCmd.Flag("config"))).
			FilterArgs(),
	)
}
