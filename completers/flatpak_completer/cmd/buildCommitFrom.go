package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/pkg/actions/os"
	"github.com/spf13/cobra"
)

var buildCommitFromCmd = &cobra.Command{
	Use:     "build-commit-from [OPTION…] DST-REPO [DST-REF…]",
	Short:   "Make a new commit from existing commits",
	GroupID: "build",
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(buildCommitFromCmd).Standalone()

	buildCommitFromCmd.Flags().StringP("body", "b", "", "Full description")
	buildCommitFromCmd.Flags().Bool("disable-fsync", false, "Do not invoke fsync()")
	buildCommitFromCmd.Flags().String("end-of-life", "", "Mark build as end-of-life")
	buildCommitFromCmd.Flags().String("end-of-life-rebase", "", "Mark refs matching the OLDID prefix as end-of-life, to be replaced with the given NEWID")
	buildCommitFromCmd.Flags().String("extra-collection-id", "", "Add an extra collection id ref and binding")
	buildCommitFromCmd.Flags().Bool("force", false, "Always commit, even if same content")
	buildCommitFromCmd.Flags().String("gpg-homedir", "", "GPG Homedir to use when looking for keyrings")
	buildCommitFromCmd.Flags().String("gpg-sign", "", "GPG Key ID to sign the commit with")
	buildCommitFromCmd.Flags().BoolP("help", "h", false, "Show help options")
	buildCommitFromCmd.Flags().Bool("no-summary-index", false, "Don't generate a summary index")
	buildCommitFromCmd.Flags().Bool("no-update-summary", false, "Don't update the summary")
	buildCommitFromCmd.Flags().Bool("ostree-verbose", false, "Show OSTree debug information")
	buildCommitFromCmd.Flags().String("src-ref", "", "Source repo ref")
	buildCommitFromCmd.Flags().String("src-repo", "", "Source repo dir")
	buildCommitFromCmd.Flags().StringP("subject", "s", "", "One line subject")
	buildCommitFromCmd.Flags().String("subset", "", "Add to a named subset")
	buildCommitFromCmd.Flags().String("timestamp", "", "Override the timestamp of the commit (NOW for current time)")
	buildCommitFromCmd.Flags().String("token-type", "", "Set type of token needed to install this commit")
	buildCommitFromCmd.Flags().Bool("untrusted", false, "Do not trust SRC-REPO")
	buildCommitFromCmd.Flags().Bool("update-appstream", false, "Update the appstream branch")
	buildCommitFromCmd.Flags().BoolP("verbose", "v", false, "Show debug information, -vv for more detail")
	rootCmd.AddCommand(buildCommitFromCmd)

	//  TODO flag
	carapace.Gen(buildCommitFromCmd).FlagCompletion(carapace.ActionMap{
		// "body": carapace.ActionValues(),
		// "end-of-life": carapace.ActionValues(),
		// "end-of-life-rebase": carapace.ActionValues(),
		// "extra-collection-id": carapace.ActionValues(),
		"gpg-homedir": carapace.ActionDirectories(),
		"gpg-sign":    os.ActionGpgKeyIds(), // TODO handle gpg-homedir
		// "src-ref": carapace.ActionValues(),
		// "src-repo": carapace.ActionValues(),
		// "subject": carapace.ActionValues(),
		// "subset": carapace.ActionValues(),
		// "timestamp": carapace.ActionValues(),
		// "token-type"		: carapace.ActionValues(),
	})

	// TODO positional
}
