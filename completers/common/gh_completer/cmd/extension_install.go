package cmd

import (
	"strings"

	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/gh_completer/cmd/action"
	"github.com/carapace-sh/carapace-bin/pkg/actions/tools/git"
	"github.com/spf13/cobra"
)

var extension_installCmd = &cobra.Command{
	Use:   "install <repository>",
	Short: "Install a gh extension from a repository",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(extension_installCmd).Standalone()

	extension_installCmd.Flags().Bool("force", false, "Force upgrade extension, or ignore if latest already installed")
	extension_installCmd.Flags().String("pin", "", "Pin extension to a release tag or commit ref")
	extensionCmd.AddCommand(extension_installCmd)

	carapace.Gen(extension_installCmd).FlagCompletion(carapace.ActionMap{
		"pin": carapace.ActionCallback(func(c carapace.Context) carapace.Action {
			if len(c.Args) == 0 {
				return carapace.ActionValues()
			}
			url := c.Args[0]
			if splitted := strings.Split(c.Args[0], "/"); len(splitted) == 2 { // just owner/repo
				url = "https://github.com/" + url
			}
			return git.ActionLsRemoteRefs(git.LsRemoteRefOption{Url: url, Branches: true, Tags: true})
		}),
	})

	carapace.Gen(extension_installCmd).PositionalCompletion(
		carapace.Batch(
			action.ActionTopExtensions(),
			action.ActionOwnerRepositories(extension_installCmd),
		).ToA(),
	)
}
