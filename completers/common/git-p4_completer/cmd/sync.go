package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var syncCmd = &cobra.Command{
	Use:   "sync",
	Short: "Import new changes from p4 into Git",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(syncCmd).Standalone()

	syncCmd.Flags().StringArrayP("", "/", []string{}, "Exclude selected depot paths")
	syncCmd.Flags().String("branch", "", "Import changes into <ref> instead of refs/remotes/p4/master")
	syncCmd.Flags().String("changes-block-size", "", "Internal block size for converting revision specifiers")
	syncCmd.Flags().String("changesfile", "", "Import exactly the p4 change numbers listed in file")
	syncCmd.Flags().Bool("detect-branches", false, "Use branch detection algorithm to find new paths in p4")
	syncCmd.Flags().Bool("detect-labels", false, "Query p4 for labels associated with the depot paths")
	syncCmd.Flags().Bool("import-labels", false, "Import labels from p4 into Git")
	syncCmd.Flags().Bool("import-local", false, "Put p4 branches in refs/heads/p4/ instead of refs/remotes/p4/")
	syncCmd.Flags().Bool("keep-path", false, "Retain full p4 depot path in Git")
	syncCmd.Flags().String("max-changes", "", "Import at most n changes")
	syncCmd.Flags().Bool("silent", false, "Do not print any progress information")
	syncCmd.Flags().Bool("use-client-spec", false, "Use a client spec to find the list of interesting files")
	rootCmd.AddCommand(syncCmd)

	carapace.Gen(syncCmd).FlagCompletion(carapace.ActionMap{
		"":                   carapace.ActionValues(),
		"branch":             carapace.ActionValues(),
		"changes-block-size": carapace.ActionValues(),
		"changesfile":        carapace.ActionFiles(),
		"max-changes":        carapace.ActionValues(),
	})

	carapace.Gen(syncCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
