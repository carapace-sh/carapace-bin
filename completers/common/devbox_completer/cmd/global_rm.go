package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_rmCmd = &cobra.Command{
	Use:   "rm <pkg>...",
	Short: "Remove a package from your devbox",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_rmCmd).Standalone()

	global_rmCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_rmCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	global_rmCmd.Flag("config").Hidden = true
	globalCmd.AddCommand(global_rmCmd)

	// TODO environment
	carapace.Gen(global_rmCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
