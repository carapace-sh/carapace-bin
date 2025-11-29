package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var global_listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List installed packages",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(global_listCmd).Standalone()

	global_listCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	global_listCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	global_listCmd.Flags().Bool("outdated", false, "List outdated packages")
	global_listCmd.Flag("config").Hidden = true
	globalCmd.AddCommand(global_listCmd)

	// TODO environment
	carapace.Gen(global_listCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})
}
