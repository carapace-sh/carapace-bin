package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var p4SyncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Import new changes from p4 into Git",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(p4SyncCmd).Standalone()

	p4SyncCmd.Flags().StringArrayP("", "/", []string{}, "Exclude selected depot paths")
	p4SyncCmd.Flags().String("branch", "", "Import changes into <ref> instead of refs/remotes/p4/master")
	p4SyncCmd.Flags().String("changes-block-size", "", "Internal block size for converting revision specifiers")
	p4SyncCmd.Flags().String("changesfile", "", "Import exactly the p4 change numbers listed in file")
	p4SyncCmd.Flags().Bool("detect-branches", false, "Use branch detection algorithm to find new paths in p4")
	p4SyncCmd.Flags().Bool("detect-labels", false, "Query p4 for labels associated with the depot paths")
	p4SyncCmd.Flags().Bool("import-labels", false, "Import labels from p4 into Git")
	p4SyncCmd.Flags().Bool("import-local", false, "Put p4 branches in refs/heads/p4/ instead of refs/remotes/p4/")
	p4SyncCmd.Flags().Bool("keep-path", false, "Retain full p4 depot path in Git")
	p4SyncCmd.Flags().String("max-changes", "", "Import at most n changes")
	p4SyncCmd.Flags().Bool("silent", false, "Do not print any progress information")
	p4SyncCmd.Flags().Bool("use-client-spec", false, "Use a client spec to find the list of interesting files")
	rootCmd.AddCommand(p4SyncCmd)

	carapace.Gen(p4SyncCmd).FlagCompletion(carapace.ActionMap{
		"":                   carapace.ActionValues(),
		"branch":             carapace.ActionValues(),
		"changes-block-size": carapace.ActionValues(),
		"changesfile":        carapace.ActionFiles(),
		"max-changes":        carapace.ActionValues(),
	})

	carapace.Gen(p4SyncCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
