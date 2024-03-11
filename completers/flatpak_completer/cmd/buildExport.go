package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var buildExportCmd = &cobra.Command{
	Use:     "build-export [OPTION…] LOCATION DIRECTORY [BRANCH]",
	Short:   "Create a repository from a build directory",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildExportCmd).Standalone()

	buildExportCmd.Flags().String("arch", "", "Architecture to export for (must be host compatible)")
	buildExportCmd.Flags().StringP("body", "b", "", "Full description")
	buildExportCmd.Flags().String("collection-id", "", "Collection ID")
	buildExportCmd.Flags().Bool("disable-fsync", false, "Do not invoke fsync()")
	buildExportCmd.Flags().Bool("disable-sandbox", false, "Do not sandbox icon validator")
	buildExportCmd.Flags().String("end-of-life", "", "Mark build as end-of-life")
	buildExportCmd.Flags().String("end-of-life-rebase", "", "Mark build as end-of-life, to be replaced with the given ID")
	buildExportCmd.Flags().String("exclude", "", "Files to exclude")
	buildExportCmd.Flags().String("files", "", "Use alternative directory for the files")
	buildExportCmd.Flags().String("gpg-homedir", "", "GPG Homedir to use when looking for keyrings")
	buildExportCmd.Flags().String("gpg-sign", "", "GPG Key ID to sign the commit with")
	buildExportCmd.Flags().BoolP("help", "h", false, "Show help options")
	buildExportCmd.Flags().String("include", "", "Excluded files to include")
	buildExportCmd.Flags().String("metadata", "", "Use alternative file for the metadata")
	buildExportCmd.Flags().Bool("no-summary-index", false, "Don't generate a summary index")
	buildExportCmd.Flags().Bool("no-update-summary", false, "Don't update the summary")
	buildExportCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	buildExportCmd.Flags().BoolP("runtime", "r", false, "Commit runtime (/usr), not /app")
	buildExportCmd.Flags().StringP("subject", "s", "", "One line subject")
	buildExportCmd.Flags().String("subset", "", "Add to a named subset")
	buildExportCmd.Flags().String("timestamp", "", "Override the timestamp of the commit")
	buildExportCmd.Flags().String("token-type", "", "Set type of token needed to install this commit")
	buildExportCmd.Flags().Bool("update-appstream", false, "Update the appstream branch")
	buildExportCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(buildExportCmd)

	// TODO flag
	carapace.Gen(buildExportCmd).FlagCompletion(carapace.ActionMap{
		// "arch":               carapace.ActionValues(),
		// "body":               carapace.ActionValues(),
		// "collection-id":      carapace.ActionValues(),
		// "end-of-life":        carapace.ActionValues(),
		// "end-of-life-rebase": carapace.ActionValues(),
		// "exclude":            carapace.ActionValues(),
		// "files":              carapace.ActionValues(),
		// "gpg-homedir":        carapace.ActionValues(),
		// "gpg-sign":           carapace.ActionValues(),
		// "include":            carapace.ActionValues(),
		// "metadata":           carapace.ActionValues(),
		// "subject":            carapace.ActionValues(),
		// "subset":             carapace.ActionValues(),
		// "timestamp":          carapace.ActionValues(),
		// "token-type":         carapace.ActionValues(),
	})

	carapace.Gen(buildExportCmd).PositionalCompletion(
		carapace.ActionValues(), // TODO location
		carapace.ActionDirectories(),
		carapace.ActionValues(), // TODO branch
	)
}
