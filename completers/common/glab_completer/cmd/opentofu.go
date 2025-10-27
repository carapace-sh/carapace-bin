package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/carapace-sh/carapace-bin/completers/common/glab_completer/cmd/action"
	"github.com/spf13/cobra"
)

var opentofuCmd = &cobra.Command{
	Use:     "opentofu <command> [flags]",
	Short:   "Work with the OpenTofu or Terraform integration.",
	Aliases: []string{"terraform", "tf"},
	Run:     func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(opentofuCmd).Standalone()

	opentofuCmd.PersistentFlags().StringP("repo", "R", "", "Select another repository. Can use either `OWNER/REPO` or `GROUP/NAMESPACE/REPO` format. Also accepts full URL or Git URL.")
	rootCmd.AddCommand(opentofuCmd)

	carapace.Gen(opentofuCmd).FlagCompletion(carapace.ActionMap{
		"repo": action.ActionRepo(opentofuCmd),
	})
}
