package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/completers/apk_completer/cmd/common"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:     "add",
	Short:   "Add or modify constraints in WORLD and commit changes",
	Run:     func(cmd *cobra.Command, args []string) {},
	GroupID: "package installation and removal",
}

func init() {
	carapace.Gen(addCmd).Standalone()

	addCmd.Flags().Bool("initdb", false, "Initialize a new package database")
	addCmd.Flags().BoolP("latest", "l", false, "Always choose the latest package by version")
	addCmd.Flags().Bool("no-chown", false, "Do not change file owner or group")
	addCmd.Flags().BoolP("upgrade", "u", false, "Upgrade PACKAGES and it's dependencies")
	addCmd.Flags().StringP("virtual", "t", "", "Create virtual package NAME with given dependencies")
	common.AddGlobalFlags(addCmd)
	common.AddCommitFlags(addCmd)
	rootCmd.AddCommand(addCmd)

	carapace.Gen(addCmd).PositionalAnyCompletion(
		action.ActionPackageSearch(addCmd).FilterArgs(),
	)
}
