package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var cacheCmd = &cobra.Command{
	Use:     "cache",
	Short:   "Manage the local package cache",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "system maintenance",
}

func init() {
	carapace.Gen(cacheCmd).Standalone()

	cacheCmd.Flags().Bool("add-dependencies", false, "Add the argument dependencies to WORLD dependencies")
	cacheCmd.Flags().BoolP("available", "a", false, "Selected packages to be downloaded from active repositories")
	cacheCmd.Flags().Bool("ignore-conflict", false, "Ignore conflicts when resolving dependencies")
	cacheCmd.Flags().BoolP("latest", "l", false, "Always choose the latest package by version")
	cacheCmd.Flags().BoolP("simulate", "s", false, "Simulate the requested operation without making any changes")
	cacheCmd.Flags().BoolP("upgrade", "u", false, "Prefer upgrades of already installed packages")
	common.AddGlobalFlags(cacheCmd)
	rootCmd.AddCommand(cacheCmd)

	carapace.Gen(cacheCmd).PositionalCompletion(
		carapace.ActionValues("clean", "download", "purge", "sync"),
		// TODO dependency completion
	)
}
