package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:     "list",
	Short:   "List installed packages",
	Aliases: []string{"ls"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(listCmd).Standalone()

	listCmd.Flags().StringP("config", "c", "", "path to directory containing a devbox.json config file")
	listCmd.Flags().String("environment", "", "environment to use, when supported (e.g.secrets support dev, prod, preview.)")
	listCmd.Flags().Bool("outdated", false, "List outdated packages")
	rootCmd.AddCommand(listCmd)

	// TODO environment
	carapace.Gen(listCmd).FlagCompletion(carapace.ActionMap{
		"config": carapace.ActionDirectories(),
	})
}
