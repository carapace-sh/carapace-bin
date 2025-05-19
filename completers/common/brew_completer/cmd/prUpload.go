package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var prUploadCmd = &cobra.Command{
	Use:     "pr-upload",
	Short:   "Apply the bottle commit and publish bottles to a host",
	GroupID: "developer",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(prUploadCmd).Standalone()

	prUploadCmd.Flags().Bool("committer", false, "Specify a committer name and email in `git`'s standard author format.")
	prUploadCmd.Flags().Bool("debug", false, "Display any debugging information.")
	prUploadCmd.Flags().Bool("dry-run", false, "Print what would be done rather than doing it.")
	prUploadCmd.Flags().Bool("help", false, "Show this message.")
	prUploadCmd.Flags().Bool("keep-old", false, "If the formula specifies a rebuild version, attempt to preserve its value in the generated DSL.")
	prUploadCmd.Flags().Bool("no-commit", false, "Do not generate a new commit before uploading.")
	prUploadCmd.Flags().Bool("quiet", false, "Make some output more quiet.")
	prUploadCmd.Flags().Bool("root-url", false, "Use the specified <URL> as the root of the bottle's URL instead of Homebrew's default.")
	prUploadCmd.Flags().Bool("root-url-using", false, "Use the specified download strategy class for downloading the bottle's URL instead of Homebrew's default.")
	prUploadCmd.Flags().Bool("upload-only", false, "Skip running `brew bottle` before uploading.")
	prUploadCmd.Flags().Bool("verbose", false, "Make some output more verbose.")
	prUploadCmd.Flags().Bool("warn-on-upload-failure", false, "Warn instead of raising an error if the bottle upload fails. Useful for repairing bottle uploads that previously failed.")
	rootCmd.AddCommand(prUploadCmd)
}
