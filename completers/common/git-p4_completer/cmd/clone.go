package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var cloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Create a new Git repository from an existing p4 repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(cloneCmd).Standalone()

	cloneCmd.Flags().StringArrayP("", "/", []string{}, "Exclude selected depot paths")
	cloneCmd.Flags().Bool("bare", false, "Perform a bare clone")
	cloneCmd.Flags().String("branch", "", "Import changes into <ref> instead of refs/remotes/p4/master")
	cloneCmd.Flags().String("changes-block-size", "", "Internal block size for converting revision specifiers")
	cloneCmd.Flags().String("changesfile", "", "Import exactly the p4 change numbers listed in file")
	cloneCmd.Flags().String("destination", "", "Where to create the Git repository")
	cloneCmd.Flags().Bool("detect-branches", false, "Use branch detection algorithm to find new paths in p4")
	cloneCmd.Flags().Bool("detect-labels", false, "Query p4 for labels associated with the depot paths")
	cloneCmd.Flags().Bool("import-labels", false, "Import labels from p4 into Git")
	cloneCmd.Flags().Bool("import-local", false, "Put p4 branches in refs/heads/p4/ instead of refs/remotes/p4/")
	cloneCmd.Flags().Bool("keep-path", false, "Retain full p4 depot path in Git")
	cloneCmd.Flags().String("max-changes", "", "Import at most n changes")
	cloneCmd.Flags().Bool("silent", false, "Do not print any progress information")
	cloneCmd.Flags().Bool("use-client-spec", false, "Use a client spec to find the list of interesting files")
	rootCmd.AddCommand(cloneCmd)

	carapace.Gen(cloneCmd).FlagCompletion(carapace.ActionMap{
		"":                   carapace.ActionValues(),
		"branch":             carapace.ActionValues(),
		"changes-block-size": carapace.ActionValues(),
		"changesfile":        carapace.ActionFiles(),
		"destination":        carapace.ActionDirectories(),
		"max-changes":        carapace.ActionValues(),
	})

	carapace.Gen(cloneCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
