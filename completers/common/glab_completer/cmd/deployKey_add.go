package cmd

import (
	"github.com/carapace-sh/carapace"
	"github.com/spf13/cobra"
)

var deployKey_addCmd = &cobra.Command{
	Use:   "add [key-file]",
	Short: "Add a deploy key to a GitLab project.",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	carapace.Gen(deployKey_addCmd).Standalone()

	deployKey_addCmd.Flags().BoolP("can-push", "c", false, "If true, deploy keys can be used for pushing code to the repository.")
	deployKey_addCmd.Flags().StringP("expires-at", "e", "", "The expiration date of the deploy key, using the ISO-8601 format: YYYY-MM-DDTHH:MM:SSZ.")
	deployKey_addCmd.Flags().StringP("title", "t", "", "New deploy key's title.")
	deployKey_addCmd.MarkFlagRequired("title")
	deployKeyCmd.AddCommand(deployKey_addCmd)
}
