package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var debugbackupbundleCmd = &cobra.Command{
	Use:    "debugbackupbundle",
	Short:  "hg debugbackupbundle [--recover HASH]",
	Hidden: true,
	Run:    func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(debugbackupbundleCmd).Standalone()

	debugbackupbundleCmd.Flags().BoolP("git", "g", false, "use git extended diff format")
	debugbackupbundleCmd.Flags().BoolP("graph", "G", false, "show the revision DAG")
	debugbackupbundleCmd.Flags().StringP("limit", "l", "", "limit number of changes displayed")
	debugbackupbundleCmd.Flags().BoolP("no-merges", "M", false, "do not show merges")
	debugbackupbundleCmd.Flags().BoolP("patch", "p", false, "show patch")
	debugbackupbundleCmd.Flags().String("recover", "", "brings the specified changeset back into the repository")
	debugbackupbundleCmd.Flags().Bool("stat", false, "output diffstat-style summary of changes")
	debugbackupbundleCmd.Flags().String("style", "", "display using template map file (DEPRECATED)")
	debugbackupbundleCmd.Flags().StringP("template", "T", "", "display with template")
	rootCmd.AddCommand(debugbackupbundleCmd)
}
