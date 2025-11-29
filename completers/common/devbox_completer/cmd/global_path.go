package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_pathCmd = &cobra.Command{
	Use:   "path",
	Short: "Show path to [global] devbox config",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_pathCmd).Standalone()

	global_pathCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_pathCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	global_pathCmd.Flag("config").Hidden = true
	globalCmd.AddCommand(global_pathCmd)

	// TODO environment
	carapace.Gen(global_pathCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})
}
