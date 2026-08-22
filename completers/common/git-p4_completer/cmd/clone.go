package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var p4CloneCmd = &cobra.Command{
	Use:   "clone",
	Short: "Create a new Git repository from an existing p4 repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(p4CloneCmd).Standalone()

	p4CloneCmd.Flags().StringArrayP("", "/", []string{}, "Exclude selected depot paths")
	p4CloneCmd.Flags().Bool("bare", false, "Perform a bare clone")
	p4CloneCmd.Flags().String("branch", "", "Import changes into <ref> instead of refs/remotes/p4/master")
	p4CloneCmd.Flags().String("changes-block-size", "", "Internal block size for converting revision specifiers")
	p4CloneCmd.Flags().String("changesfile", "", "Import exactly the p4 change numbers listed in file")
	p4CloneCmd.Flags().String("destination", "", "Where to create the Git repository")
	p4CloneCmd.Flags().Bool("detect-branches", false, "Use branch detection algorithm to find new paths in p4")
	p4CloneCmd.Flags().Bool("detect-labels", false, "Query p4 for labels associated with the depot paths")
	p4CloneCmd.Flags().Bool("import-labels", false, "Import labels from p4 into Git")
	p4CloneCmd.Flags().Bool("import-local", false, "Put p4 branches in refs/heads/p4/ instead of refs/remotes/p4/")
	p4CloneCmd.Flags().Bool("keep-path", false, "Retain full p4 depot path in Git")
	p4CloneCmd.Flags().String("max-changes", "", "Import at most n changes")
	p4CloneCmd.Flags().Bool("silent", false, "Do not print any progress information")
	p4CloneCmd.Flags().Bool("use-client-spec", false, "Use a client spec to find the list of interesting files")
	rootCmd.AddCommand(p4CloneCmd)

	carapace.Gen(p4CloneCmd).FlagCompletion(carapace.ActionMap{
		"":                   carapace.ActionValues(),
		"branch":             carapace.ActionValues(),
		"changes-block-size": carapace.ActionValues(),
		"changesfile":        carapace.ActionFiles(),
		"destination":        carapace.ActionDirectories(),
		"max-changes":        carapace.ActionValues(),
	})

	carapace.Gen(p4CloneCmd).PositionalAnyCompletion(
		carapace.ActionValues(),
	)
}
