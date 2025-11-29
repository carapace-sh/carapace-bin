package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_pushCmd = &cobra.Command{
	Use:   "push <git-repo>",
	Short: "Push a [global] config. Leave empty to use jetify cloud. Can be a git repo for self storage.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_pushCmd).Standalone()

	global_pushCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_pushCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	global_pushCmd.Flag("config").Hidden = true
	globalCmd.AddCommand(global_pushCmd)

	// TODO environment
	carapace.Gen(global_pushCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})

	// TODO positional completion
}
