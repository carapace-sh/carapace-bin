package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "apfsctl",
	Short: "APFS filesystem control utility",
	Long:  "APFS filesystem control utility for managing datasets, purgeable files, and dataless items.",
	Run:   func(*cobra.Command, []string) {},
}

func Execute() error {
	return rootCmd.Execute()
}

func init() {
	carapace.Gen(rootCmd).Standalone()

	carapace.Gen(rootCmd).PositionalCompletion(
		carapace.ActionValuesDescribed(
			"create", "create APFS snapshots",
			"dataless", "mark files/directories as dataless",
			"dataset", "manage APFS dataset/directory classification",
			"purgeable", "manage purgeable flags on files",
		),
	)
}
