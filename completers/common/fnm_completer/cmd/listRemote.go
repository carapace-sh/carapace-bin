package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var listRemoteCmd = &cobra.Command{
	Use:     "list-remote",
	Short:   "List all remote Node.js versions",
	Aliases: []string{"ls-remote"},
	Run:     func(*cobra.Command, []string) {},
}

func init() {
	carapace.Gen(listRemoteCmd).Standalone()

	listRemoteCmd.Flags().String("filter", "", "Filter versions by a user-defined version or a semver range")
	listRemoteCmd.Flags().Bool("latest", false, "Only show the latest matching version")
	listRemoteCmd.Flags().String("lts", "", "Show only LTS versions (optionally filter by LTS codename)")
	listRemoteCmd.Flags().String("sort", "asc", "Version sorting order")
	rootCmd.AddCommand(listRemoteCmd)

	carapace.Gen(listRemoteCmd).FlagCompletion(carapace.ActionMap{
		"sort": carapace.ActionValues("asc", "desc"),
	})
}
