package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var versionlockCmd = &cobra.Command{
	Use:   "versionlock <subcommand> <package-spec>...",
	Short: "protect packages from updates to newer versions",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionlockAddCmd = &cobra.Command{
	Use:   "add <package-spec>...",
	Short: "add a versionlock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionlockExcludeCmd = &cobra.Command{
	Use:   "exclude <package-spec>...",
	Short: "add an exclude within versionlock",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionlockClearCmd = &cobra.Command{
	Use:   "clear",
	Short: "remove all versionlock entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionlockDeleteCmd = &cobra.Command{
	Use:   "delete <package-spec>...",
	Short: "remove matching versionlock entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

var versionlockListCmd = &cobra.Command{
	Use:   "list",
	Short: "list current versionlock entries",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(versionlockCmd).Standalone()

	versionlockCmd.AddCommand(versionlockAddCmd)
	versionlockCmd.AddCommand(versionlockExcludeCmd)
	versionlockCmd.AddCommand(versionlockClearCmd)
	versionlockCmd.AddCommand(versionlockDeleteCmd)
	versionlockCmd.AddCommand(versionlockListCmd)

	rootCmd.AddCommand(versionlockCmd)
}
